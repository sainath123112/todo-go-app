package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/internal/user/handler"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello from user Service")
}

func main() {

	r := mux.NewRouter()
	//r.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/username", handler.GetUsernameHandler).Methods("GET")
	r.HandleFunc("/register", handler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", handler.AuthenticationMiddleware(handler.GetUserHandler))
	http.Handle("/", r)
	log.Println("Listening on port: 8081")
	http.ListenAndServe(":8081", nil)

}
