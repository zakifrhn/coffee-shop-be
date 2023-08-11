package main

import (
	"inter/internal/routers"
	"inter/pkg"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func init() {
	//set package viper in here
	//govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("env.config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database, err := pkg.Pgdb()

	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
