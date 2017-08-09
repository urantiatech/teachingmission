package english

import (
	"context"
	"net/http"

	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	_ "github.com/urantiatech/teachingmission/apps/webapp/multilingual"
	"github.com/urantiatech/teachingmission/apps/webapp/views"
)

const lang = "en"

var ctx context.Context

func init() {
	ctx = context.WithValue(context.Background(), api.ContextKey("lang"), lang)
}

func Home(rw http.ResponseWriter, req *http.Request) {
	var p *api.Page
	keywords := req.FormValue("search")

	if keywords != "" {
		p = api.SearchPage(ctx, keywords)
		views.Render(rw, views.SearchLayout, p)
	} else {
		p = api.HomePage(ctx)
		views.Render(rw, views.HomeLayout, p)
	}
}

func Transcript(rw http.ResponseWriter, req *http.Request) {
	p := api.TranscriptPage(ctx)
	views.Render(rw, views.TranscriptLayout, p)
}
