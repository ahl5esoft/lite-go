package ginsvc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/ahl5esoft/lite-go/contract"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/ahl5esoft/lite-go/service/errorsvc"
	"github.com/ahl5esoft/lite-go/service/iocsvc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type routeParam struct {
	API      string `uri:"api"`
	Endpoint string `uri:"endpoint"`
}

func NewPostOption(
	apiFactory contract.IApiFactory,
	logFactory contract.ILogFactory,
) Option {
	return func(app *gin.Engine) {
		validate := validator.New()
		app.POST("/:endpoint/:api", func(ctx *gin.Context) {
			var api contract.IApi
			var resp message.API
			var err error
			defer func() {
				ctx.JSON(http.StatusOK, resp)
			}()

			log := logFactory.Build().AddLabel("route", ctx.Request.RequestURI)
			defer func() {
				var cErr error
				if rv := recover(); rv != nil {
					var ok bool
					if cErr, ok = rv.(error); !ok {
						cErr = fmt.Errorf("%v", rv)
					}

					log.AddLabel(
						"stack",
						"%s",
						debug.Stack(),
					)
				}

				if cErr == nil && err != nil {
					cErr = err
				}

				if cErr != nil {
					if sErr, ok := cErr.(contract.IError); ok {
						resp.Error = sErr.GetCode()
						if sErr.GetData() != nil {
							resp.Data = sErr.GetData()
						} else {
							resp.Data = sErr.Error()
						}
					} else {
						log.Error(cErr)
						resp.Error = errorcode.Panic
					}
				}
			}()

			if len(ctx.Request.Header) > 0 {
				log.AddLabel("header", "%v", ctx.Request.Header)
			}

			var rp routeParam
			_ = ctx.ShouldBindUri(&rp)
			if api, err = apiFactory.Build(rp.Endpoint, rp.API); err != nil {
				return
			}

			iocsvc.Inject(api, func(v reflect.Value) reflect.Value {
				if traceable, ok := v.Interface().(contract.ITraceable); ok {
					if span, ok := ctx.Value("ParentSpan").(opentracing.Span); ok {
						newCtx := opentracing.ContextWithSpan(ctx, span)
						return traceable.WithContext(newCtx)
					}
				}

				return v
			})

			if ctx.Request.ContentLength > 0 {
				var bodyBytes []byte
				if bodyBytes, err = io.ReadAll(ctx.Request.Body); err != nil {
					return
				}

				ctx.Set("body", bodyBytes)

				log.AddLabel(
					"body",
					string(bodyBytes),
				)
				if err = json.Unmarshal(bodyBytes, &api); err != nil {
					return
				}

				if err = validate.Struct(api); err != nil {
					err = errorsvc.Newf(errorcode.Verify, "")
					return
				}
			}

			if v, ok := api.(contract.IApiSession[*gin.Context]); ok {
				err = v.SetSession(ctx)
			}
			if err != nil {
				return
			}

			resp.Data, err = api.Call()
		})
	}
}
