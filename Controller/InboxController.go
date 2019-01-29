package controller

import (
	// "encoding/json"
	// "github.com/gorilla/websocket"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	//	"fmt"

	"exchange.com/exchange/helper"
	// . "exchange.com/exchange/helper"
	// . "exchange.com/exchange/model"
	"exchange.com/exchange/system"
)

type InboxController struct {
	system.Controller
}

// not in use at the moment
func (controller *InboxController) Chat_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	// dbSess := controller.GetDb(c).NewSession(nil)

	listingId := c.URLParams["listing-id"]

	// countryList, err := CountryModel.GetAll(dbSess)

	//if err != nil {
	//	helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
	//	return
	//}

	// json.NewEncoder(w).Encode(&countryList)

	c.Env["title"] = "Chat | Exchange"

	helper.RetrieveFlashes(session, c)

	c.Env["listingId"] = listingId

	widgets := helper.Parse(t, "inbox_chat", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}
