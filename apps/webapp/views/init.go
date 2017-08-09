package views

import (
	"html/template"
	"path/filepath"
)

var tmpl *template.Template

const (
	HomeLayout         = "home-layout"
	SearchLayout       = "search-layout"
	TranscriptLayout   = "transcript-layout"
	SubscriptionLayout = "subscription-layout"
	FAQLayout          = "faq-layout"
)

func init() {
	ParseTemplates()
}

func ParseTemplates() {
	pages := filepath.Join("templates", "*.tmpl")
	tmpl = template.Must(template.ParseGlob(pages))

	layouts := filepath.Join("templates/layouts", "*.tmpl")
	tmpl = template.Must(tmpl.ParseGlob(layouts))

	partials := filepath.Join("templates/partials", "*.tmpl")
	tmpl = template.Must(tmpl.ParseGlob(partials))

	search := filepath.Join("templates/search", "*.tmpl")
	tmpl = template.Must(tmpl.ParseGlob(search))
}
