package main

import (
	"github.com/cleysonph/clean-arch-go/main/adapters"
	"github.com/cleysonph/clean-arch-go/main/factories"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/api/clients", adapters.IrisAdapter(factories.GetFindAllClientsController()))
	app.Post("/api/clients", adapters.IrisAdapter(factories.GetCreateClientController()))

	app.Listen(":3001")
}
