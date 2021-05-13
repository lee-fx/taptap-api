package config

import (
	"encoding/json"
	"os"
	"strconv"
)

type Configuration struct {
	MYSQL_NAME string `json:"mysql_name"`
	MYSQL_PASSWORD string `json:"mysql_password"`
	MYSQL_HOST string `json:"mysql_host"`
	MYSQL_PORT string `json:"mysql_port"`
	MYSQL_APP_DB string `json:"mysql_app_db"`
	MYSQL_ADMIN_DB string `json:"mysql_admin_db"`

	REDIS_ADDR string `json:"redis_addr"`
	REDIS_PORT string `json:"redis_port"`
	REDIS_PASSWORD string `json:"redis_password"`
	REDIS_APP_DB string `json:"redis_app_db"`
	REDIS_ADMIN_DB string `json:"redis_admin_db"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetMysqlName() string {
	return configuration.MYSQL_NAME
}

func GetMysqlPassword() string {
	return configuration.MYSQL_PASSWORD
}

func GetMysqlHost() string {
	return configuration.MYSQL_HOST
}

func GetMysqlPort() string {
	return configuration.MYSQL_PORT
}

func GetMysqlAppDb() string {
	return configuration.MYSQL_APP_DB
}

func GetMysqlAdminDb() string {
	return configuration.MYSQL_ADMIN_DB
}



func GetRedisAddr() string {
	return configuration.REDIS_ADDR
}

func GetRedisPort() string {
	return configuration.REDIS_PORT
}

func GetRedisPassword() string {
	return configuration.REDIS_PASSWORD
}

func GetRedisAppDb() int {
	db, _ := strconv.Atoi(configuration.REDIS_APP_DB)
	return db
}

func GetRedisAdminDb() int {
	db, _ := strconv.Atoi(configuration.REDIS_ADMIN_DB)
	return db
}

