package model

import "gorm.io/gorm"

type Gopher struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Job  string `json:"job" form:"job"`
}
