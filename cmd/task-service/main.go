package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/internal/task/handler"
)

type Config struct {
	TaskServiceStruct TaskService `yaml:"task_service"`
}
type TaskService struct {
	Port string `yaml:"port"`
}

func main() {

	data, err := os.ReadFile("configs/task/config.yaml")

	if err != nil {
		log.Fatal("Error while reading file")
	}
	var config Config

	yaml.Unmarshal(data, &config)

	r := mux.NewRouter()
	r.Use(handler.AuthenticationMiddleware)
	r.HandleFunc("/tasks/{userid}", handler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{userid}/{id}", handler.GetSingleTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{userid}", handler.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{userid}/{id}", handler.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{userid}/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.Handle("/", r)
	log.Println("Listening on port: " + config.TaskServiceStruct.Port)

	http.ListenAndServe(":"+config.TaskServiceStruct.Port, nil)

}
