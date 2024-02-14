package database

import (
	"Crud-Golang-RabbitMQ/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// representa uma conexão com BD
type DB struct {
	Conn *gorm.DB
}

// NewDB inicializa a conexão com o BD usando GORM
func NewDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &DB{Conn: db}, nil
}
