package repository

import (
	"gorillaMailer/models"

	"github.com/jinzhu/gorm"
)

// GetSubscribers repository
func GetSubscribers() *gorm.DB {
	var subscribers []models.Subscriber

	result := db.Find(&subscribers)

	return result
}

// CreateSubscriber repository
func CreateSubscriber(subscriber *models.Subscriber) (bool, *models.Subscriber) {
	db.Create(&subscriber)
	created := db.NewRecord(subscriber)

	if !created {
		return true, subscriber
	}

	return false, subscriber
}

// GetUniqueSubscriber repository
func GetUniqueSubscriber(id int64) *gorm.DB {
	var subscriber models.Subscriber

	finded := db.Where("id = ?", id).First(&subscriber)

	return finded
}
