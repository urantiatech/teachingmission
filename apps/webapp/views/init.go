package views

import (
	// "fmt"
	"path/filepath"
	"text/template"
)

// Page is 2-dimensional map. Accessed as Page['page']['layout']
var TemplateSet map[string]map[string]*template.Template

func init() {
	// Initialize the Map.
	TemplateSet = make(map[string]map[string]*template.Template)
	ParseTemplates()
}

func ParseTemplates() {
	var base *template.Template

	// Get all base files without duplicate block definitions
	masterGlob := filepath.Join("templates/master-layouts", "*.tmpl")
	partialGlob := filepath.Join("templates/partials", "*.tmpl")
	widgetGlob := filepath.Join("templates/widgets", "*.tmpl")

	// Define Base template
	base = template.New("").Funcs(funcMap)
	base = template.Must(base.ParseGlob(masterGlob))
	base = template.Must(base.ParseGlob(partialGlob))
	base = template.Must(base.ParseGlob(widgetGlob))

	// Parse Page Layouts
	layoutFiles, err := filepath.Glob("templates/page-layouts/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse Pages
	pageFiles, err := filepath.Glob("templates/pages/*.tmpl")
	if err != nil {
		panic(err)
	}

	// For each page template
	for _, page := range pageFiles {
		// Initialize the map using basename
		p := filepath.Base(page)
		TemplateSet[p] = make(map[string]*template.Template)

		// Clone the Base and add Layouts to each Page
		for _, layout := range layoutFiles {
			// Get the base filename
			l := filepath.Base(layout)
			// fmt.Printf("TemplateSet[%s][%s]\n", p, l)

			TemplateSet[p][l] = template.Must(base.Clone())
			TemplateSet[p][l] = template.Must(TemplateSet[p][l].ParseFiles(layout))
			TemplateSet[p][l] = template.Must(TemplateSet[p][l].ParseFiles(page))
		}
	}

}
