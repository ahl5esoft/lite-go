package ginsvc

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 端口
func NewPortOption(name string, port int) Option {
	return func(app *gin.Engine) {
		fmt.Printf(
			"%s:%d[%s]\n",
			name,
			port,
			time.Now().Format("2006-01-02 15:04:05"),
		)
		err := app.Run(
			fmt.Sprintf(":%d", port),
		)
		if err != nil {
			panic(err)
		}
	}
}
