package main

import (
	"github.com/cleysonph/clean-arch-go/main/adapters"
	"github.com/cleysonph/clean-arch-go/main/factories"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/clients", adapters.GinAdapter(factories.GetFindAllClientsController()))
	r.POST("/api/clients", adapters.GinAdapter(factories.GetCreateClientController()))

	r.Run(":3000")
}
