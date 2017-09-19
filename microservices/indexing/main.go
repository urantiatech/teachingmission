package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var listen = flag.String("listen", ":9014", "HTTP listen address")
	flag.Parse()

	var svc IndexingService
	svc = indexingService{}

	indexHandler := httptransport.NewServer(
		makeIndexEndpoint(svc),
		decodeIndexRequest,
		encodeResponse,
	)

	searchHandler := httptransport.NewServer(
		makeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)

	http.Handle("/index", indexHandler)
	http.Handle("/search", searchHandler)

	http.ListenAndServe(*listen, nil)
}
