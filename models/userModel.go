package models

import (
	"time"

	"github.com/Isaiah-peter/instagram-clone/database"
	"github.com/Isaiah-peter/instagram-clone/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = &gorm.DB{}

type User struct {
	gorm.Model
	UserName     string `json:"user_name"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image"`
}

type Token struct {
	UserID int64
	Email  string
	jwt.StandardClaims
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

func GetUserById(id int) (*User, *gorm.DB) {
	var user *User
	d := db.Where("ID = ?", id).Find(&user)
	return user, d
}

func FindOne(email string, password string) map[string]interface{} {
	var user *User

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp

	}

	expireAt := time.Now().Add(time.Minute * 1000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &Token{
		UserID: int64(user.ID),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}
