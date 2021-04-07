package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

//Vai retornar as rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigureRoutes(r)
}
