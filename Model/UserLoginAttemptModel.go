package model

import (
	"github.com/gocraft/dbr"
	"time"

	. "exchange.com/exchange/helper"
)

var UserLoginAttemptModel UserLoginAttempt

type UserLoginAttempt struct {
	UserId       int64     `json:"UserId"`
	Tries        int       `json:"Tries"`
	BlockExpires time.Time `json:"BlockExpires"`
}

func (UserLoginAttemptModel UserLoginAttempt) GetLoginAttempt(dbSess *dbr.Session, userId int64) (UserLoginAttempt, error) {
	userLoginAttempt := UserLoginAttempt{}
	err := dbSess.SelectBySql("SELECT * FROM user_login_attempts WHERE user_id=?;", userId).LoadStruct(&userLoginAttempt)

	return userLoginAttempt, err
}

func (UserLoginAttemptModel UserLoginAttempt) UpdateLoginAttempt(dbSess *dbr.Session, userId int64, tries int) error {
	var err error

	// reached number of invalid tries
	if tries >= INVALID_PASSWORD_TRIES {
		_, err = dbSess.UpdateBySql("INSERT INTO user_login_attempts(user_id, tries, block_expires) VALUES (?, ?, DATE_ADD(NOW(), INTERVAL ? MINUTE)) ON DUPLICATE KEY UPDATE tries=?, block_expires=DATE_ADD(NOW(), INTERVAL ? MINUTE);", userId, tries, INVALID_PASSWORD_BLOCK_MINUTES, tries, INVALID_PASSWORD_BLOCK_MINUTES).Exec()
	} else { // haven't reach
		_, err = dbSess.UpdateBySql("INSERT INTO user_login_attempts(user_id, tries, block_expires) VALUES (?, ?, NOW()) ON DUPLICATE KEY UPDATE tries=?;", userId, tries, tries).Exec()
	}

	return err
}
