package views

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	api "github.com/urantiatech/teachingmission/apigateways/webgw/api"
)

type Language struct {
	Subdomain string
	Name      string
	Flag      string
}

type LanguageMap map[string]Language

type Page struct {
	LangCode    string
	LanguageMap LanguageMap
	URI         string
	api.Page
}

var lmap LanguageMap

func init() {
	lmap = make(LanguageMap)
}

func AddLanguage(langcode string, l Language) {
	lmap[langcode] = Language{l.Subdomain, l.Name, l.Flag}
}

func Render(ctx context.Context, rw http.ResponseWriter, r *http.Request, langcode string,
	layout string, page string, uri string, req interface{}) {
	//var rw http.ResponseWriter
	ParseTemplates() // Parsing should be removed from here for production

	// fmt.Printf("Page=[%s], Layout=[%s], gwAPI=[%s]\n", page, layout, uri)

	// Handle null request
	if req == nil {
		// fmt.Println("received null request")
	}

	tgt, err := url.Parse("http://localhost:9000" + uri)
	if err != nil {
		InternalServerError(rw, langcode, err)
		return
	}

	endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodePageResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		InternalServerError(rw, langcode, err)
		return
	}

	p := resp.(api.Page)
	// Not sure whether list of Languages should come from API Gateway or Enabled from Frontend
	// As of now they are enabled by the Frontend in main function
	pg := Page{langcode, lmap, r.URL.Path, p}

	// Calculate breadcrumb based on URI
	breadcrumb := api.Breadcrumb{}
	tokens := strings.Split(r.URL.Path, "/")

	for i, t := range tokens {
		var b string
		switch i {
		case 0:
			b = "/"
		case len(tokens) - 1:
			b = pg.Title
		default:
			b = t
		}
		breadcrumb.Titles = append(breadcrumb.Titles, b)
	}
	pg.Breadcrumb = &breadcrumb

	var master string
	if page == "page-home-search.tmpl" || page == "page-home-browse.tmpl" {
		master = "master-home.tmpl"
	} else {
		master = "master.tmpl"
	}

	err = TemplateSet[page][layout].ExecuteTemplate(rw, master, pg)
	if err != nil {
		InternalServerError(rw, langcode, err)
		return
	}
}

func RenderUserForm(_ context.Context, rw http.ResponseWriter, r *http.Request, langcode, layout, page string) {
	ParseTemplates() // Parsing should be removed from here for production
	_ = TemplateSet[page][layout].ExecuteTemplate(rw, "blank.tmpl", nil)
}

func InternalServerError(rw http.ResponseWriter, langcode string, err error) {
	_ = TemplateSet["500.tmpl"]["layout-blank.tmpl"].ExecuteTemplate(rw, "blank.tmpl", nil)
	fmt.Println(err.Error())
}

func NotFound(rw http.ResponseWriter, langcode string) {
	_ = TemplateSet["404.tmpl"]["layout-blank.tmpl"].ExecuteTemplate(rw, "blank.tmpl", nil)
}
