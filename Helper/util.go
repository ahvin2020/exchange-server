package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const FLASH_SUCCESS = "success_flash"
const FLASH_ERROR = "error_flash"

// support email
const SUPPORT_EMAIL = "no-reply@exchange.com"
const SUPPORT_EMAIL_NAME = "Support Email"

// profile pic
const PROFILE_PIC_DIR = "./Upload/profiles/"
const PROFILE_PIC_SIZE = 5242880
const PROFILE_PIC_SIZE_STRING = "5MB"

const HMAC_SALT = "Z3EU6cN54RMvbAxmc4vaHDBRzd5CAwuT"
const LOGIN_TOKEN_HOUR_VALIDITY = 24 // how long can a login token last?
const PASSWORD_LENGTH = 8
const PASSWORD_TOKEN_LENGTH = 32        // how long is a reset password token?
const PASSWORD_TOKEN_HOUR_VALIDITY = 24 // how long can a reset password token last?

const INVALID_PASSWORD_BLOCK_MINUTES = 5 // block user for 5 minutes if he tries to many wrong passwords
const INVALID_PASSWORD_BLOCK_MSG = "You entered too many invalid passwords, please try again after 5 minutes"
const INVALID_PASSWORD_TRIES = 5

const PROFILE_PIC_RAND_LENGTH = 16

func Int64ToString(l int64) string {
	return strconv.FormatInt(l, 10)
}

const RUNES string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func BaseUrl(r *http.Request) string {
	var proto string
	if r.TLS != nil {
		proto = `https`
	} else {
		proto = `http`
	}

	host := r.Host
	return strings.Join([]string{proto, `://`, host}, ``)
}

func IsImageFile(file multipart.File) (bool, error) {
	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err := file.Read(buff)
	if err != nil {
		return false, err
	}

	var isImage bool = false

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg", "image/jpg":
		isImage = true
	case "image/gif":
		isImage = true
	case "image/png":
		isImage = true
	default:
		isImage = false
	}

	return isImage, nil
}

func GetFileSize(file multipart.File) (int64, error) {
	fileSize, err := file.Seek(0, 2) //2 = from end
	if err != nil {
		return -1, err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return -1, err
	}

	return fileSize, nil
}

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err
}

func StringToBool(str string) (bool, error) {
	b, err := strconv.ParseBool(str)
	return b, err
}

func GenerateHash(str string) string {
	mac := hmac.New(md5.New, []byte(HMAC_SALT))
	mac.Write([]byte(str))
	hashString := hex.EncodeToString(mac.Sum(nil))

	return hashString
}

func GetIp(r *http.Request) string {
	var forwarded string = r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	} else {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		return ip
	}
}

func GetUserProfilePicLink(profile_pic string) string {
	if profile_pic == "" {
		return "/public/images/Facebook-Blank-Photo.jpg"
	} else {
		return "/upload/profiles/" + profile_pic
	}
}

func RandStr(strSize int) string {
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = RUNES[v%byte(len(RUNES))]
	}
	return string(bytes)
}

func RetrieveFlashes(session *sessions.Session, c web.C) {
	c.Env["FlashError"] = session.Flashes(FLASH_ERROR)
	c.Env["FlashSuccess"] = session.Flashes(FLASH_SUCCESS)
}
