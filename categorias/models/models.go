package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Categoria struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Nombre      string        `json:"nombre"`
		Descripcion string        `json:"descripcion"`
		NumMensajes int           `json:"numero_mensajes"`
		NumTemas    int           `json:"numero_temas"`
	}

	Tema struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		User      string        `json:"user"`
		Asunto    string        `json:"asunto"`
		NumRes    int           `json:"numero_respuestas"`
		NumVistas int           `json:"numero_vistas"`
	}

	Mensaje struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		User      string        `json:"user"`
		Contenido string        `json:"contenido"`
		Fecha     time.Time     `json:"fecha,omitempty"`
	}
)
