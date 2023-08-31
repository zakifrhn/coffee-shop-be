package main

import (
	"inter/internal/routers"
	"inter/pkg"
	"log"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func init() {

	govalidator.SetFieldsRequiredByDefault(true)

	//set package viper in here
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

//! migration database
//? migration sql (add the table)
//! after dockerfile
//migrate -path ./migrations/ -database "postgresql://postgres:Fazztrak2023@localhost:5440/dbcoffee?sslmode=disable&search_path=coffeshop" -verbose up

//! after dockercompose
//migrate -path ./migrations/ -database "postgresql://postgres:Fazztrak2023@localhost:5445/dbcoffee?sslmode=disable&search_path=coffeshop" -verbose up

//! before docker
//migrate -path ./migrations/ -database "postgresql://postgres:Fazztrak2023@localhost/webgolang?sslmode=disable&search_path=coffeshop" -verbose up

//?migration sql(delete the table)
//migrate -path ./migrations/ -database "postgresql://postgres:Fazztrak2023@localhost/webgolang?sslmode=disable&search_path=coffeshop" -verbose down
