package adapters

import (
	"io"

	"github.com/cleysonph/clean-arch-go/infra/controllers"
	"github.com/cleysonph/clean-arch-go/infra/web"
	"github.com/gin-gonic/gin"
)

func GinAdapter(controller controllers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		request := web.HttpRequest{
			Body: jsonData,
		}
		response := controller.Execute(request)
		c.Data(response.StatusCode, response.Headers.ContentType, []byte(response.Body))
	}
}
