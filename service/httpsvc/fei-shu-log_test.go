package httpsvc

import (
	"testing"
)

func Test_NewFeiShuLog_it_Info(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		NewFeiShuLog("dev-ops:", "飞书weebhook地址").
			AddLabel("测试", "dev-ops by go").
			Info()
	})
}
