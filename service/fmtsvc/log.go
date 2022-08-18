package fmtsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/logsvc"
)

// 创建格式化输出
func NewLog() contract.ILog {
	action := func(label map[string]string) {
		fmt.Printf("%v\n", label)
	}
	return &logsvc.LogProxy{
		DebugAction: action,
		ErrorAction: func(label map[string]string, err error) {
			if label == nil {
				label = map[string]string{
					"err": err.Error(),
				}
			} else {
				label["err"] = err.Error()
			}
			action(label)
		},
		InfoAction:    action,
		WarningAction: action,
	}
}
