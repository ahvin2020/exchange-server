package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gocraft/dbr"
	"github.com/zenazn/goji/web"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2-unstable"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"exchange.com/exchange/helper"
	. "exchange.com/exchange/helper"
	. "exchange.com/exchange/model"
	"exchange.com/exchange/system"
)

type UserController struct {
	system.Controller
}

// forgot password page route
func (controller *UserController) Forgot_Password_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	// no need to reset password if user already login
	if session.Values["AccessToken"] != nil {
		return "/", http.StatusSeeOther
	}

	c.Env["title"] = "Forgot Password | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_forgot_password", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// forgot password action
func (controller *UserController) Forgot_Password_Action(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	}

	// get post value
	email := r.FormValue("email")

	// trim all white spaces
	email = strings.TrimSpace(email)

	isValidEmail := govalidator.IsEmail(email)

	// invalid input?
	if isValidEmail == false {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	}

	c.Env["email"] = email

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// get user
	user, err := UserModel.GetUserByEmail(dbSess, email)

	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	} else if user.Id == 0 { // user not found?
		session.AddFlash("Email is not in use", FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	}

	// generate a random password
	token := helper.RandStr(PASSWORD_TOKEN_LENGTH)
	tokenHash := helper.GenerateHash(token)

	err = ResetPasswordTicketModel.Save(dbSess, user.Id, tokenHash)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	}

	resetLink := helper.BaseUrl(r) + "/reset-password/" + token
	emailVar := map[string]string{
		"reset_link": resetLink,
	}

	emailBody := helper.Parse(t, "reset_password_email", emailVar)

	m := gomail.NewMessage()
	m.SetHeader("From", SUPPORT_EMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset your password")
	m.SetBody("text/html", emailBody)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "<email>", "<password>")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Forgot_Password_Page(c, r)
	}

	// send this password to the user
	session.AddFlash("A reset password link has been sent to your inbox..", FLASH_SUCCESS)

	return "/", http.StatusSeeOther
}

// login page route
func (controller *UserController) Login_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	// no need to login if user already login
	if session.Values["AccessToken"] != nil {
		return "/", http.StatusSeeOther
	}

	c.Env["title"] = "Login | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_login", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// login
func (controller *UserController) Login_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	// get post value
	username := r.FormValue("username")
	password := r.FormValue("password")

	// trim all white spaces
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// invalid input?
	if username == "" || password == "" {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	// save the data just in case fail, can auto populate form
	c.Env["username"] = username

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// get user
	user, err := UserModel.GetPasswordByUsername(dbSess, username)
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Login_Page(c, r)
	} else if user.Id == 0 { // user not found?
		session.AddFlash("User not found", FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	userLoginAttempt, err := UserLoginAttemptModel.GetLoginAttempt(dbSess, user.Id)
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	// there was a fail attempt
	if userLoginAttempt.UserId > 0 {
		// check whether expiry time has passed
		currentTime := time.Now()
		if currentTime.Unix() <= userLoginAttempt.BlockExpires.Unix() {
			session.AddFlash(INVALID_PASSWORD_BLOCK_MSG, FLASH_ERROR)
			return controller.Login_Page(c, r)
		}
	}

	// check whether password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// wrong password!
	if err != nil {
		// increment invalid login attempt by 1
		err2 := UserLoginAttemptModel.UpdateLoginAttempt(dbSess, user.Id, userLoginAttempt.Tries+1)
		if err2 != nil { // there's an error?
			session.AddFlash(err2.Error(), FLASH_ERROR)
			return controller.Login_Page(c, r)
		}

		session.AddFlash("Wrong password", FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	// is correct, retrieve access token
	accessToken, err := UserModel.Login(dbSess, username)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Login_Page(c, r)
	}

	// reset invalid login attempt
	if userLoginAttempt.UserId > 0 {
		err = UserLoginAttemptModel.UpdateLoginAttempt(dbSess, user.Id, 0)
		if err != nil { // there's an error?
			session.AddFlash(err.Error(), FLASH_ERROR)
			return controller.Login_Page(c, r)
		}
	}

	// save this access token in session
	session.Values["AccessToken"] = accessToken

	return "/", http.StatusSeeOther
}

// logout
func (controller *UserController) Logout_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c)

	session.Values["AccessToken"] = nil

	return "/", http.StatusSeeOther
}

