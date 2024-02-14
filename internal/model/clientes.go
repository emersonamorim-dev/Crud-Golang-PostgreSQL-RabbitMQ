package model

import (
	"time"
)

// Clientes representa a camada de regra de negócios.
// swagger:model Clientes
type Clientes struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
	Nome       string     `gorm:"type:varchar(100);not null" json:"nome"`
	Sobrenome  string     `gorm:"type:varchar(100);not null" json:"sobrenome"`
	Contato    string     `gorm:"type:varchar(100);not null" json:"contato"`
	Endereco   string     `gorm:"type:varchar(200);not null" json:"endereco"`
	Nascimento string     `gorm:"type:date;not null" json:"nascimento"`
	CPF        string     `gorm:"type:varchar(11);not null;unique" json:"cpf"`
}
