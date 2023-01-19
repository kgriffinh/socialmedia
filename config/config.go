package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWT_KEY                  string = ""
	CLOUDINARY_CLOUD_NAME    string = ""
	CLOUDINARY_API_KEY       string = ""
	CLOUDINARY_API_SECRET    string = ""
	CLOUDINARY_UPLOAD_FOLDER string = ""
)

type AppConfig struct {
	DBUser                   string
	DBPass                   string
	DBHost                   string
	DBPort                   int
	DBName                   string
	jwtKey                   string
	CLOUDINARY_CLOUD_NAME    string
	CLOUDINARY_API_KEY       string
	CLOUDINARY_API_SECRET    string
	CLOUDINARY_UPLOAD_FOLDER string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}

	JWT_KEY = app.jwtKey
	CLOUDINARY_CLOUD_NAME = app.CLOUDINARY_CLOUD_NAME
	CLOUDINARY_API_KEY = app.CLOUDINARY_API_KEY
	CLOUDINARY_API_SECRET = app.CLOUDINARY_API_SECRET
	CLOUDINARY_UPLOAD_FOLDER = app.CLOUDINARY_UPLOAD_FOLDER
	return &app
}
