package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":9000", "HTTP listen address")
	)
	flag.Parse()

	var svc WebApiService
	svc = webapiService{}

	articleHandler := httptransport.NewServer(
		makeArticleEndpoint(svc),
		decodeArticleRequest,
		encodeResponse,
	)

	listingsHandler := httptransport.NewServer(
		makeListingsEndpoint(svc),
		decodeListingsRequest,
		encodeResponse,
	)

	transcriptHandler := httptransport.NewServer(
		makeTranscriptEndpoint(svc),
		decodeTranscriptRequest,
		encodeResponse,
	)

	searchHandler := httptransport.NewServer(
		makeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)

	http.Handle("/article", articleHandler)
	http.Handle("/transcript", transcriptHandler)
	http.Handle("/search", searchHandler)
	http.Handle("/listings", listingsHandler)

	http.ListenAndServe(*listen, nil)
}
