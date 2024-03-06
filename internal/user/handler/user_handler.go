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
	"github.com/sainath/todo-go-app/pkg/util"
	"golang.org/x/crypto/bcrypt"
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
	err := db.Model(&model.User{}).Preload("Tasks").First(&user, id).Error
	if err != nil {
		log.Fatalln(err)
	}
	userGetResponse.Id = user.ID
	userGetResponse.Email = user.Email
	userGetResponse.FirstName = user.FirstName
	userGetResponse.LastName = user.LastName
	userGetResponse.Tasks = user.Tasks
	json.NewEncoder(w).Encode(&userGetResponse)

}

func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	pathParameters := r.URL.Query()

	var user model.User
	id, _ := strconv.Atoi(pathParameters.Get("id"))
	db.First(&user, id)
	json.NewEncoder(w).Encode(user.Email)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserLogin
	var user model.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&userLogin)
	err := db.Model(&model.User{}).Where("email = ?", userLogin.Username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(&model.TokenResponse{Message: "No user found with email " + userLogin.Username, Token: " "})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLogin.Password)); err != nil {
		json.NewEncoder(w).Encode(&model.TokenResponse{Message: "Invalid Password! Try Again..! ", Token: " "})
		return
	}

	token, err := util.GenerateJwtToken(userLogin.Username)

	if err != nil {
		json.NewEncoder(w).Encode([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&model.TokenResponse{Message: "User Logedin Successfully!", Token: token})

}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
