package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/urantiatech/teachingmission/microservices/quote/db"
)

func main() {
	var listen = flag.String("listen", ":9004", "HTTP listen address")
	var database = flag.String("database", "teachingmission", "The database name")
	flag.Parse()

	db.Name = *database

	var svc QuoteService
	svc = quoteService{}

	quoteHandler := httptransport.NewServer(
		makeQuoteEndpoint(svc),
		decodeQuoteRequest,
		encodeResponse,
	)

	http.Handle("/quote", quoteHandler)

	http.ListenAndServe(*listen, nil)
}
