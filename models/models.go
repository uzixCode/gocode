package models

type Model struct {
	Name   string  `json:"name"`
	IsGorm bool    `json:"is_gorm"`
	Fields []Field `json:"fields"`
}
