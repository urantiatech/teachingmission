package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var svc TranscriptService
	svc = transcriptService{}

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

	http.Handle("/transcript", transcriptHandler)
	http.Handle("/search", searchHandler)

	http.ListenAndServe(*listen, nil)
}
