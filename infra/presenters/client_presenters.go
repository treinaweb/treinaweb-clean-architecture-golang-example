package presenters

import (
	"encoding/json"

	"github.com/cleysonph/clean-arch-go/enterprise/entities"
)

type FindAllClientsOutput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
}

func (FindAllClientsOutput) ToJson(clients []entities.Client) []byte {
	output := make([]FindAllClientsOutput, len(clients))
	for i, client := range clients {
		output[i] = FindAllClientsOutput{
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			Cpf:       client.Cpf,
		}
	}
	json, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	return json
}

type CreateClientOutput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
}

func (CreateClientOutput) ToJson(client entities.Client) []byte {
	output := CreateClientOutput{
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Email:     client.Email,
		Cpf:       client.Cpf,
	}
	json, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	return json
}

type CreateClientIntput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
}

func (CreateClientIntput) FromJson(jsonData []byte) entities.Client {
	var clientInput CreateClientIntput
	err := json.Unmarshal(jsonData, &clientInput)
	if err != nil {
		panic(err)
	}
	client, err := entities.NewClient(clientInput.FirstName, clientInput.LastName, clientInput.Email, clientInput.Cpf)
	if err != nil {
		panic(err)
	}
	return client
}
