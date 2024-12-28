package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               *string `json:"ID" gorm:"primaryKey;type:uuid;"`
	Emaul            *string `json:"email" gorm:"unique;"`
	Password         *string `json:"-" gorm:"default:'';type:varchar(255);"`
	FirstName        *string `json:"first_name" gorm:"default:'';type:varchar(255);"`
	LastName         *string `json:"last_name" gorm:"default:'';type:varchar(255);"`
	Address          *string `json:"address" gorm:"default:'';type:varchar(255);"`
	Nik              *string `json:"nik" gorm:"default:'';type:varchar(255);"`
	PhoneNumber      *string `json:"phone_number" gorm:"default:'';type:varchar(255);"`
	Type             *string `json:"type" gorm:"default:'';type:varchar(255);"`
	IsConfirmed      *bool   `json:"is_confirmed" gorm:"default:true;"`
	ConfirmationCode *string `json:"-" gorm:"type:varchar(255);"`
}
