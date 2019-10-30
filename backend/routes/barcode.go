package routes

import (
	"encoding/json"
	"net/http"

	"github.com/henne-/ctdogeusnbuhasys/backend/db"
	"github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/sirupsen/logrus"
)

func ListBarcodes(w http.ResponseWriter, r *http.Request) {
	var slice []models.Barcode
	db.DB.Find(&[]models.Barcode{}).Scan(&slice)

	data, err := json.Marshal(slice)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	w.Write(data)}

func GetBarcode(w http.ResponseWriter, r *http.Request) {

}

func CreateBarcode(w http.ResponseWriter, r *http.Request) {

}

func UpdateBarcode(w http.ResponseWriter, r *http.Request) {

}

func GetEanBarcode(w http.ResponseWriter, r *http.Request) {

}


