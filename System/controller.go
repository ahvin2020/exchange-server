package system

import (
	"bytes"
	"html/template"

	"github.com/gocraft/dbr"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
)

// option
type Option struct {
	Value, Id, Text string
	Selected        string
}

type Controller struct {
}

func (controller *Controller) GetSession(c web.C) *sessions.Session {
	return c.Env["Session"].(*sessions.Session)
}

func (controller *Controller) GetTemplate(c web.C) *template.Template {
	return c.Env["Template"].(*template.Template)
}

func (controller *Controller) GetDb(c web.C) *dbr.Connection {
	return c.Env["Db"].(*dbr.Connection)
}

func (controller *Controller) Parse(t *template.Template, name string, data interface{}) string {
	var doc bytes.Buffer
	t.ExecuteTemplate(&doc, name, data)
	return doc.String()
}
