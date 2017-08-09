package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urantiatech/teachingmission/apps/webapp/routes"
)

func main() {
	// Parse command line parameters
	var (
		domain string
		dir    string
		port   int
	)
	flag.StringVar(&domain, "domain", "teachingmission.org", "the domain name to serve")
	flag.StringVar(&dir, "dir", "static", "the directory to serve files from")
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()

	// Define language specific routes
	lang := map[string]func(*mux.Router){
		"www": routes.Default,
		"en":  routes.English,
		"fr":  routes.French,
		"hi":  routes.Hindi,
		"es":  routes.Spanish,
		"ru":  routes.Russian,
		"cn":  routes.Chinese,
	}

	// Create a router
	r := mux.NewRouter()

	for subdomain, addRoutes := range lang {
		s := r.Host(subdomain + "." + domain).Subrouter()
		// Static routes
		s.PathPrefix("/images/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/font/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/css/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/js/").Handler(http.FileServer(http.Dir(dir)))
		addRoutes(s)
	}

	p := fmt.Sprint(port)
	log.Fatal(http.ListenAndServe(":"+p, r))
}
