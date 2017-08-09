package french

import (
	"context"
	"net/http"

	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	_ "github.com/urantiatech/teachingmission/apps/webapp/multilingual"
	"github.com/urantiatech/teachingmission/apps/webapp/views"
)

const lang = "fr"

var ctx context.Context

func init() {
	ctx = context.WithValue(context.Background(), api.ContextKey("lang"), lang)
}

func Home(rw http.ResponseWriter, req *http.Request) {
	p := api.HomePage(ctx)
	views.Render(rw, views.HomeLayout, p)
}
