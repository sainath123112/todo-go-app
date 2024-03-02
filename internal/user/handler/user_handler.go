package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/internal/user/model"
	"github.com/sainath/todo-go-app/internal/user/repository"
	"github.com/sainath/todo-go-app/internal/user/service"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = repository.DbUserConfiguration()
	if err != nil {
		log.Fatal("Unable to connect User database")
	}
	db.AutoMigrate(&model.User{})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var userGetResponse model.UserGetResponse
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db.Model(&model.User{}).Preload("Tasks").First(&user, id)
	userGetResponse.Id = user.ID
	userGetResponse.Email = user.Email
	userGetResponse.FirstName = user.FirstName
	userGetResponse.LastName = user.LastName
	userGetResponse.Tasks = user.Tasks
	json.NewEncoder(w).Encode(&userGetResponse)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var userDetails model.UserDetails
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	if service.IsUserExist(user.Email) {
		json.NewEncoder(w).Encode(&model.UserExistsMessage{Message: "User Already Exists! Try with different email..!"})
	} else {
		hased_password, err := service.HashPassword(user.PasswordHash)
		if err != nil {
			log.Fatalln("Unable to bcrypt the password")
		}
		user.PasswordHash = string(hased_password)

		if err = service.CreateUser(user); err != nil {
			log.Fatalln("Unable to create User")
		} else {
			db.Model(&user).Where("email = ?", user.Email).First(&userDetails)
			json.NewEncoder(w).Encode(&model.UserPostResponse{Message: "User Created Successfully!", Data: userDetails})
		}
	}

}
