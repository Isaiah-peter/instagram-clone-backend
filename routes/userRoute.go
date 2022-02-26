package routes

import (
	"github.com/Isaiah-peter/instagram-clone/controllers"
	"github.com/gorilla/mux"
)

var UserRoute = func(route *mux.Router) {
	route.HandleFunc("/user", controllers.GetUser).Methods("GET")
	route.HandleFunc("/login", controllers.Login).Methods("POST")
	route.HandleFunc("/register", controllers.Register).Methods("POST")
	route.HandleFunc("/user/:id", controllers.GetUserById).Methods("GET")
}
