package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Rutas para la entidad categoria
	router = setCategoriaRouters(router)
	return router
}
