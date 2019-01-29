package model

import (
	//	"fmt"
	"github.com/gocraft/dbr"
	"time"

	. "exchange.com/exchange/helper"
)

var ResetPasswordTicketModel ResetPasswordTicket

type ResetPasswordTicket struct {
	UserId      int64    `json:"UserId"`
	TokenHash   string    `json:"TokenHash"`
	TokenExpiry time.Time `json:"TokenExpiry"`
	TokenUsed   int8      `json:"TokenUsed"`
}

func (ResetPasswordTicketModel ResetPasswordTicket) Save(dbSess *dbr.Session, userId int64, tokenHash string) error {
	_, err := dbSess.UpdateBySql("INSERT INTO reset_password_tickets (user_id, token_hash, token_expiry, token_used) VALUES (?, ?, DATE_ADD(NOW(), INTERVAL ? HOUR), 0) ON DUPLICATE KEY UPDATE token_hash=?, token_expiry=DATE_ADD(NOW(), INTERVAL ? HOUR), token_used=0;", userId, tokenHash, PASSWORD_TOKEN_HOUR_VALIDITY, tokenHash, PASSWORD_TOKEN_HOUR_VALIDITY).Exec()
	return err
}

func (ResetPasswordTicketModel ResetPasswordTicket) InvalidateToken(tx *dbr.Tx, tokenHash string) error {
	_, err := tx.UpdateBySql("UPDATE reset_password_tickets SET token_used=1 WHERE token_hash=? LIMIT 1;", tokenHash).Exec()
	return err
}


func (ResetPasswordTicketModel ResetPasswordTicket) GetTicketByToken(dbSess *dbr.Session, tokenHash string) (ResetPasswordTicket, error) {
	ticket := ResetPasswordTicket{}
	err := dbSess.SelectBySql("SELECT user_id FROM reset_password_tickets WHERE token_hash=? AND token_expiry >= NOW() AND token_used=0 LIMIT 1;", tokenHash).LoadStruct(&ticket)

	return ticket, err
}
