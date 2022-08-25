package httpsvc

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/logsvc"
	jsoniter "github.com/json-iterator/go"
)

// 飞书日志
func NewFeiShuLog(keyword, url string) contract.ILog {
	logAction := func(labels [][2]string) {
		var res []interface{}
		underscore.Chain(labels).Map(func(r [2]string, _ int) interface{} {
			return []map[string]interface{}{
				{
					"tag":  "text",
					"text": fmt.Sprintf("%s: %s", r[0], r[1]),
				},
			}
		}).Value(&res)
		s, err := jsoniter.MarshalToString(map[string]interface{}{
			"msg_type": "post",
			"content": map[string]interface{}{
				"post": map[string]interface{}{
					"zh_cn": map[string]interface{}{
						"content": res,
						"title":   keyword,
					},
				},
			},
		})
		if err == nil {
			http.Post(
				url,
				"application/json",
				bytes.NewBufferString(s),
			)
		}
	}
	return &logsvc.LogProxy{
		DebugAction: logAction,
		ErrorAction: func(err error, labels [][2]string) {
			labels = append(labels, [2]string{"错误", err.Error()})
			labels = append(labels, [2]string{
				"堆栈",
				string(
					debug.Stack(),
				),
			})
			logAction(labels)
		},
		InfoAction:    logAction,
		WarningAction: logAction,
	}
}
