package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/urantiatech/teachingmission/microservices/dictionary/db"
)

func main() {
	var listen = flag.String("listen", ":9010", "HTTP listen address")
	var database = flag.String("database", "teachingmission", "The database name")
	flag.Parse()

	db.Name = *database

	var svc DefinitionService
	svc = definitionService{}

	definitionHandler := httptransport.NewServer(
		makeDefinitionEndpoint(svc),
		decodeDefinitionRequest,
		encodeResponse,
	)

	tooltipHandler := httptransport.NewServer(
		makeTooltipEndpoint(svc),
		decodeTooltipRequest,
		encodeResponse,
	)

	taptargetHandler := httptransport.NewServer(
		makeTaptargetEndpoint(svc),
		decodeTaptargetRequest,
		encodeResponse,
	)

	textHandler := httptransport.NewServer(
		makeTextEndpoint(svc),
		decodeTextRequest,
		encodeResponse,
	)

	http.Handle("/definition", definitionHandler)
	http.Handle("/tooltip", tooltipHandler)
	http.Handle("/taptarget", taptargetHandler)
	http.Handle("/text", textHandler)

	http.ListenAndServe(*listen, nil)
}
