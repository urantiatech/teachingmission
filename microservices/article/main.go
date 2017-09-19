package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/urantiatech/teachingmission/microservices/article/db"
)

func main() {
	var listen = flag.String("listen", ":9006", "HTTP listen address")
	var database = flag.String("database", "teachingmission", "The database name")
	flag.Parse()

	db.Name = *database

	var svc ArticleService
	svc = articleService{}

	articleHandler := httptransport.NewServer(
		makeArticleEndpoint(svc),
		decodeArticleRequest,
		encodeResponse,
	)

	http.Handle("/article", articleHandler)

	http.ListenAndServe(*listen, nil)
}
