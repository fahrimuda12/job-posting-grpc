package models

import (
	"time"

	"github.com/google/uuid"
)


type Job struct {
	ID     		uuid.UUID	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CompanyID  	uuid.UUID	`gorm:"type:uuid;not null"`
	Title      	string		`gorm:"not null type:text"`
	Description string		`gorm:"not null; type:text"`
	CreatedAt 	time.Time	`gorm:"autoCreateTime"`
	Company		Company		`gorm:"foreignKey:CompanyID"`

}



func (Job) TableName() string {
	return "jobs"
}