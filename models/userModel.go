package models

import (
	"github.com/Isaiah-peter/instagram-clone/database"
	"github.com/Isaiah-peter/instagram-clone/utils"
	"gorm.io/gorm"
)

var db = &gorm.DB{}

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	Hashpassword, _ := utils.Hashpassword(u.Password)
	u.Password = Hashpassword
	db.Create(u)
	return u
}
