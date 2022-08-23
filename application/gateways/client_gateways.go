package gateways

import "github.com/cleysonph/clean-arch-go/enterprise/entities"

type ClientGateway interface {
	FindAll() []entities.Client
	Create(client entities.Client)
}
