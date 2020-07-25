package models

import (
	"github.com/jinzhu/gorm"
)

// Subscriber model
type Subscriber struct {
	gorm.Model
	Name  string
	Email string
	Phone string
}
