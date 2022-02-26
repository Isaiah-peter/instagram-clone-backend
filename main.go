package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Isaiah-peter/instagram-clone/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	route := mux.NewRouter()
	route.Handle("/", handlers.MethodHandler{})
	route.HandleFunc("/welcome", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("welcome"))
	})
	route.HandleFunc("/test/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)
		num, _ := strconv.Atoi(id["id"])
		str := strconv.Itoa(num)
		rw.Write([]byte("testing to get id" + " " + str))
	})
	routes.UserRoute(route)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(route)))
}
