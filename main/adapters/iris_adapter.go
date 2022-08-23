package adapters

import (
	"github.com/cleysonph/clean-arch-go/infra/controllers"
	"github.com/cleysonph/clean-arch-go/infra/web"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func IrisAdapter(controller controllers.Controller) func(c *context.Context) {
	return func(ctx iris.Context) {
		jsonData, err := ctx.GetBody()
		if err != nil {
			panic(err)
		}
		request := web.HttpRequest{
			Body: jsonData,
		}
		response := controller.Execute(request)
		ctx.StatusCode(response.StatusCode)
		ctx.Header("Content-Type", response.Headers.ContentType)
		ctx.Write(response.Body)
	}
}
