package models

import "github.com/google/uuid"

type Job struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CompanyID   uuid.UUID `gorm:"type:uuid;not null"`
	Title      	string		`gorm:"not null type:text"`
	Description string		`gorm:"not null; type:text"`
}

func (Job) TableName() string {
	return "jobs"
}