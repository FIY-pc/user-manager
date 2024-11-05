package model

import (
	"github.com/FIY-pc/user-manager/internal/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB

func InitPostgres() {
	var err error
	postgresDB, err = gorm.Open(postgres.Open(config.Config.Postgres.Dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// 数据库迁移
	err = postgresDB.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
	// 初始化管理员
	InitAdmin()
}

func InitAdmin() {
	if _, err := GetUser(config.Config.User.InitAdmin.Email); err != nil {
		admin := User{
			Nickname:   config.Config.User.InitAdmin.Nickname,
			Email:      config.Config.User.InitAdmin.Email,
			Password:   config.Config.User.InitAdmin.Password,
			Permission: 2,
		}
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}
		admin.Password = string(hashPassword)
		_, err = CreateUser(admin)
		if err != nil {
			panic(err.Error())
		}
	}
}
