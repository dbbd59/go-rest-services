package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Date        string `json:"date"`
	Company     string `json:"company"`
	Link        string `json:"link"`
	LinkDisplay string `json:"linkDisplay"`
	JobTitle    string `json:"jobTitle"`
	Descr       string `json:"descr"`
}

type CreateJobInput struct {
	Date        string `json:"date" binding:"required"`
	Company     string `json:"company" binding:"required"`
	Link        string `json:"link"`
	LinkDisplay string `json:"linkDisplay"`
	JobTitle    string `json:"jobTitle" binding:"required"`
	Descr       string `json:"descr"`
}

type UpdateJobInput struct {
	Date        string `json:"date"`
	Company     string `json:"company"`
	Link        string `json:"link"`
	LinkDisplay string `json:"linkDisplay"`
	JobTitle    string `json:"jobTitle"`
	Descr       string `json:"descr"`
}
