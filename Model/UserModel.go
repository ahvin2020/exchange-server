package model

import (
	"github.com/gocraft/dbr"
	"time"

	"exchange.com/exchange/helper"
	. "exchange.com/exchange/helper"
)

var UserModel User

type User struct {
	Id          int64     `json:"Id"`
	Username    string    `json:"Username"`
	Email       string    `json:"Email"`
	Password    string    `json:"Password"`
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Gender      string    `json:"Gender"`
	Birthday    time.Time `json:"Birthday"`
	CountryId   int16     `json:"CountryId"`
	Bio         string    `json:"Bio"`
	ProfilePic  string    `json:"ProfilePic"`
	Token       string    `json:"Token"`
	TokenExpiry time.Time `json:"TokenExpiry"`
	CreateIp    string    `json:"CreateIp"`
	Created     time.Time `json:"Created"`
}

func (UserModel User) Edit(dbSess *dbr.Session, accessToken string, email string, firstName string, lastName string, gender string, birthday string, country string, bio string) error {
	_, err := dbSess.UpdateBySql("UPDATE users SET email=?, first_name=?, last_name=?, gender=?, birthday=?, country_id=?, bio=? WHERE token=? LIMIT 1;", email, firstName, lastName, gender, birthday, country, bio, accessToken).Exec()
	return err
}

func (UserModel User) EditPasswordByToken(dbSess *dbr.Session, accessToken string, password string) error {
	_, err := dbSess.UpdateBySql("UPDATE users SET password=? WHERE token=? LIMIT 1;", password, accessToken).Exec()
	return err
}

func (UserModel User) EditPasswordByUserId(tx *dbr.Tx, userId int64, password string) error {
	_, err := tx.UpdateBySql("UPDATE users SET password=? WHERE id=? LIMIT 1;", password, userId).Exec()
	return err
}

func (UserModel User) Login(dbSess *dbr.Session, username string) (string, error) {
	// get access token and expiry date
	user := User{}
	err := dbSess.SelectBySql("SELECT id, token, token_expiry FROM users WHERE username=? LIMIT 1;", username).LoadStruct(&user)

	if err != nil {
		return "", err
	}

	accessToken := user.Token

	// get time diff
	currentTime := time.Now()
	if currentTime.Unix() > user.TokenExpiry.Unix() {
		accessToken = helper.RandStr(32)
	}

	// update token expiry
	_, err = dbSess.UpdateBySql("UPDATE users SET token=?, token_expiry=DATE_ADD(NOW(), INTERVAL ? HOUR) WHERE username=? LIMIT 1;", accessToken, LOGIN_TOKEN_HOUR_VALIDITY, username).Exec()

	if err != nil {
		return "", err
	} else {
		return accessToken, err
	}
}

func (UserModel User) GetPasswordByUsername(dbSess *dbr.Session, username string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, password FROM users WHERE username=? LIMIT 1;", username).LoadStruct(&user)

	return user, err
}

func (UserModel User) GetPasswordByAccessToken(dbSess *dbr.Session, accessToken string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, password FROM users WHERE token=? LIMIT 1;", accessToken).LoadStruct(&user)

	return user, err
}

// get the minimum user info
func (UserModel User) GetUserByAccessTokenMin(dbSess *dbr.Session, accessToken string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, first_name, last_name, token FROM users WHERE token=? AND token_expiry >= NOW() LIMIT 1;", accessToken).LoadStruct(&user)

	return user, err
}

// get the minimum user info
// func (UserModel User) GetUserByIdMin(dbSess *dbr.Session, id int64) (User, error) {
// 	user := User{}
// 	err := dbSess.SelectBySql("SELECT id, username, first_name, last_name, token FROM users WHERE id=? LIMIT 1;", id).LoadStruct(&user)

// 	return user, err
// }

// get user profile pic
func (UserModel User) GetProfilePicByAccessToken(dbSess *dbr.Session, accessToken string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, profile_pic, token FROM users WHERE token=? AND token_expiry >= NOW() LIMIT 1;", accessToken).LoadStruct(&user)

	return user, err
}

// get all user info
func (UserModel User) GetUserByAccessTokenFull(dbSess *dbr.Session, accessToken string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, email, first_name, last_name, gender, birthday, country_id, bio, profile_pic, token FROM users WHERE token=? AND token_expiry >= NOW() LIMIT 1;", accessToken).LoadStruct(&user)

	return user, err
}

func (UserModel User) GetUserByEmail(dbSess *dbr.Session, email string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, email, token FROM users WHERE email=? LIMIT 1;", email).LoadStruct(&user)

	return user, err
}

func (UserModel User) GetUserByEmailOrUsername(dbSess *dbr.Session, email string, username string) ([]*User, error) {
	userList := []*User{}
	_, err := dbSess.SelectBySql("SELECT id, username, email, token FROM users WHERE email=? OR username=? LIMIT 2;", email, username).LoadStructs(&userList)

	return userList, err
}

// get the minimum user info
func (UserModel User) GetUserByUsernameMin(dbSess *dbr.Session, username string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, first_name, last_name, token FROM users WHERE username=? LIMIT 1;", username).LoadStruct(&user)

	return user, err
}

// get all user info
func (UserModel User) GetUserByUsernameFull(dbSess *dbr.Session, username string) (User, error) {
	user := User{}
	err := dbSess.SelectBySql("SELECT id, username, email, first_name, last_name, gender, birthday, country_id, bio, profile_pic, token, created FROM users WHERE username=? LIMIT 1;", username).LoadStruct(&user)

	return user, err
}

func (UserModel User) Signup(dbSess *dbr.Session, email string, username string, hashPassword string, country string, ipAddress string) error {
	_, err := dbSess.UpdateBySql("INSERT INTO users (email, username, password, country_id, create_ip, created) VALUES (?, ?, ?, ?, ?, NOW());", email, username, hashPassword, country, ipAddress).Exec()
	return err
}

func (UserModel User) UpdateProfilePicByToken(tx *dbr.Tx, token string, profilePic string) error {
	_, err := tx.UpdateBySql("UPDATE users SET profile_pic=? WHERE token=? LIMIT 1;", profilePic, token).Exec()
	return err
}
