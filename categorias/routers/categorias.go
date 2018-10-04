package routers

import (
	"foro/categorias/controllers"

	"github.com/gorilla/mux"
)

func setCategoriaRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/categorias", controllers.GetCategorias).Methods("GET")
	//router.HandleFunc("/categorias", controllers.CreateCategoria).Methods("POST")
	//router.HandleFunc("/categorias/{id}", controllers.GetCategoriaById).Methods("GET")
	//router.HandleFunc("/categorias/{id}", controllers.DeleteCategoria).Methods("DELETE")
	router.HandleFunc("/temas", controllers.GetTemas).Methods("GET")
	router.HandleFunc("/temas", controllers.CreateTema).Methods("POST")
	router.HandleFunc("/mensajes", controllers.GetMensajes).Methods("GET")
	router.HandleFunc("/mensajes", controllers.CreateMensaje).Methods("POST")
	router.HandleFunc("/mensajes/{id}", controllers.GetMensajeByID).Methods("GET")
	router.HandleFunc("/mensajes/{id}", controllers.DeleteMensaje).Methods("DELETE")
	//router.HandleFunc("/categorias/{id}/Tema", controllers.GetTemaById).Methods("GET")
	return router
}