func (controller *UserController) Profile_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	username := c.URLParams["username"]
	if username == "" {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	username = strings.TrimSpace(username)

	// get user
	dbSess := controller.GetDb(c).NewSession(nil) // get db
	user, err := UserModel.GetUserByUsernameFull(dbSess, username)
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	} else if user.Id == 0 { // not found?
		pageController := PageController{}
		return pageController.Not_Found_Page(c, r)
	}

	c.Env["user"] = user
	c.Env["userProfilePicLink"] = helper.GetUserProfilePicLink(user.ProfilePic)

	// check whether can follow/unfollow user or not
	canFollow := false
	isFollowing := false

	accessToken := session.Values["AccessToken"]
	if accessToken != nil {
		c.Env["accessToken"] = accessToken.(string)

		// get current user
		currentUser, err := UserModel.GetUserByAccessTokenMin(dbSess, accessToken.(string))
		if err != nil && err != dbr.ErrNotFound { // there's an error?
			session.AddFlash(err.Error(), FLASH_ERROR)
			return "/", http.StatusSeeOther
		} else if user.Id != currentUser.Id { // if not own page, check whether following this user
			canFollow = true

			isFollowing, err = UserFollowModel.GetIsFollowingUser(dbSess, currentUser.Id, user.Id)
			if err != nil && err != dbr.ErrNotFound { // there's an error?
				session.AddFlash(err.Error(), FLASH_ERROR)
				return "/", http.StatusSeeOther
			}
		}
	} else {
		c.Env["accessToken"] = nil
	}

	c.Env["canFollow"] = canFollow
	c.Env["isFollowing"] = isFollowing

	// get country list
	countryList, err := CountryModel.GetAll(dbSess)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["userCountry"] = countryList[user.CountryId].Name

	// get number of followers
	followerCount, err := UserFollowModel.GetFollowerCount(dbSess, user.Id)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["followerCount"] = followerCount

	// get number of followings
	followingCount, err := UserFollowModel.GetFollowingCount(dbSess, user.Id)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["followingCount"] = followingCount

	c.Env["title"] = user.Username + "'s Page | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_profile", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// reset password page route
