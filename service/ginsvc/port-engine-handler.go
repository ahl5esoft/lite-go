package ginsvc

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPortEngineHandler(port int, project string) IHandler {
	return NewEngineHandler(func(app *gin.Engine) error {
		fmt.Printf(
			"%s:%d[%s]\n",
			project,
			port,
			time.Now().Format("2006-01-02 15:04:05"),
		)
		return app.Run(
			fmt.Sprintf(":%d", port),
		)
	})
}
