package service_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/queue"
	"Crud-Golang-RabbitMQ/internal/service"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

type mockClienteRepository struct{}

func (m *mockClienteRepository) Create(ctx context.Context, cliente model.Clientes) error {
	return nil
}

func (m *mockClienteRepository) ListarTodos(ctx context.Context) ([]model.Clientes, error) {
	return []model.Clientes{}, nil
}

func (m *mockClienteRepository) GetByID(ctx context.Context, id uint) (model.Clientes, error) {
	if id == 1 {
		return model.Clientes{}, nil
	}
	return model.Clientes{}, errors.New("cliente não encontrado")
}

func (m *mockClienteRepository) Update(ctx context.Context, cliente model.Clientes) error {
	return nil
}

func (m *mockClienteRepository) Delete(ctx context.Context, id uint) error {
	if id != 1 {
		return errors.New("falha ao deletar cliente")
	}
	return nil
}

type mockQueueClient struct{}

func (m *mockQueueClient) PublishMessage(queueName string, message []byte) error {
	return nil
}

// Defini a interface QueueClient conforme esperado pela camada de serviço
type QueueClient interface {
	PublishMessage(queueName string, message []byte) error
}

// função para o createMockRabbitMQClient
func createMockRabbitMQClient() *queue.RabbitMQ {
	return &queue.RabbitMQ{}
}

// Testei a função createMockRabbitMQClient para injetar a dependência mockada
func TestClienteService_CreateCliente(t *testing.T) {
	repo := &mockClienteRepository{}
	queueClient := createMockRabbitMQClient()

	svc := service.NewClienteService(repo, queueClient)

	err := svc.CreateCliente(context.Background(), model.Clientes{})
	assert.NoError(t, err)
}

func TestClienteService_ListarClientes(t *testing.T) {
	repo := &mockClienteRepository{}
	queueClient := createMockRabbitMQClient()

	svc := service.NewClienteService(repo, queueClient)

	clientes, err := svc.ListarClientes(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, clientes)
}

func TestClienteService_GetClienteByID(t *testing.T) {
	repo := &mockClienteRepository{}
	queueClient := createMockRabbitMQClient()

	svc := service.NewClienteService(repo, queueClient)

	cliente, err := svc.GetClienteByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, cliente)

	cliente, err = svc.GetClienteByID(context.Background(), 2)
	assert.Error(t, err)
	assert.EqualError(t, err, "cliente não encontrado")
	assert.Equal(t, model.Clientes{}, cliente)
}

func TestClienteService_UpdateCliente(t *testing.T) {
	repo := &mockClienteRepository{}
	queueClient := createMockRabbitMQClient()

	svc := service.NewClienteService(repo, queueClient)

	err := svc.UpdateCliente(context.Background(), model.Clientes{})
	assert.NoError(t, err)
}

func TestClienteService_DeleteCliente(t *testing.T) {
	repo := &mockClienteRepository{}
	queueClient := createMockRabbitMQClient()

	svc := service.NewClienteService(repo, queueClient)

	err := svc.DeleteCliente(context.Background(), 1)
	assert.NoError(t, err)

	err = svc.DeleteCliente(context.Background(), 2)
	assert.Error(t, err)
	assert.EqualError(t, err, "falha ao deletar cliente")
}

func createAMQPConnection() (*amqp.Connection, error) {
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao RabbitMQ: %v", err)
	}

	return conn, nil
}
