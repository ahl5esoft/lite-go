package iocsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/service/fmtsvc"
	"github.com/ahl5esoft/lite-go/service/logsvc"
	"github.com/ahl5esoft/lite-go/service/mongosvc"
	"github.com/ahl5esoft/lite-go/service/pathsvc"
	"github.com/ahl5esoft/lite-go/service/timesvc"
	"github.com/ahl5esoft/lite-go/service/yamlsvc"
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
		logsvc.NewLogFactory(func() contract.ILog {
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

	if cfg.GetRedis().Host != "" {
		// 	addr := cfg.GetRedis().Host
		// 	if !strings.Contains(addr, ":") {
		// 		addr = fmt.Sprintf("%s:6379", addr)
		// 	}
		// 	Set(
		// 		goredissvc.NewRedis(
		// 			goredissvc.RedisOptions(&redis.Options{
		// 				Addr:     addr,
		// 				Password: cfg.GetRedis().Password,
		// 			}),
		// 		),
		// 	)
	}

	Set(
		mongosvc.NewStringGenerator(),
	)

	return
}
