package iocsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/service/cmdsvc"
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
func Init[T any](yaml string, t *T) (err error) {
	ioPath := pathsvc.NewIOPath()
	Set(ioPath)

	configLoader := yamlsvc.NewConfigLoader(ioPath, yaml)
	if err = configLoader.Load(t); err != nil {
		return
	}

	var cfg mcontract.IDefaultConfig
	var ok bool
	if cfg, ok = any(t).(mcontract.IDefaultConfig); !ok {
		err = fmt.Errorf("非mcontract.IDefaultConfig")
		return
	}

	Set(configLoader)

	Set(
		cmdsvc.NewCommandFactory(func(name string, args []string) contract.ICommand {
			return execsvc.NewCommand(name, args)
		}),
	)

	if cfg.GetGateway() != "" {
		Set(
			genericsvc.NewFactory[contract.IRpcFactory](func() contract.IRpc {
				return httpsvc.NewRpc(
					cfg.GetGateway(),
				)
			}),
		)
	}

	Set(
		genericsvc.NewFactory[contract.ILogFactory](func() contract.ILog {
			return fmtsvc.NewLog()
		}),
	)

	if cfg.GetMongo() != "" {
		Set(
			mongosvc.NewDbFactory(
				cfg.GetName(),
				cfg.GetMongo(),
			),
		)
	}

	Set(
		timesvc.NewNowTime(),
	)

	if cfg.GetRedis().GetAddr() != "" {
		Set(
			goredissvc.NewRedis(
				goredissvc.OptionsRedisOption(&redis.Options{
					Addr:     cfg.GetRedis().GetAddr(),
					Password: cfg.GetRedis().GetPassword(),
				}),
			),
		)
	}

	Set(
		mongosvc.NewStringGenerator(),
	)

	return
}
