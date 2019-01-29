package controller

import (
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	
	"exchange.com/exchange/system"
	"exchange.com/exchange/helper"
)

type PageController struct {
	system.Controller
}

func (controller *PageController) Index_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)
	 
	c.Env["title"] = "Index | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "page_index", c.Env)
	c.Env["content"] = template.HTML(widgets)
	
	return helper.Parse(t, "main", c.Env), http.StatusOK
}

func (controller *PageController) Not_Found_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)
	 
	c.Env["title"] = "Not Found | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "page_not_found", c.Env)
	c.Env["content"] = template.HTML(widgets)
	
	return helper.Parse(t, "main", c.Env), http.StatusOK
}