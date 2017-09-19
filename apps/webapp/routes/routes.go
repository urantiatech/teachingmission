package routes

import (
	"github.com/gorilla/mux"
	"github.com/urantiatech/teachingmission/apps/webapp/routes/english"
	"github.com/urantiatech/teachingmission/apps/webapp/routes/french"
	"github.com/urantiatech/teachingmission/apps/webapp/routes/russian"
	"github.com/urantiatech/teachingmission/apps/webapp/routes/spanish"
)

type LanguageRoutes struct {
	Language string
	Flag     string
	Routes   func(*mux.Router)
}

func Default(r *mux.Router) {
	// r.HandleFunc("/", english.Home)
	r.HandleFunc("/", english.Search)
	r.HandleFunc("/about", english.About)
	r.HandleFunc("/contact", english.Contact)
	r.HandleFunc("/article/{id}", english.Article)
	r.HandleFunc("/page/{id}", english.Page)
	r.HandleFunc("/transcript/{id}", english.Transcript)

	r.HandleFunc("/archives", english.Archives)

	// User Routes
	//r.HandleFunc("/login", english.UserLogin)
	//r.HandleFunc("/register", english.UserRegister)
	//r.HandleFunc("/forgot-password", english.UserForgotPassword)
	//r.HandleFunc("/lock-screen", english.UserLockScreen)
	//r.HandleFunc("/profile", english.UserProfile)

}

func English(r *mux.Router) {
	// Redirect all en.teachingmission.org routs to www.teachingmission.org
}

func French(r *mux.Router) {
	// r.HandleFunc("/", french.Home)
	r.HandleFunc("/", french.Search)
	r.HandleFunc("/about", french.About)
	r.HandleFunc("/contact", french.Contact)
	r.HandleFunc("/article/{id}", french.Article)
	r.HandleFunc("/transcript/{id}", french.Transcript)

	// User Routes
	r.HandleFunc("/login", french.UserLogin)
	r.HandleFunc("/register", french.UserRegister)
	r.HandleFunc("/profile", french.UserProfile)
}

func Russian(r *mux.Router) {
	// r.HandleFunc("/", russian.Home)
	r.HandleFunc("/", russian.Search)
	r.HandleFunc("/about", russian.About)
	r.HandleFunc("/contact", russian.Contact)
	r.HandleFunc("/article/{id}", russian.Article)
	r.HandleFunc("/transcript/{id}", russian.Transcript)

	// User Routes
	r.HandleFunc("/login", russian.UserLogin)
	r.HandleFunc("/register", russian.UserRegister)
	r.HandleFunc("/profile", russian.UserProfile)
}

func Spanish(r *mux.Router) {
	// r.HandleFunc("/", spanish.Home)
	r.HandleFunc("/", spanish.Search)
	r.HandleFunc("/about", spanish.About)
	r.HandleFunc("/contact", spanish.Contact)
	r.HandleFunc("/article/{id}", spanish.Article)
	r.HandleFunc("/transcript/{id}", spanish.Transcript)

	// User Routes
	r.HandleFunc("/login", spanish.UserLogin)
	r.HandleFunc("/register", spanish.UserRegister)
	r.HandleFunc("/profile", spanish.UserProfile)
}
