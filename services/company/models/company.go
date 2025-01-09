package models

import (
	"github.com/google/uuid"
)


type Company struct {
	ID     		uuid.UUID	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      	string		`gorm:"not null"`
}

func (Company) TableName() string {
	return "companies"
}