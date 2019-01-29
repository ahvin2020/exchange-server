package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"io"
	"io/ioutil"
	"strings"
	"net/http"
	"html/template"
	//	"strconv"

	"exchange.com/exchange/helper"
	. "exchange.com/exchange/model"
	"exchange.com/exchange/system"
)

type UserCurrencyController struct {
	system.Controller
}

// home page route
func (controller *UserCurrencyController) Index_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)

	// With that kind of flags template can "figure out" what route is being rendered

	c.Env["Title"] = "Default Project - free Go website project template"

	widgets := helper.Parse(t, "user_currency_index", nil)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

func (controller *UserCurrencyController) Exchange_Rate_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	amount := r.FormValue("amount")
	fromCurrency := r.FormValue("from")
	toCurrency := r.FormValue("to")

	resp, err := http.Get("http://www.google.com/finance/converter?a=" + amount + "&from=" + fromCurrency + "&to=" + toCurrency)
	if err != nil {
		fmt.Println(err)
	}
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	split := strings.Split(string(body), "bld>")
	split = strings.Split(split[1], toCurrency)
	value := strings.Trim(split[0], " ")

	io.WriteString(w, value)
}

func (controller *UserCurrencyController) List_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	//	userId := c.URLParams["id"]

	dbSess := controller.GetDb(c).NewSession(nil)

	userCurrencyList := []*UserCurrency{}
	err := UserCurrencyModel.GetAll(dbSess, &userCurrencyList)

	if err != nil {
		fmt.Println(err.Error())
		//		session.AddFlash("Invalid Email or Password", "auth")
		//		return controller.SignIn(c, r)
	}

	//	for i := 0; i < len(userCurrencyList); i++ {
	//		fmt.Println(helper.Int64ToString(userCurrencyList[i].Id))
	//	}

	json.NewEncoder(w).Encode(&userCurrencyList)
}

func (controller *UserCurrencyController) Update_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		//		ReturnErr(w, err.Error())
		return
	}

	id := r.FormValue("id")
	userId := r.FormValue("user_id")
	buyAmount := r.FormValue("buy_amount")
	buyCurrency := r.FormValue("buy_currency")
	sellAmount := r.FormValue("sell_amount")
	sellCurrency := r.FormValue("sell_currency")

	sess := controller.GetDb(c).NewSession(nil)

	var updateId int64 = -1
	
	// TOOD: is this the best way to check?
	if id == "-1" {
		var res sql.Result = nil
		res, err = sess.UpdateBySql(`INSERT INTO user_currencies (user_id, sell_amount, sell_currency, buy_amount, buy_currency) VALUES ( ?, ?, ?, ?, ?)`,
			userId, sellAmount, sellCurrency, buyAmount, buyCurrency).Exec()
		
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		
		updateId, err = res.LastInsertId()
		
		io.WriteString(w, helper.Int64ToString(updateId))
	} else {
		_, err = sess.UpdateBySql(`UPDATE user_currencies SET sell_amount=?, sell_currency=?, buy_amount=?, buy_currency=? WHERE id=?;`,
			sellAmount, sellCurrency, buyAmount, buyCurrency, id).Exec()
		
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		
		io.WriteString(w, id)
	}
}

func (controller *UserCurrencyController) Delete_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		//		ReturnErr(w, err.Error())
		return
	}

	id := r.FormValue("id")

	sess := controller.GetDb(c).NewSession(nil)

	_, err = sess.UpdateBySql(`DELETE FROM user_currencies WHERE id = ?`, id).Exec()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	io.WriteString(w, "OK")
}
