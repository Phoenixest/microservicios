package controllers

import (
	"foro/categorias/common"

	mgo "gopkg.in/mgo.v2"
)

// Estructura usada para mantener el contexto de las peticiones HTTP
type Context struct {
	MongoSession *mgo.Session
}

// Cierra la mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Retorna una mgo.collection para el nombre dado
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// Crea un nuevo obejo de contexto para cada petición HTTP
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
