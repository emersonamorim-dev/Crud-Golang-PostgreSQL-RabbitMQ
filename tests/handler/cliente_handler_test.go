package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Crud-Golang-RabbitMQ/internal/handler"
	"Crud-Golang-RabbitMQ/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestClienteHandler(t *testing.T) {
	// Mock do serviço de cliente
	mockService := &mockClienteService{}
	handler := handler.NewClientesHandler(mockService)

	// Config do roteador Gin
	router := gin.Default()
	router.POST("/clientes", handler.CreateCliente)
	router.GET("/clientes", handler.ListarClientes)
	router.GET("/clientes/:id", handler.GetClienteByID)
	router.PUT("/clientes/:id", handler.UpdateCliente)
	router.DELETE("/clientes/:id", handler.DeleteCliente)

	t.Run("CreateCliente_Success", func(t *testing.T) {
		// Caso de sucesso ao criar um cliente
		cliente := model.Clientes{
			Nome:       "Emerson",
			Sobrenome:  "Amorim",
			Contato:    "emerson_tecno@hotmail.com",
			Endereco:   "123 Main St",
			Nascimento: "1981-02-18",
			CPF:        "12345678901",
		}
		payload, _ := json.Marshal(cliente)
		req, _ := http.NewRequest("POST", "/clientes", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})
}

// Implementei do serviço de cliente mock para uso nos testes
type mockClienteService struct{}

func (m *mockClienteService) CreateCliente(ctx context.Context, cliente model.Clientes) error {
	// Mock da função CreateCliente do serviço
	return nil
}

func (m *mockClienteService) ListarClientes(ctx context.Context) ([]model.Clientes, error) {
	// Mock da função ListarClientes do serviço
	return nil, nil
}

func (m *mockClienteService) GetClienteByID(ctx context.Context, id uint) (model.Clientes, error) {
	// Mock da função GetClienteByID do serviço
	return model.Clientes{}, nil
}

func (m *mockClienteService) UpdateCliente(ctx context.Context, cliente model.Clientes) error {
	// Mock da função UpdateCliente do serviço
	return nil
}

func (m *mockClienteService) DeleteCliente(ctx context.Context, id uint) error {
	// Mock da função DeleteCliente do serviço
	return nil
}
