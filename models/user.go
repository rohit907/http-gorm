package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primary_key": not null`
	Name  string `json:"name"`
	Email string `json:"email"`
}
