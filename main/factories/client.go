package factories

import (
	"github.com/cleysonph/clean-arch-go/application/gateways"
	"github.com/cleysonph/clean-arch-go/application/usecases"
	"github.com/cleysonph/clean-arch-go/infra/controllers"
	gatewaysImpl "github.com/cleysonph/clean-arch-go/infra/gateways"
)

var clientGateway gateways.ClientGateway

func GetClientGateway() gateways.ClientGateway {
	if clientGateway == nil {
		clientGateway = gatewaysImpl.NewClientMySqlGateway()
	}
	return clientGateway
}

func GetFindAllClientsUseCase() usecases.FindAllClientsUseCase {
	return usecases.NewFindAllClientsUseCase(GetClientGateway())
}

func GetFindAllClientsController() controllers.Controller {
	return controllers.NewFindAllClientsController(GetFindAllClientsUseCase())
}

func GetCreateClientUseCase() usecases.CreateClientUseCase {
	return usecases.NewCreateClientUseCase(GetClientGateway())
}

func GetCreateClientController() controllers.Controller {
	return controllers.NewCreateClientController(GetCreateClientUseCase())
}
