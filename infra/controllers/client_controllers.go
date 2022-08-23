package controllers

import (
	"github.com/cleysonph/clean-arch-go/application/usecases"
	"github.com/cleysonph/clean-arch-go/infra/presenters"
	"github.com/cleysonph/clean-arch-go/infra/web"
)

func NewFindAllClientsController(findAllClientsUseCase usecases.FindAllClientsUseCase) Controller {
	return FindAllClientsController{findAllClientsUseCase}
}

type FindAllClientsController struct {
	findAllClientsUseCase usecases.FindAllClientsUseCase
}

// Execute implements Controller
func (c FindAllClientsController) Execute(request web.HttpRequest) web.HttpResponse {
	clients := c.findAllClientsUseCase.Execute()
	return web.HttpResponse{
		StatusCode: 200,
		Body:       presenters.FindAllClientsOutput{}.ToJson(clients),
		Headers:    web.Headers{ContentType: "application/json"},
	}
}

func NewCreateClientController(createClientUseCase usecases.CreateClientUseCase) Controller {
	return CreateClientController{createClientUseCase}
}

type CreateClientController struct {
	createClientUseCase usecases.CreateClientUseCase
}

// Execute implements Controller
func (c CreateClientController) Execute(request web.HttpRequest) web.HttpResponse {
	client := presenters.CreateClientIntput{}.FromJson(request.Body)
	createdClient := c.createClientUseCase.Execute(client)
	return web.HttpResponse{
		StatusCode: 201,
		Body:       presenters.CreateClientOutput{}.ToJson(createdClient),
		Headers:    web.Headers{ContentType: "application/json"},
	}
}
