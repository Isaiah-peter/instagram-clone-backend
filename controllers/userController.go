package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Isaiah-peter/instagram-clone/database"
	"github.com/Isaiah-peter/instagram-clone/models"
	"github.com/Isaiah-peter/instagram-clone/utils"
	"github.com/gorilla/mux"
)

var (
	db = database.GetDB()
)

type ReturnUser struct {
	Id           uint
	UserName     string
	FullName     string
	Email        string
	ProfileImage string
}

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

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var user *models.User
	db.Find(&user)
	result := ReturnUser{
		Id:           user.ID,
		UserName:     user.UserName,
		FullName:     user.FullName,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}
	res, _ := json.Marshal(result)
	utils.Result(res, w, r)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var Id = mux.Vars(r)
	id, err := strconv.Atoi(Id["id"])
	if err != nil {
		panic(err)
	}
	var user *models.User
	db.Find(&user, id)
	result := ReturnUser{
		Id:           user.ID,
		UserName:     user.UserName,
		FullName:     user.FullName,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}
	res, _ := json.Marshal(result)
	utils.Result(res, w, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.Write([]byte("fail to decode"))
	}
	utils.UseToken(r)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	user1, db := models.GetUserById(id)
	if user.Email != "" {
		user1.Email = user.Email
	}
	if user.FullName != "" {
		user1.FullName = user.FullName
	}
	if user.ProfileImage != "" {
		user1.ProfileImage = user.ProfileImage
	}
	if user.UserName != "" {
		user1.UserName = user.UserName
	}
	if user.Password != "" {
		hash, _ := utils.Hashpassword(user.Password)
		user1.Password = hash
	}
	db.Save(user1)
	res, _ := json.Marshal(user1)
	w.Write(res)
}
