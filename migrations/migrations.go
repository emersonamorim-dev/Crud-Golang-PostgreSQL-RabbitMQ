package migrations

import (
	"Crud-Golang-RabbitMQ/internal/model"

	"gorm.io/gorm"
)

// representa uma migração para criar a tabela de clientes.
type CreateClientesTable struct{}

// Migrate aplica a migração para criar a tabela de clientes.
func (m *CreateClientesTable) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Clientes{})
	if err != nil {
		return err
	}
	return nil
}

// Rollback reverte a migração para criar a tabela de clientes.
func (m *CreateClientesTable) Rollback(db *gorm.DB) error {
	err := db.Migrator().DropTable("clientes")
	if err != nil {
		return err
	}
	return nil
}
