package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	ID   uint    `json:"id" gorm:"primary_key"`
	Name string  `json:"name"`
	Perc float64 `json:"perc"`
}

type CreateSkillInput struct {
	Name string  `json:"name" binding:"required"`
	Perc float64 `json:"perc" binding:"required"`
}

type UpdateSkillInput struct {
	Name string  `json:"name"`
	Perc float64 `json:"perc"`
}
