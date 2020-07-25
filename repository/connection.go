package repository

import (
	"fmt"
	"gorillaMailer/common"
	"gorillaMailer/models"

	"github.com/jinzhu/gorm"

	// MySQL dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connectionString string = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", common.Env("USER"), common.Env("PASSWORD"), common.Env("IP"), common.Env("DATABASE"))

var db, err = gorm.Open("mysql", connectionString)

func init() {
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database!")

	db.AutoMigrate(&models.Subscriber{})
}
