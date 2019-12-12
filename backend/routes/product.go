package routes

import (
	"encoding/json"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	var product models.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	product.ID = nil
	product.CreatedAt = nil
	product.UpdatedAt = nil

	db.DB.Save(&product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}




