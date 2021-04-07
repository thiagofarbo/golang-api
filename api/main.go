package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnviromentVariables()

	fmt.Println("Rodando API")
	r := router.Gerar()

	fmt.Printf("Database is running on por %d", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))

}
