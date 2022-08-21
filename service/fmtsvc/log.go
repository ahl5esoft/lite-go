package fmtsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/logsvc"
)

// 创建格式化输出
func NewLog() contract.ILog {
	action := func(labels [][2]string) {
		for _, r := range labels {
			fmt.Println(r[0], "->", r[1])
		}
	}
	return &logsvc.LogProxy{
		DebugAction: action,
		ErrorAction: func(err error, labels [][2]string) {
			if labels == nil {
				labels = make([][2]string, 0)
			}
			labels = append(labels, [2]string{"err", err.Error()})
			action(labels)
		},
		InfoAction:    action,
		WarningAction: action,
	}
}
