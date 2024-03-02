package service

import (
	"log"

	"github.com/sainath/todo-go-app/internal/user/model"
	"github.com/sainath/todo-go-app/internal/user/repository"
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
}
func HashPassword(password string) ([]byte, error) {
	hashcode, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return hashcode, err
}

func CreateUser(user model.User) error {
	result := db.Model(&user).Create(&user)

	return result.Error
}

func IsUserExist(userEmail string) bool {
	if err := db.Where("email = ?", userEmail).First(&model.User{}).Error; err == gorm.ErrRecordNotFound {
		return false
	} else {
		return true
	}
}
