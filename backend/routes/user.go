package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/henne-/ctdogeusnbuhasys/backend/db"
	"github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/sirupsen/logrus"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var slice []models.User
	db.DB.Find(&[]models.User{}).Scan(&slice)

	data, err := json.Marshal(slice)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	w.Write(data)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	var user models.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	user.ID = nil
	user.CreatedAt = nil
	user.UpdatedAt = nil
	
	db.DB.Save(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
