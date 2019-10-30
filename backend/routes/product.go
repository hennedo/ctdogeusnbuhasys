package routes

import (
	"encoding/json"
	"net/http"

	"github.com/henne-/ctdogeusnbuhasys/backend/db"
	"github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/sirupsen/logrus"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	var slice []models.Product
	db.DB.Find(&[]models.Product{}).Scan(&slice)

	data, err := json.Marshal(slice)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}




