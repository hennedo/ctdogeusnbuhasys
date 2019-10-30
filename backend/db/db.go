package db

import (
	"github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func Load() {
	logrus.Info("Loading")
	db, err := gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		logrus.Fatal(err)
	}
	
	db.AutoMigrate(&models.User{}, &models.Barcode{}, &models.Log{}, &models.Category{}, &models.Product{})

	DB = db
}

func Close() {
	DB.Close()
}
