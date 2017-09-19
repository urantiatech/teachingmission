package english

import (
	"net/http"

	"github.com/urantiatech/teachingmission/apps/webapp/controllers"
)

const langcode = "en"

func Search(rw http.ResponseWriter, r *http.Request) {
	controllers.Search(langcode, rw, r)
}

func About(rw http.ResponseWriter, r *http.Request) {
	controllers.About(langcode, rw, r)
}

func Contact(rw http.ResponseWriter, r *http.Request) {
	controllers.Contact(langcode, rw, r)
}

func Article(rw http.ResponseWriter, r *http.Request) {
	controllers.Article(langcode, rw, r)
}

func Transcript(rw http.ResponseWriter, r *http.Request) {
	controllers.Transcript(langcode, rw, r)
}

func Page(rw http.ResponseWriter, r *http.Request) {
	controllers.Page(langcode, rw, r)
}

func Archives(rw http.ResponseWriter, r *http.Request) {
	controllers.Archives(langcode, rw, r)
}

func UserLogin(rw http.ResponseWriter, r *http.Request) {
	controllers.UserLogin(langcode, rw, r)
}

func UserRegister(rw http.ResponseWriter, r *http.Request) {
	controllers.UserRegister(langcode, rw, r)
}

func UserForgotPassword(rw http.ResponseWriter, r *http.Request) {
	controllers.UserForgotPassword(langcode, rw, r)
}

func UserLockScreen(rw http.ResponseWriter, r *http.Request) {
	controllers.UserLockScreen(langcode, rw, r)
}

func UserProfile(rw http.ResponseWriter, r *http.Request) {
	controllers.UserProfile(langcode, rw, r)
}

func NotFound(rw http.ResponseWriter, r *http.Request) {
	controllers.NotFound(langcode, rw, r)
}