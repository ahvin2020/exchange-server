package controller

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"net/http"
	//	"fmt"

	"exchange.com/exchange/helper"
	. "exchange.com/exchange/helper"
	. "exchange.com/exchange/model"
	"exchange.com/exchange/system"
)

type CountryController struct {
	system.Controller
}

// not in use at the moment
func (controller *CountryController) List_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	dbSess := controller.GetDb(c).NewSession(nil)

	countryList, err := CountryModel.GetAll(dbSess)

	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	json.NewEncoder(w).Encode(&countryList)
}
