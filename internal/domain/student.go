package domain

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name        string
	RollNo      int
	Age         int
	Department  string
	PhoneNumber string
	ImageUrl      string
}
