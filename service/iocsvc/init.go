package iocsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	iockey "github.com/ahl5esoft/lite-go/model/enum/ioc-key"
	"github.com/ahl5esoft/lite-go/service/cmdsvc"
	"github.com/ahl5esoft/lite-go/service/cryptosvc"
	"github.com/ahl5esoft/lite-go/service/execsvc"
	"github.com/ahl5esoft/lite-go/service/fmtsvc"
	"github.com/ahl5esoft/lite-go/service/genericsvc"
	"github.com/ahl5esoft/lite-go/service/goredissvc"
	"github.com/ahl5esoft/lite-go/service/httpsvc"
	"github.com/ahl5esoft/lite-go/service/mongosvc"
	"github.com/ahl5esoft/lite-go/service/pathsvc"
	"github.com/ahl5esoft/lite-go/service/timesvc"
	"github.com/ahl5esoft/lite-go/service/yamlsvc"
	"github.com/go-redis/redis/v8"
)

// 初始化
func Init[T mcontract.IDefaultConfig](yaml string) (res T, err error) {
	ioPath := pathsvc.NewIOPath()
	Set(ioPath)

	configLoader := yamlsvc.NewConfigLoader(ioPath, yaml)
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
		genericsvc.NewFactory[contract.ILogFactory](func() contract.ILog {
			return fmtsvc.NewLog()
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
		timesvc.NewNowTime(),
	)

	if res.GetRedis().GetAddr() != "" {
		Set(
			goredissvc.NewRedis(
				goredissvc.OptionsRedisOption(&redis.Options{
					Addr:     res.GetRedis().GetAddr(),
					Password: res.GetRedis().GetPassword(),
				}),
			),
		)
	}

	Set(
		mongosvc.NewStringGenerator(),
	)

	return
}
