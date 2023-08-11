package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Pgdb() (*sqlx.DB, error) {
	//menggunakan .env
	// host := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")

	// config := fmt.Sprintf("host=%s dbname=%s password=%s user=%s sslmode=disable",
	// host, dbName, password, user)

	//menggunakan yml
	host := viper.GetString("database.host")
	dbName := viper.GetString("database.name")
	port := viper.GetString("database.port")
	password := viper.GetString("database.pass")
	user := viper.GetString("database.user")

	config := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s sslmode=disable",
		host, dbName, port, password, user)

	return sqlx.Connect("postgres", config)
}
