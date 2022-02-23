package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaiah-peter/instagram-clone/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var message string
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "fail to translate json"
		w.Write([]byte(message))
		return
	}
	user.CreateUser()
	w.WriteHeader(http.StatusOK)
	message = "welcome fam"
	w.Write([]byte(message))
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u := models.FindOne(newUser.Email, newUser.Password)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(u)
}
