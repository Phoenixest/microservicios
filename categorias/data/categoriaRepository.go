package data

import (
	"foro/categorias/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CategoriaRepository struct {
	C *mgo.Collection
}

func (r *CategoriaRepository) GetAllC() []models.Categoria {
	var categorias []models.Categoria
	iter := r.C.Find(nil).Iter()
	result := models.Categoria{}
	for iter.Next(&result) {
		categorias = append(categorias, result)
	}
	return categorias
}

func (r *CategoriaRepository) GetAllT() []models.Tema {
	var temas []models.Tema
	iter := r.C.Find(nil).Iter()
	result := models.Tema{}
	for iter.Next(&result) {
		temas = append(temas, result)
	}
	return temas
}

func (r *CategoriaRepository) GetAllM() []models.Mensaje {
	var mensajes []models.Mensaje
	iter := r.C.Find(nil).Iter()
	result := models.Mensaje{}
	for iter.Next(&result) {
		mensajes = append(mensajes, result)
	}
	return mensajes
}

func (r *CategoriaRepository) CreateT(tema *models.Tema) error {
	obj_id := bson.NewObjectId()
	tema.ID = obj_id
	err := r.C.Insert(&tema)
	return err
}

func (r *CategoriaRepository) CreateM(mensaje *models.Mensaje) error {
	obj_id := bson.NewObjectId()
	mensaje.ID = obj_id
	err := r.C.Insert(&mensaje)
	return err
}

func (r *CategoriaRepository) GetByID(id string) (mensaje models.Mensaje, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&mensaje)
	return
}

func (r *CategoriaRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
