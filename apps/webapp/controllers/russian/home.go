package russian

import (
	"net/http"

	"github.com/urantiatech/teachingmission/apps/webapp/multilingual"
)

const lang = "ru"

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(multilingual.Get(lang, "title")))
}
