package views

import (
	"fmt"
	"net/http"

	api "github.com/urantiatech/teachingmission/apigateways/webgw/api"
)

func Render(rw http.ResponseWriter, layout string, p *api.Page) {

	ParseTemplates() // Parsing should be removed from here for production

	err := tmpl.ExecuteTemplate(rw, layout, p)
	if err != nil {
		fmt.Println(err.Error())
	}
}
