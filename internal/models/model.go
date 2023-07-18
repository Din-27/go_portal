package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id_user    int
	first_name string
	last_name  string
	email      string
	paswword   string
}