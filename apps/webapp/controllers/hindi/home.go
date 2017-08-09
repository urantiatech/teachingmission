package hindi

import (
	"net/http"

	"github.com/urantiatech/teachingmission/apps/webapp/multilingual"
)

const lang = "hi"

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(multilingual.Get(lang, "title")))
}
