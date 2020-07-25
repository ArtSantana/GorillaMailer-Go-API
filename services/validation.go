package services

import (
	"encoding/json"
	"fmt"
	"gorillaMailer/models"

	"github.com/jinzhu/gorm"
)

// ValidationExist Service
func ValidationExist(subscriber *gorm.DB) (bool, []byte) {
	result, err := json.Marshal(subscriber.Value)
	var subscriberModel models.Subscriber

	if err != nil {
		return false, result
	}
	json.Unmarshal([]byte(result), &subscriberModel)

	if subscriberModel.ID < 1 {
		fmt.Println(subscriberModel)
		return false, result
	}

	return true, result
}
