package iocsvc

import (
	"os"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	iockey "github.com/ahl5esoft/lite-go/model/enum/ioc-key"
	"github.com/ahl5esoft/lite-go/service/cmdsvc"
	"github.com/ahl5esoft/lite-go/service/cryptosvc"
	"github.com/ahl5esoft/lite-go/service/execsvc"
	"github.com/ahl5esoft/lite-go/service/filesvc"
	"github.com/ahl5esoft/lite-go/service/fmtsvc"
	"github.com/ahl5esoft/lite-go/service/genericsvc"
	"github.com/ahl5esoft/lite-go/service/goredissvc"
	"github.com/ahl5esoft/lite-go/service/httpsvc"
	"github.com/ahl5esoft/lite-go/service/mongosvc"
	"github.com/ahl5esoft/lite-go/service/ossvc"
	"github.com/ahl5esoft/lite-go/service/pathsvc"
	"github.com/ahl5esoft/lite-go/service/redissvc"
	"github.com/ahl5esoft/lite-go/service/timesvc"
	"github.com/ahl5esoft/lite-go/service/yamlsvc"

	"github.com/go-redis/redis/v8"
)

// 初始化
func Init[T mcontract.IDefaultConfig](yaml string) (res T, err error) {
	var wd string
	if wd, err = os.Getwd(); err != nil {
		return
	}

	osPath := pathsvc.NewOsPath(wd)
	Set(osPath)

	configLoader := yamlsvc.NewConfigLoader(osPath, yaml)
	if err = configLoader.Load(&res); err != nil {
		return
	}

	Set(configLoader)

	SetWithName(
		iockey.AuthCrypto,
		cryptosvc.NewAesCrypto(
			[]byte(
				res.GetAuthCipher(),
			),
		),
	)

	Set(
		cmdsvc.NewCommandFactory(func(name string, args []string) contract.ICommand {
			return execsvc.NewCommand(name, args)
		}),
	)

	if res.GetMongo() != "" {
		Set(
			mongosvc.NewDbFactory(
				res.GetName(),
				res.GetMongo(),
			),
		)
	}

	Set(
		filesvc.NewFileFactory(osPath, func(self contract.IFileFactory, path string) contract.IFileEntry {
			return ossvc.NewFileEntry(self, osPath, path)
		}, func(entry contract.IFileEntry) contract.IFile {
			return ossvc.NewFile(entry)
		}),
	)

	Set(
		genericsvc.NewFactory[contract.ILogFactory](func() contract.ILog {
			return fmtsvc.NewLog()
		}),
	)

	if res.GetRedis().GetAddr() != "" {
		goRedis := goredissvc.NewRedis(
			goredissvc.OptionsRedisOption(&redis.Options{
				Addr:     res.GetRedis().GetAddr(),
				Password: res.GetRedis().GetPassword(),
			}),
		)
		Set(goRedis)
		Set(
			redissvc.NewNowTime(goRedis),
		)
	} else {
		Set(
			timesvc.NewNowTime(),
		)
	}

	if res.GetGateway() != "" {
		Set(
			genericsvc.NewFactory[contract.IRpcFactory](func() contract.IRpc {
				return httpsvc.NewRpc(
					res.GetGateway(),
				)
			}),
		)
	}

	Set(
		mongosvc.NewStringGenerator(),
	)

	return
}
