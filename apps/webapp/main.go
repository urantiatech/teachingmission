package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urantiatech/teachingmission/apps/webapp/routes"
	"github.com/urantiatech/teachingmission/apps/webapp/views"
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
	langRouteMap := map[string]routes.LanguageRoutes{
		"www": {"English", "United-States.png", routes.Default},
		//"en":  {"English", "United-States.png", routes.English},
		// "fr": {"French", "France.png", routes.French},
		//"hi": {"Hindi", "India.png", routes.Hindi},
		// "es": {"Spanish", "Spain.png", routes.Spanish},
		// "ru": {"Russian", "Russia.png", routes.Russian},
		// "cn": {"Chinese", "China.png", routes.Chinese},
	}

	// Create a router
	r := mux.NewRouter()

	for subdomain, lang := range langRouteMap {
		s := r.Host(subdomain + "." + domain).Subrouter()

		// Static routes
		s.PathPrefix("/images/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/font/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/css/").Handler(http.FileServer(http.Dir(dir)))
		s.PathPrefix("/js/").Handler(http.FileServer(http.Dir(dir)))
		lang.Routes(s)

		// Add Frontend Flags
		langcode := subdomain
		if subdomain == "www" {
			langcode = "en"
		}
		l := views.Language{subdomain, lang.Language, lang.Flag}
		views.AddLanguage(langcode, l)
	}

	p := fmt.Sprint(port)
	log.Fatal(http.ListenAndServe(":"+p, r))
}
