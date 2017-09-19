package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	"github.com/urantiatech/teachingmission/apps/webapp/constants"
	_ "github.com/urantiatech/teachingmission/apps/webapp/multilingual"
	"github.com/urantiatech/teachingmission/apps/webapp/views"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// Search page is a new Homepage
func Search(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a search request
	var req api.SearchRequest
	req.Language = langcode

	req.Query = r.FormValue("q") // Query string

	if len(req.Query) < constants.MinimumQuerySize {
		// Show home search form.
		list := api.ListingsRequest{
			Language:   langcode,
			Categories: true,
			Sections:   true,
			Groups:     true,
			Beings:     true,
			Teachers:   true,
			Receivers:  true}
		views.Render(ctx, rw, r, langcode, "layout-12.tmpl", "page-home-search.tmpl", "/listings", list)
		return
	}

	req.Category = r.FormValue("c")   // Category
	req.Being = r.FormValue("b")      // Being Type
	req.Teacher = r.FormValue("t")    // Teacher
	req.Receiver = r.FormValue("r")   // Receiver
	req.StartDate = r.FormValue("y1") // Start Date
	req.EndDate = r.FormValue("y2")   // End Date

	p, err := strconv.Atoi(r.FormValue("p"))
	if err != nil || p <= 0 {
		p = 1 // First page is p=1
	}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil || limit <= 0 || limit > constants.SearchResultsLimit {
		limit = constants.SearchResultsLimit
	}

	req.Skip = (p - 1) * limit
	req.Limit = limit
	views.Render(ctx, rw, r, langcode, "layout-12.tmpl", "page-search-results.tmpl", "/search", req)
}

func About(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a article request
	var req api.ArticleRequest
	req.Language = langcode
	req.Id = constants.AboutPage
	views.Render(ctx, rw, r, langcode, "layout-8-4.tmpl", "page-article.tmpl", "/article", req)
}

func Contact(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a article request
	var req api.ArticleRequest
	req.Language = langcode
	req.Id = constants.ContactPage
	views.Render(ctx, rw, r, langcode, "layout-8-4.tmpl", "page-contact.tmpl", "/article", req)
}

func Article(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a article request
	var req api.ArticleRequest
	req.Language = langcode
	vars := mux.Vars(r)
	req.Id = vars["id"]
	views.Render(ctx, rw, r, langcode, "layout-8-4.tmpl", "page-article.tmpl", "/article", req)
}

func Page(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a article request
	var req api.ArticleRequest
	req.Language = langcode
	vars := mux.Vars(r)
	req.Id = vars["id"]
	views.Render(ctx, rw, r, langcode, "layout-12.tmpl", "page-wide.tmpl", "/article", req)
}

func Transcript(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a transcript request
	var req api.TranscriptRequest
	req.Language = langcode
	vars := mux.Vars(r)
	req.Id = vars["id"]
	views.Render(ctx, rw, r, langcode, "layout-9-3.tmpl", "page-transcript.tmpl", "/transcript", req)
}

func Archives(langcode string, rw http.ResponseWriter, r *http.Request) {
	// Create a listing request
	var req api.ListingsRequest
	req.Language = langcode

	req.Categories = true
	req.Groups = true
	req.Beings = true
	req.Teachers = true
	req.Receivers = true
	req.Sections = true

	views.Render(ctx, rw, r, langcode, "layout-12.tmpl", "page-home-browse.tmpl", "/listings", req)
}

func UserLogin(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.RenderUserForm(ctx, rw, r, langcode, "layout-blank.tmpl", "user-login.tmpl")
}

func UserRegister(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.RenderUserForm(ctx, rw, r, langcode, "layout-blank.tmpl", "user-register.tmpl")
}

func UserForgotPassword(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.RenderUserForm(ctx, rw, r, langcode, "layout-blank.tmpl", "user-forgot-password.tmpl")
}

func UserLockScreen(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.RenderUserForm(ctx, rw, r, langcode, "layout-blank.tmpl", "user-lock-screen.tmpl")
}

func UserProfile(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.RenderUserForm(ctx, rw, r, langcode, "layout-blank.tmpl", "profile.tmpl")
}

func NotFound(langcode string, rw http.ResponseWriter, r *http.Request) {
	views.NotFound(rw, langcode)
}
