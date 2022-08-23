package gateways

import (
	"github.com/cleysonph/clean-arch-go/application/gateways"
	"github.com/cleysonph/clean-arch-go/enterprise/entities"
	"github.com/cleysonph/clean-arch-go/infra/db"
)

func initialClients() []entities.Client {
	client, err := entities.NewClient("Cleyson", "Lima", "cleyson@mail.com", "12345678912")
	if err != nil {
		panic(err)
	}
	return []entities.Client{client}
}

var clients = initialClients()

func NewClientInMemoryGateway() gateways.ClientGateway {
	return ClientInMemoryGateway{}
}

type ClientInMemoryGateway struct{}

// Create implements gateways.ClientGateway
func (ClientInMemoryGateway) Create(client entities.Client) {
	clients = append(clients, client)
}

// FindAll implements gateways.ClientGateway
func (ClientInMemoryGateway) FindAll() []entities.Client {
	return clients
}

type clientDb struct {
	id        int
	nome      string
	sobrenome string
	email     string
	cpf       string
}

func NewClientMySqlGateway() gateways.ClientGateway {
	return ClientMySqlGateway{}
}

type ClientMySqlGateway struct{}

// Create implements gateways.ClientGateway
func (ClientMySqlGateway) Create(client entities.Client) {
	connection := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection()

	_, err := connection.Exec(
		"INSERT INTO clients (nome, sobrenome, email, cpf) VALUES (?, ?, ?, ?)",
		client.FirstName,
		client.LastName,
		client.Email,
		client.Cpf,
	)
	if err != nil {
		panic(err)
	}
}

// FindAll implements gateways.ClientGateway
func (ClientMySqlGateway) FindAll() []entities.Client {
	connection := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection()

	query, err := connection.Query("SELECT * FROM clients")
	if err != nil {
		panic(err)
	}

	var clients []entities.Client
	for query.Next() {
		data := clientDb{}
		err := query.Scan(&data.id, &data.nome, &data.sobrenome, &data.email, &data.cpf)
		if err != nil {
			panic(err)
		}
		client, err := entities.NewClient(data.nome, data.sobrenome, data.email, data.cpf)
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}
	return clients
}
