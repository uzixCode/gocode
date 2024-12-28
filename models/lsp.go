package models

import (
	"gorm.io/gorm"
)

type Lsp struct {
	gorm.Model
	ID        string `json:"ID" gorm:"type:uuid;primaryKey"`
	Name      string `json:"name" gorm:"not null;default:'';type:varchar(255)"`
	Profile   string `json:"profile" gorm:"default:'';"`
	Email     string `json:"email" gorm:"not null;default:'';type:varchar(255)"`
	Wa        string `json:"wa" gorm:"not null;default:'';type:varchar(255)"`
	Type      string `json:"type" gorm:"not null;default:'';type:varchar(255)"`
	Tomp      string `json:"tomp" gorm:"not null;default:'';type:varchar(255)"`
	IsSuspend bool   `json:"is_suspend" gorm:"not null;default:false;"`
}