func (controller *UserController) Reset_Password_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	// no need to reset password if user already login
	if session.Values["AccessToken"] != nil {
		return "/", http.StatusSeeOther
	}

	// get token
	token := c.URLParams["token"]
	if token == "" {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	// generate hash token
	tokenHash := helper.GenerateHash(token)

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// check if token is still valid or not
	ticket, err := ResetPasswordTicketModel.GetTicketByToken(dbSess, tokenHash)
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	} else if ticket.UserId == 0 { // not found?
		session.AddFlash("Token is not valid or expired. Please request for a new reset password link.", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	c.Env["token"] = token

	c.Env["title"] = "Reset Password | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_reset_password", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// reset password action
func (controller *UserController) Reset_Password_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Reset_Password_Page(c, r)
	}

	// get post value
	token := r.FormValue("token")
	password := r.FormValue("password")

	// trim all white spaces
	token = strings.TrimSpace(token)
	password = strings.TrimSpace(password)

	// invalid input?
	if token == "" || len(password) < PASSWORD_LENGTH {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	tokenHash := helper.GenerateHash(token)

	// check if token is still valid or not
	ticket, err := ResetPasswordTicketModel.GetTicketByToken(dbSess, tokenHash)
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	} else if ticket.UserId == 0 { // not found?
		session.AddFlash("Token is not valid or expired. Please request for a new reset password link.", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	// save just in case need to return to previous page
	c.Env["token"] = token

	// hash password to save into db
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	// create transaction
	tx, err := dbSess.Begin()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	defer tx.RollbackUnlessCommitted()

	// update the password now
	err = UserModel.EditPasswordByUserId(tx, ticket.UserId, string(hashPassword))
	if err != nil { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	// set the password reset token to be invalid
	err = ResetPasswordTicketModel.InvalidateToken(tx, string(tokenHash))
	if err != nil { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	err = tx.Commit()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	// send this password to the user
	session.AddFlash("Password has been updated. You may login using the new password now", FLASH_SUCCESS)

	return "/login", http.StatusSeeOther
}

// settings password page
func (controller *UserController) Settings_Password_Page(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c)

	accessToken := session.Values["AccessToken"]
	if accessToken == nil {
		return "/", http.StatusSeeOther
	}

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	user, err := UserModel.GetUserByAccessTokenFull(dbSess, accessToken.(string))
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	} else if user.Id == 0 { // user not found?
		session.AddFlash("User not found", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	c.Env["user"] = user
	c.Env["userProfilePicLink"] = helper.GetUserProfilePicLink(user.ProfilePic)

	t := controller.GetTemplate(c)
	c.Env["title"] = "Change Password | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_settings_password", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// settings password action
func (controller *UserController) Settings_Password_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	}

	// get post value
	password := r.FormValue("password")
	newPassword := r.FormValue("new_password")
	accessToken := r.FormValue("token")

	// trim the strings
	password = strings.TrimSpace(password)
	newPassword = strings.TrimSpace(newPassword)
	accessToken = strings.TrimSpace(accessToken)

	// invalid input?
	if password == "" || len(newPassword) < PASSWORD_LENGTH || accessToken == "" {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	}

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// get user
	user, err := UserModel.GetPasswordByAccessToken(dbSess, accessToken)

	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	} else if user.Id == 0 { // user not found?
		session.AddFlash("User not found", FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	}

	// check whether password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// wrong password!
	if err != nil {
		session.AddFlash("Wrong password", FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	}

	// is correct, save new password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	// edit profile
	err = UserModel.EditPasswordByToken(dbSess, accessToken, string(hashPassword))
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Settings_Password_Page(c, r)
	}

	// no error
	session.AddFlash("Password updated", FLASH_SUCCESS)

	return controller.Settings_Password_Page(c, r)
}

// settings profile page
func (controller *UserController) Settings_Profile_Page(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c)

	accessToken := session.Values["AccessToken"]
	if accessToken == nil {
		return "/", http.StatusSeeOther
	}

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// get country list
	countryList, err := CountryModel.GetAll(dbSess)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["countryList"] = countryList

	// get gender list
	genderList, err := GenderModel.GetAll(dbSess)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["genderList"] = genderList

	// get user
	user, err := UserModel.GetUserByAccessTokenFull(dbSess, accessToken.(string))
	if err != nil && err != dbr.ErrNotFound { // there's an error?
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	} else if user.Id == 0 { // user not found?
		session.AddFlash("User not found", FLASH_ERROR)
		return "/", http.StatusSeeOther
	}

	t := controller.GetTemplate(c)
	c.Env["title"] = "Edit your profile | Exchange"

	c.Env["user"] = user
	c.Env["userProfilePicLink"] = helper.GetUserProfilePicLink(user.ProfilePic)

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_settings_profile", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// settings profile action
func (controller *UserController) Settings_Profile_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Settings_Profile_Page(c, r)
	}

	// get post value
	email := r.FormValue("email")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	gender := r.FormValue("gender")
	birthday := r.FormValue("birthday")
	country := r.FormValue("country")
	bio := r.FormValue("bio")
	accessToken := r.FormValue("token")

	// trim the strings
	email = strings.TrimSpace(email)
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	birthdayTime, err := time.Parse("02-01-2006", birthday)
	if err == nil {
		birthday = helper.IntToString(birthdayTime.Year()) + "-" + helper.IntToString(int(birthdayTime.Month())) + "-" + helper.IntToString(birthdayTime.Day())
	} else {
		birthday = "0000-00-00"
	}
	bio = strings.TrimSpace(bio)
	accessToken = strings.TrimSpace(accessToken)

	// validate the inputs
	isValidEmail := govalidator.IsEmail(email)
	isValidCountry := govalidator.Contains(country, "number:")

	// invalid input?
	if isValidEmail == false || isValidCountry == false || accessToken == "" {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return controller.Settings_Profile_Page(c, r)
	}

	if govalidator.Contains(gender, "string:") {
		gender = gender[7:len(gender)]
	} else {
		gender = ""
	}
	country = country[7:len(country)]

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// edit profile
	err = UserModel.Edit(dbSess, accessToken, email, firstName, lastName, gender, birthday, country, bio)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Settings_Profile_Page(c, r)
	}

	// no error
	session.AddFlash("Profile updated", FLASH_SUCCESS)

	return controller.Settings_Profile_Page(c, r)
}

