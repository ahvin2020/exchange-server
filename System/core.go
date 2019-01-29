package system

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"crypto/sha256"

	"github.com/gocraft/dbr"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"github.com/pelletier/go-toml"
	"github.com/zenazn/goji/web"
)

//type CsrfProtection struct {
//	Key    string
//	Cookie string
//	Header string
//	Secure bool
//}

//var Db *dbr.Connection

type Application struct {
	Config   *toml.TomlTree
	Template *template.Template
	Store    *sessions.CookieStore
	Db *dbr.Connection
	//	CsrfProtection *CsrfProtection
}

func (application *Application) Init(filename *string) {

	config, err := toml.LoadFile(*filename)
	if err != nil {
		glog.Fatalf("TOML load failed: %s\n", err)
	}

	hash := sha256.New()
	io.WriteString(hash, config.Get("cookie.mac_secret").(string))
	application.Store = sessions.NewCookieStore(hash.Sum(nil))
	application.Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   config.Get("cookie.secure").(bool),
	}

	// connect to db
	dbConfig := config.Get("database").(*toml.TomlTree)
	CONN_STRING := dbConfig.Get("user").(string) + ":" + dbConfig.Get("password").(string) + "@tcp(" + dbConfig.Get("hostname").(string) + ":" + dbConfig.Get("port").(string) + ")/" + dbConfig.Get("database").(string) + "?charset=utf8&parseTime=true&loc=Asia%2FSingapore"
	db, err := sql.Open("mysql", CONN_STRING)
	if err != nil {
		fmt.Println("Error in db open : " + err.Error())
		panic(fmt.Sprintf("Error when connecting SQL: '%v'", err))
	}
	fmt.Println("MYSQL DBR CONNECTED")
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(500)
	application.Db = dbr.NewConnection(db, nil)

	// init csrf protection
	//	application.CsrfProtection = &CsrfProtection{
	//		Key:    config.Get("csrf.key").(string),
	//		Cookie: config.Get("csrf.cookie").(string),
	//		Header: config.Get("csrf.header").(string),
	//		Secure: config.Get("cookie.secure").(bool),
	//	}

	application.Config = config
}

func (application *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}
	err := filepath.Walk(application.Config.Get("general.template_path").(string), fn)

	if err != nil {
		return err
	}

	appTemplate := template.New("templ")
	appTemplate = appTemplate.Delims("[[", "]]")
	appTemplate, err = appTemplate.ParseFiles(templates...)

	if err != nil {
		log.Println(err.Error())
	}
	application.Template = appTemplate
	
	return nil
}

func (application *Application) Close() {
	glog.Info("Bye!")
}

func (application *Application) Route(controller interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		c.Env["Content-Type"] = "text/html"
		
		methodValue := reflect.ValueOf(controller).MethodByName(route)
		methodInterface := methodValue.Interface()
		method := methodInterface.(func(c web.C, r *http.Request) (string, int))

		body, code := method(c, r)

		if session, exists := c.Env["Session"]; exists {
			
			err := session.(*sessions.Session).Save(r, w)
			if err != nil {
				glog.Errorf("Can't save session: %v", err)
			}
		}

		switch code {
		case http.StatusOK:
			if _, exists := c.Env["Content-Type"]; exists {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			}
			io.WriteString(w, body)
		case http.StatusSeeOther, http.StatusFound:
			http.Redirect(w, r, body, code)
		}
	}
	return fn
}