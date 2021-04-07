package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	URL_DB = ""
	PORT   = 0
)

//Carrega as variveis de ambiente da aplicacap
func LoadEnviromentVariables() {

	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	PORT, error = strconv.Atoi(os.Getenv(("API_PORT")))
	if error != nil {
		PORT = 5000
	}

	URL_DB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

}
