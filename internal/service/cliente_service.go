package service

import (
	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/queue"
	"Crud-Golang-RabbitMQ/internal/repository"
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// ClienteService define a interface para o serviço de clientes
type ClienteService interface {
	CreateCliente(ctx context.Context, cliente model.Clientes) error
	ListarClientes(ctx context.Context) ([]model.Clientes, error)
	GetClienteByID(ctx context.Context, id uint) (model.Clientes, error)
	UpdateCliente(ctx context.Context, cliente model.Clientes) error
	DeleteCliente(ctx context.Context, id uint) error
}

// Implementei o suporte adicional para operações RabbitMQ
type service struct {
	repo        repository.ClientesRepository
	queueClient *queue.RabbitMQ
}


type QueueClient interface {
	PublishMessage(queueName string, message amqp.Publishing) error
}


func NewClienteService(repo repository.ClientesRepository, queueClient *queue.RabbitMQ) ClienteService {
	return &service{
		repo:        repo,
		queueClient: queueClient,
	}
}


// cria um novo cliente e publica uma mensagem no RabbitMQ
func (s *service) CreateCliente(ctx context.Context, cliente model.Clientes) error {
	// Insere o cliente no BD
	err := s.repo.Create(ctx, cliente)
	if err != nil {
		return fmt.Errorf("não foi possível inserir o cliente: %v", err)
	}

	// Serializa o cliente para JSON para publicar na fila
	clienteJSON, err := json.Marshal(cliente)
	if err != nil {
		return fmt.Errorf("não foi possível serializar o cliente: %v", err)
	}

	// Publica a mensagem na fila do RabbitMQ
	err = s.queueClient.Publish("clientesQueue", clienteJSON)
	if err != nil {
		return fmt.Errorf("falha ao publicar mensagem de criação de cliente: %v", err)
	}

	return nil
}

// ListarClientes retorna todos os clientes disponíveis.
func (s *service) ListarClientes(ctx context.Context) ([]model.Clientes, error) {
	return s.repo.ListarTodos(ctx)
}

// GetClienteByID busca um cliente pelo ID
func (s *service) GetClienteByID(ctx context.Context, id uint) (model.Clientes, error) {
	return s.repo.GetByID(ctx, uint(id))
}

// UpdateCliente atualiza um cliente existente
func (s *service) UpdateCliente(ctx context.Context, cliente model.Clientes) error {
	return s.repo.Update(ctx, cliente)
}

// DeleteCliente remove um cliente pelo ID
func (s *service) DeleteCliente(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, uint(id))
}

