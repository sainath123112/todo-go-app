package model

import (
	"github.com/sainath/todo-go-app/internal/task/model"
	"gorm.io/gorm"
)

type TasksArray interface {
}

type User struct {
	*gorm.Model
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"password_hash"`
	Tasks        []model.Task `json:"-"`
}

type TokenResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type UserPostResponse struct {
	Message string      `json:"message"`
	Data    UserDetails `json:"data"`
}

type UserGetResponse struct {
	Id        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Tasks     []model.Task `json:"tasks"`
}

type UserExistsMessage struct {
	Message string `json:"message"`
}
