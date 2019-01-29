package model

import (
	"github.com/gocraft/dbr"
	//	"fmt"
	//	"log"

	//	_ "github.com/go-sql-driver/mysql"
)

var UserCurrencyModel UserCurrency

type UserCurrency struct {
	Id           int64   `json:"Id"`
	UserId       int64   `json:"UserId"`
	SellAmount   float64 `json:"SellAmount"`
	SellCurrency string  `json:"SellCurrency"`
	BuyAmount    float64 `json:"BuyAmount"`
	BuyCurrency  string  `json:"BuyCurrency"`
}

func (UserCurrencyModel UserCurrency) GetAll(dbSess *dbr.Session, userCurrencyList *[]*UserCurrency) error {
	_, err := dbSess.SelectBySql("SELECT * FROM user_currencies;").
		LoadStructs(userCurrencyList)
	return err
}
