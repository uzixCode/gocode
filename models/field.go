package models

type Field struct {
	Name     string `json:"name"`
	DataType string `json:"data_type"`
	JSONTag  string `json:"json_tag"`
	GormTag  string `json:"gorm_tag"`
}
