package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/internal/task/model"
	"github.com/sainath/todo-go-app/internal/task/repository"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type ErrorStruct struct {
	Message      string `json:"message"`
	ErrorMessage error  `json:"error"`
}

func init() {
	db, err = repository.DbCongiguration()
	if err != nil {
		log.Fatal("Unable to connect DB")
	}
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userid"])
	var taskArray []model.TaskResponse

	if result := db.Model(&model.Task{}).Where("user_id = ?", userId).Find(&taskArray); len(taskArray) > 0 {
		json.NewEncoder(w).Encode(&taskArray)
	} else {
		json.NewEncoder(w).Encode(&ErrorStruct{Message: "No records found for this user", ErrorMessage: result.Error})
	}

}

func GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskResponse model.TaskResponse
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	userId, _ := strconv.Atoi(vars["userid"])
	err := db.Model(&model.Task{}).Where("id = ? and user_id = ?", id, userId).First(&taskResponse, id).Error
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(&ErrorStruct{Message: "No task found for task id: " + vars["id"], ErrorMessage: err})
		return
	}
	json.NewEncoder(w).Encode(&taskResponse)

}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userid"])
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&task)
	task.UserId = userId
	db.Create(&task)
	json.NewEncoder(w).Encode(&task)

}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db.First(&task, id)
	json.NewDecoder(r.Body).Decode(&task)
	db.Save(&task)
	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if result := db.Unscoped().Delete(&task, id); result.Error != nil {
		json.NewEncoder(w).Encode(&ErrorStruct{Message: "Unable to delete task", ErrorMessage: result.Error})
	}

}
