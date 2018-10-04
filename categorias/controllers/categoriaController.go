package controllers

import (
	"encoding/json"
	"net/http"

	"foro/categorias/common"
	"foro/categorias/data"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Handler para peticion HTTP Get - "/categorias"
// Devuelve todas las categorias
func GetCategorias(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("categorias")
	repo := &data.CategoriaRepository{c}
	categorias := repo.GetAllC()
	j, err := json.Marshal(CategoriasResource{Data: categorias})
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Get - "/temas"
// Devuelve todas los temas
func GetTemas(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("temas")
	repo := &data.CategoriaRepository{c}
	temas := repo.GetAllT()
	j, err := json.Marshal(TemasResource{Data: temas})
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Get - "/mensajes"
// Devuelve todas los mensajes
func GetMensajes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("mensajes")
	repo := &data.CategoriaRepository{c}
	mensajes := repo.GetAllM()
	j, err := json.Marshal(MensajesResource{Data: mensajes})
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Post - "/temas"
// Inserta un nuevo tema
func CreateTema(w http.ResponseWriter, r *http.Request) {
	var dataResourse TemaResource
	// Decodifica el json de el tema entrante
	err := json.NewDecoder(r.Body).Decode(&dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Tema no valido", 500)
		return
	}
	tema := &dataResourse.Data

	// crea un nuevo contexto
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("temas")
	// Inserta un tema
	repo := &data.CategoriaRepository{c}
	repo.CreateT(tema)
	j, err := json.Marshal(dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Post - "/mensajes"
// Inserta un nuevo mensaje
func CreateMensaje(w http.ResponseWriter, r *http.Request) {
	var dataResourse MensajeResource
	// Decodifica el json de el mensaje entrante
	err := json.NewDecoder(r.Body).Decode(&dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Mensaje no valido", 500)
		return
	}
	mensaje := &dataResourse.Data

	// crea un nuevo contexto
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("mensajes")
	// Inserta un mensaje
	repo := &data.CategoriaRepository{c}
	repo.CreateM(mensaje)
	j, err := json.Marshal(dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Get - "/mensajes/{id}"
// Obtiene un mensaje por su ID
func GetMensajeByID(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del url
	vars := mux.Vars(r)
	id := vars["id"]

	// crea un nuevo contexto
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("mensajes")
	repo := &data.CategoriaRepository{c}

	// Obtiene el mensaje por el ID
	mensaje, err := repo.GetByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
			return
		}
	}

	j, err := json.Marshal(MensajeResource{Data: mensaje})
	if err != nil {
		common.DisplayAppError(w, err, "Ha ocurrido un error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler para peticion HTTP Delete - "/mensajes/{id}"
// Borra un mensaje por su ID
func DeleteMensaje(rw http.ResponseWriter, req *http.Request) {
	// Obtiene el ID de la url
	vars := mux.Vars(req)
	id := vars["id"]

	// Crea un nuevo contexto
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("mensajes")
	repo := &data.CategoriaRepository{c}

	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(rw, err, "Ha ocurrido un error", 500)
			return
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}
