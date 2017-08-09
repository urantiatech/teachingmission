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

	var svc ExampleService
	svc = exampleService{}

	exampleHandler := httptransport.NewServer(
		makeExampleEndpoint(svc),
		decodeExampleRequest,
		encodeResponse,
	)

	http.Handle("/example", exampleHandler)

	http.ListenAndServe(*listen, nil)
}
