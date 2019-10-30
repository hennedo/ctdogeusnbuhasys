package routes

import (
	"encoding/json"
	"net/http"

	"github.com/henne-/ctdogeusnbuhasys/backend/db"
	"github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/sirupsen/logrus"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
	var slice []models.Category
	db.DB.Find(&[]models.Category{}).Scan(&slice)

	data, err := json.Marshal(slice)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	w.Write(data)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

}

func GetCategory(w http.ResponseWriter, r *http.Request) {

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

}
