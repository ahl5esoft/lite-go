package ginsvc

import (
	"net/http"
	"regexp"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/gin-gonic/gin"
)

var versionReg = regexp.MustCompile(`\d+\.\d+\.\d+`)

func NewGetVersionEngineHandler(fileFactory contract.IFileFactory, filePath contract.IFilePath, project string) IHandler {
	return NewEngineHandler(func(app *gin.Engine) (err error) {
		file := fileFactory.BuildFile(
			filePath.Getwd(),
			"README.md",
		)

		var text string
		if err = file.Read(&text); err != nil {
			return
		}

		version := "0.0.0"
		versions := versionReg.FindAllString(text, 1)
		if len(versions) > 0 {
			version = versions[0]
		}

		app.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, message.ApiResponse{
				Data: map[string]string{
					"project": project,
					"version": version,
				},
			})
		})

		return

	})
}
