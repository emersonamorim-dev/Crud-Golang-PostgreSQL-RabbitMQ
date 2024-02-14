package repository

import (
	"Crud-Golang-RabbitMQ/internal/model"
	"context"

	"gorm.io/gorm"
)

// ClientesRepository interface define os métodos que o repositório de clientes deve implementar.
type ClientesRepository interface {
	Create(ctx context.Context, c model.Clientes) error
	ListarTodos(ctx context.Context) ([]model.Clientes, error)
	GetByID(ctx context.Context, id uint) (model.Clientes, error)
	Update(ctx context.Context, c model.Clientes) error
	Delete(ctx context.Context, id uint) error
}

type gormClientesRepository struct {
	db *gorm.DB
}

// NewClientesRepository cria um novo repositório de clientes utilizando GORM.
func NewClientesRepository(db *gorm.DB) ClientesRepository {
	return &gormClientesRepository{db: db}
}

// Insere um novo cliente no BD utilizando GORM.
func (r *gormClientesRepository) Create(ctx context.Context, c model.Clientes) error {
	result := r.db.WithContext(ctx).Create(&c)
	return result.Error
}

// Lista todos os clientes disponíveis no banco de dados.
func (r *gormClientesRepository) ListarTodos(ctx context.Context) ([]model.Clientes, error) {
	var clientes []model.Clientes
	result := r.db.WithContext(ctx).Find(&clientes)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientes, nil
}

// GetByID busca um cliente pelo ID utilizando GORM.
func (r *gormClientesRepository) GetByID(ctx context.Context, id uint) (model.Clientes, error) {
	var c model.Clientes
	result := r.db.WithContext(ctx).First(&c, id)
	if result.Error != nil {
		return model.Clientes{}, result.Error
	}
	return c, nil
}

// Update atualiza um cliente existente utilizando GORM.
func (r *gormClientesRepository) Update(ctx context.Context, c model.Clientes) error {
	result := r.db.WithContext(ctx).Save(&c)
	return result.Error
}

// Delete remove um cliente pelo ID utilizando GORM.
func (r *gormClientesRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&model.Clientes{}, id)
	return result.Error
}
