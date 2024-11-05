package model

import (
	"errors"
	"github.com/FIY-pc/user-manager/internal/tools"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Nickname   string `json:"name" `
	Password   string `json:"password"`
	Email      string `json:"email" gorm:"unique"`
	Permission int    `json:"permission"`
}

func GetUser(email string) (User, error) {
	var user User
	var result *gorm.DB
	if email != "" {
		result = postgresDB.Where("email = ?", email).First(&user)
	} else {
		return user, errors.New("get user failed due to invalid params")
	}
	return user, result.Error
}

func UpdateUser(user User) (User, error) {
	var result *gorm.DB
	if user.Email != "" {
		result = postgresDB.Where("email = ?", user.Email).Updates(user)
	} else {
		return user, errors.New("update user failed due to invalid params")
	}
	return user, result.Error
}

func DeleteUser(email string) (User, error) {
	var result *gorm.DB
	if email != "" {
		result = postgresDB.Where("email =?", email).Delete(&User{})
	} else {
		return User{}, errors.New("delete user failed due to invalid params")
	}
	return User{}, result.Error
}

func CreateUser(user User) (User, error) {
	var result *gorm.DB
	// 非空检查
	if user.Email == "" {
		return User{}, errors.New("create user failed due to invalid params")
	}
	if user.Password == "" {
		return User{}, errors.New("create user failed due to invalid params")
	}
	// 格式检查

	if user.Nickname == "" {
		user.Nickname = tools.GenerateRandName()
	}
	result = postgresDB.Create(&user)
	log.Println(result)
	return user, result.Error
}
