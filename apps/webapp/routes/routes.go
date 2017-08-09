package routes

import (
	"github.com/gorilla/mux"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/chinese"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/english"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/french"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/hindi"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/russian"
	"github.com/urantiatech/teachingmission/apps/webapp/controllers/spanish"
)

func Default(r *mux.Router) {
	r.HandleFunc("/", english.Home)
	//	r.HandleFunc("/search", english.Search)
	r.HandleFunc("/transcript", english.Transcript)
}

func English(r *mux.Router) {
	// Redirect all en.teachingmission.org routs to www.teachingmission.org
	r.HandleFunc("/", english.Home)
}

func French(r *mux.Router) {
	r.HandleFunc("/", french.Home)
}

func Chinese(r *mux.Router) {
	r.HandleFunc("/", chinese.Home)
}

func Russian(r *mux.Router) {
	r.HandleFunc("/", russian.Home)
}

func Spanish(r *mux.Router) {
	r.HandleFunc("/", spanish.Home)
}

func Hindi(r *mux.Router) {
	r.HandleFunc("/", hindi.Home)
}
