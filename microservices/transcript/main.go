package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/urantiatech/teachingmission/microservices/transcript/db"
)

func main() {
	var listen = flag.String("listen", ":9008", "HTTP listen address")
	var database = flag.String("database", "teachingmission", "The database name")
	flag.Parse()

	db.Name = *database

	var svc TranscriptService
	svc = transcriptService{}

	transcriptHandler := httptransport.NewServer(
		makeTranscriptEndpoint(svc),
		decodeTranscriptRequest,
		encodeResponse,
	)

	translationsHandler := httptransport.NewServer(
		makeTranslationsEndpoint(svc),
		decodeTranslationsRequest,
		encodeResponse,
	)

	searchHandler := httptransport.NewServer(
		makeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)

	http.Handle("/transcript", transcriptHandler)
	http.Handle("/translations", translationsHandler)
	http.Handle("/search", searchHandler)

	http.ListenAndServe(*listen, nil)
}
