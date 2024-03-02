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
	r.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", handler.PostUserHandler).Methods("POST")
	http.Handle("/", r)
	log.Println("Listening on port: 8081")
	http.ListenAndServe(":8081", nil)

}
