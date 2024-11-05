package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var Config configStruct
var PathLevel map[string]map[string]int

type configStruct struct {
	Server   ServerConfig   `json:"server"`
	Postgres PostgresConfig `json:"postgres"`
	Jwt      JwtConfig      `json:"jwt"`
	Bcrypt   BcryptConfig   `json:"bcrypt"`
	User     UserConfig     `json:"user"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type PostgresConfig struct {
	Dsn string `json:"dsn"`
}

type JwtConfig struct {
	Secret string `json:"secret"`
	Exp    int64  `json:"exp"`
}

type BcryptConfig struct {
	Cost int `json:"cost"`
}

type UserConfig struct {
	Nickname  NicknameConfig  `json:"nickname"`
	InitAdmin InitAdminConfig `json:"init_admin"`
}

type NicknameConfig struct {
	RandMin int `json:"rand_min"`
	RandMax int `json:"rand_max"`
}

type InitAdminConfig struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InitConfig() {
	InitDefault()
	InitPathLevel()
}

func InitDefault() {
	var err error
	var file []byte
	var dir string
	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	parts := strings.Split(dir, string("/user-manager"))
	path := fmt.Sprintf("%s/user-manager/config/default.json", parts[0])
	file, err = os.ReadFile(path)
	err = json.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}
}

func InitPathLevel() {
	var err error
	var file []byte
	var dir string
	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	parts := strings.Split(dir, string("/user-manager"))
	path := fmt.Sprintf("%s/user-manager/config/pathLevel.json", parts[0])
	file, err = os.ReadFile(path)
	err = json.Unmarshal(file, &PathLevel)
	if err != nil {
		panic(err)
	}
}
