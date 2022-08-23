package usecases

import (
	"github.com/cleysonph/clean-arch-go/application/gateways"
	"github.com/cleysonph/clean-arch-go/enterprise/entities"
)

func NewFindAllClientsUseCase(clientGateway gateways.ClientGateway) FindAllClientsUseCase {
	return FindAllClientsUseCase{clientGateway}
}

type FindAllClientsUseCase struct {
	clientGateway gateways.ClientGateway
}

func (uc *FindAllClientsUseCase) Execute() []entities.Client {
	return uc.clientGateway.FindAll()
}

func NewCreateClientUseCase(clientGateway gateways.ClientGateway) CreateClientUseCase {
	return CreateClientUseCase{clientGateway}
}

type CreateClientUseCase struct {
	clientGateway gateways.ClientGateway
}

func (uc *CreateClientUseCase) Execute(client entities.Client) entities.Client {
	uc.clientGateway.Create(client)
	return client
}
