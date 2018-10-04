package controllers

import (
	"foro/categorias/models"
)

type (
	// Para Get - /categorias
	CategoriasResource struct {
		Data []models.Categoria `json:"data"`
	}

	// Para Get - /temas
	TemasResource struct {
		Data []models.Tema `json:"data"`
	}

	// Para Get - /mensajes
	MensajesResource struct {
		Data []models.Mensaje `json:"data"`
	}

	// Para Post/Put - /temas
	TemaResource struct {
		Data models.Tema `json:"data"`
	}

	// Para Post/Put - /mensajes
	MensajeResource struct {
		Data models.Mensaje `json:"data"`
	}
)
