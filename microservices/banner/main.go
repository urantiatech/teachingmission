package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":9002", "HTTP listen address")
	)
	flag.Parse()

	var svc BannerService
	svc = bannerService{}

	bannerHandler := httptransport.NewServer(
		makeBannerEndpoint(svc),
		decodeBannerRequest,
		encodeResponse,
	)

	http.Handle("/banner", bannerHandler)

	http.ListenAndServe(*listen, nil)
}