// signup page route
func (controller *UserController) Signup_Page(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)
	session := controller.GetSession(c) // get session
	dbSess := controller.GetDb(c).NewSession(nil)

	// get country list
	countryList, err := CountryModel.GetAll(dbSess)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return "/", http.StatusSeeOther
	}
	c.Env["countryList"] = countryList

	c.Env["title"] = "Signup | Exchange"

	helper.RetrieveFlashes(session, c)

	widgets := helper.Parse(t, "user_signup", c.Env)
	c.Env["content"] = template.HTML(widgets)

	return helper.Parse(t, "main", c.Env), http.StatusOK
}

// signup action
func (controller *UserController) Signup_Action(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c) // get session

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	// get post value
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	country := r.FormValue("country")

	// trim input
	email = strings.TrimSpace(email)
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	country = strings.TrimSpace(country)

	// validate the inputs
	isValidEmail := govalidator.IsEmail(email)
	isValidCountry := govalidator.Contains(country, "number:")

	// invalid input?
	if isValidEmail == false || username == "" || len(password) < PASSWORD_LENGTH || isValidCountry == false {
		session.AddFlash("Invalid request", FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	// remove "number:" from country
	country = country[7:len(country)]

	// save the data just in case fail, can auto populate form
	c.Env["email"] = email
	c.Env["username"] = username
	c.Env["country"] = country

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// check whether email or username already exists
	userList, err := UserModel.GetUserByEmailOrUsername(dbSess, email, username)

	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	for _, user := range userList {
		if user.Email == email {
			session.AddFlash("Email already in use", FLASH_ERROR)
		}
		if user.Username == username {
			session.AddFlash("Username already in use", FLASH_ERROR)
		}
	}

	// we found an existing user, return error
	if len(userList) > 0 {
		return controller.Signup_Page(c, r)
	}

	// create hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	// get ip address
	ipAddress := helper.GetIp(r)

	// register this user
	err = UserModel.Signup(dbSess, email, username, string(hashPassword), country, ipAddress)
	if err != nil {
		session.AddFlash(err.Error(), FLASH_ERROR)
		return controller.Signup_Page(c, r)
	}

	// no error
	session.AddFlash("Register success", FLASH_SUCCESS)
	return "/", http.StatusSeeOther
}

// upload file
func (controller *UserController) UploadProfilePic_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	profilePic, header, err := r.FormFile("profile_pic")
	token := r.FormValue("token")

	// check whether the file is ok or not
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	} else if token == "" || filepath.Ext(header.Filename) == "" {
		helper.ReturnErr(w, "Invalid request", ERR_SYSTEM)
		return
	}

	defer profilePic.Close()

	dbSess := controller.GetDb(c).NewSession(nil) // get db

	// check whether user exists
	user, err := UserModel.GetProfilePicByAccessToken(dbSess, token)

	if err != nil && err != dbr.ErrNotFound { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	} else if user.Id == 0 { // user not found?
		helper.ReturnErr(w, "User not found or token expired", ERR_AUTH)
		return
	}

	// check whether the file is jpg or png format
	isFileImage, err := helper.IsImageFile(profilePic)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	} else if isFileImage == false {
		helper.ReturnErr(w, "Not image", ERR_SYSTEM)
		return
	}

	// check file size
	fileSize, err := helper.GetFileSize(profilePic)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	} else if fileSize > PROFILE_PIC_SIZE {
		helper.ReturnErr(w, "Size exceeded "+PROFILE_PIC_SIZE_STRING, ERR_SYSTEM)
		return
	}

	// create transaction
	tx, err := dbSess.Begin()
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	defer tx.RollbackUnlessCommitted()

	// if user already has profile pic, use its name, else generate a random name
	var profilePicName string

	if user.ProfilePic == "" {
		profilePicName = user.Username + "_" + helper.RandStr(PROFILE_PIC_RAND_LENGTH) + filepath.Ext(header.Filename)

		// save in db
		err = UserModel.UpdateProfilePicByToken(tx, token, profilePicName)
		if err != nil {
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}
	} else {
		profilePicName = user.ProfilePic
	}

	// create the temporary file
	dir, err := filepath.Abs(PROFILE_PIC_DIR)
	fmt.Println(dir)

	out, err := os.Create(dir + profilePicName)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, profilePic)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	data := struct {
		UserProfilePicLink string `json:"userProfilePicLink"`
	}{
		UserProfilePicLink: helper.GetUserProfilePicLink(profilePicName),
	}

	json.NewEncoder(w).Encode(&data)
}
