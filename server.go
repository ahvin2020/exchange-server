package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/context"
	"net/http"
	"runtime"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"exchange.com/exchange/controller"
	"exchange.com/exchange/system"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// load config file
	configFile := flag.String("config", "config.toml", "Path to configuration file")
	flag.Parse()
	defer glog.Flush()

	// init core
	var application = &system.Application{}
	application.Init(configFile)
	application.LoadTemplates()

	// setup static files
	static := web.New()

	publicPath := application.Config.Get("general.public_path").(string)
	static.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(publicPath))))
	http.Handle("/public/", static)

	uploadPath := application.Config.Get("general.upload_path").(string)
	static.Get("/upload/*", http.StripPrefix("/upload/", http.FileServer(http.Dir(uploadPath))))
	http.Handle("/upload/", static)

	bowerPath := application.Config.Get("general.bower_path").(string)
	static.Get("/bower_components/*", http.StripPrefix("/bower_components/", http.FileServer(http.Dir(bowerPath))))
	http.Handle("/bower_components/", static)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)
	goji.Use(application.ApplyDb)
	goji.Use(application.ApplyAuth)

	//	goji.Use(application.ApplyIsXhr)
	//	goji.Use(application.ApplyCsrfProtection)
	goji.Use(context.ClearHandler)

	// Couple of files - in the real world you would use nginx to serve them.
	goji.Get("/robots.txt", http.FileServer(http.Dir(publicPath)))
	goji.Get("/favicon.ico", http.FileServer(http.Dir(publicPath+"/images")))

	// --- routes ---
	// page
	pageController := &controller.PageController{}
	goji.Get("/", application.Route(pageController, "Index_Page"))

	// country
	countryController := &controller.CountryController{}
	goji.Get("/api/country_list", countryController.List_Api)

	// user
	userController := &controller.UserController{}
	goji.Get("/login", application.Route(userController, "Login_Page"))
	goji.Post("/login", application.Route(userController, "Login_Action"))
	goji.Get("/logout", application.Route(userController, "Logout_Action"))
	goji.Get("/signup", application.Route(userController, "Signup_Page"))
	goji.Post("/signup", application.Route(userController, "Signup_Action"))
	goji.Get("/forgot-password", application.Route(userController, "Forgot_Password_Page"))
	goji.Post("/forgot-password", application.Route(userController, "Forgot_Password_Action"))
	goji.Get("/:username", application.Route(userController, "Profile_Page"))
	goji.Get("/reset-password/:token", application.Route(userController, "Reset_Password_Page"))
	goji.Post("/reset-password", application.Route(userController, "Reset_Password_Action"))
	goji.Get("/settings/password", application.Route(userController, "Settings_Password_Page"))
	goji.Post("/settings/password", application.Route(userController, "Settings_Password_Action"))
	goji.Get("/settings/profile", application.Route(userController, "Settings_Profile_Page"))
	goji.Post("/settings/profile", application.Route(userController, "Settings_Profile_Action"))
	goji.Post("/api/upload-profile-pic", userController.UploadProfilePic_Api)

	// user follows
	userFollowController := &controller.UserFollowController{}
	goji.Post("/api/follow-user", userFollowController.FollowUser_Api)
	goji.Get("/api/get-followers", userFollowController.GetFollowers_Api)
	goji.Get("/api/get-followings", userFollowController.GetFollowings_Api)

	// user currency
	userCurrencyController := &controller.UserCurrencyController{}
	goji.Get("/usercurrency", application.Route(userCurrencyController, "Index_Page"))
	goji.Get("/api/exchange_rate", userCurrencyController.Exchange_Rate_Api)
	goji.Get("/api/usercurrency_list", userCurrencyController.List_Api)
	goji.Post("/api/usercurrency_update", userCurrencyController.Update_Api)
	goji.Post("/api/usercurrency_delete", userCurrencyController.Delete_Api)

	graceful.PostHook(func() {
		application.Close()
	})

	goji.Serve()
}
