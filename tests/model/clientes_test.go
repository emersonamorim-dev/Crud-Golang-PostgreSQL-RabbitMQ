package model_test

import (
	"Crud-Golang-RabbitMQ/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientesModel(t *testing.T) {
	now := time.Now()
	cliente := model.Clientes{
		ID:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  nil,
		Nome:       "Emerson",
		Sobrenome:  "Amorim - Full Stack",
		Contato:    "11999887766",
		Endereco:   "Rua sn, 123",
		Nascimento: "1981-02-18",
		CPF:        "12345678901",
	}

	// Testa se os campos são atribuídos corretamente
	assert.Equal(t, uint(1), cliente.ID)
	assert.Equal(t, now, cliente.CreatedAt)
	assert.Equal(t, now, cliente.UpdatedAt)
	assert.Nil(t, cliente.DeletedAt)
	assert.Equal(t, "Emerson", cliente.Nome)
	assert.Equal(t, "Amorim - Full Stack", cliente.Sobrenome)
	assert.Equal(t, "11999887766", cliente.Contato)
	assert.Equal(t, "Rua sn, 123", cliente.Endereco)
	assert.Equal(t, "1981-02-18", cliente.Nascimento)
	assert.Equal(t, "12345678901", cliente.CPF)

}
