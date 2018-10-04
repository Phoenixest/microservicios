package main

import (
	"log"
	"net/http"

	"foro/categorias/common"
	"foro/categorias/routers"
)

// EPunto de entrada del programa
func main() {

	// Llama la lógica de inicio
	common.StartUp()
	// Obtiene el router mux
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
