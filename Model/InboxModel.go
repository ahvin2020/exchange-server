package model

import (
// "github.com/gocraft/dbr"
// "time"

// "exchange.com/exchange/helper"
// . "exchange.com/exchange/helper"
)

var InboxModel Inbox

type Inbox struct {
	Id int64 `json:"Id"`
	// Username    string    `json:"Username"`
	// Email       string    `json:"Email"`
	// Password    string    `json:"Password"`
	// FirstName   string    `json:"FirstName"`
	// LastName    string    `json:"LastName"`
	// Gender      string    `json:"Gender"`
	// Birthday    time.Time `json:"Birthday"`
	// CountryId   int16     `json:"CountryId"`
	// Bio         string    `json:"Bio"`
	// ProfilePic  string    `json:"ProfilePic"`
	// Token       string    `json:"Token"`
	// TokenExpiry time.Time `json:"TokenExpiry"`
	// CreateIp    string    `json:"CreateIp"`
	// Created     time.Time `json:"Created"`
}

// func (UserModel User) Edit(dbSess *dbr.Session, accessToken string, email string, firstName string, lastName string, gender string, birthday string, country string, bio string) error {
// 	_, err := dbSess.UpdateBySql("UPDATE users SET email=?, first_name=?, last_name=?, gender=?, birthday=?, country_id=?, bio=? WHERE token=? LIMIT 1;", email, firstName, lastName, gender, birthday, country, bio, accessToken).Exec()
// 	return err
// }
