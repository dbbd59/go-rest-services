package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	ID   uint    `json:"id" gorm:"primary_key"`
	Name string  `json:"name"`
	Perc float64 `json:"perc"`
}
