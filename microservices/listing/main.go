package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/urantiatech/teachingmission/microservices/listing/db"
)

func main() {
	var listen = flag.String("listen", ":9012", "HTTP listen address")
	var database = flag.String("database", "teachingmission", "The database name")
	flag.Parse()

	db.Name = *database

	var svc ListingService
	svc = listingService{}

	categoriesHandler := httptransport.NewServer(
		makeCategoriesEndpoint(svc),
		decodeCategoriesRequest,
		encodeResponse,
	)

	sectionsHandler := httptransport.NewServer(
		makeSectionsEndpoint(svc),
		decodeSectionsRequest,
		encodeResponse,
	)

	groupsHandler := httptransport.NewServer(
		makeGroupsEndpoint(svc),
		decodeGroupsRequest,
		encodeResponse,
	)

	teachersHandler := httptransport.NewServer(
		makeTeachersEndpoint(svc),
		decodeTeachersRequest,
		encodeResponse,
	)

	beingsHandler := httptransport.NewServer(
		makeBeingsEndpoint(svc),
		decodeBeingsRequest,
		encodeResponse,
	)

	receiversHandler := httptransport.NewServer(
		makeReceiversEndpoint(svc),
		decodeReceiversRequest,
		encodeResponse,
	)

	listingsHandler := httptransport.NewServer(
		makeListingsEndpoint(svc),
		decodeListingsRequest,
		encodeResponse,
	)

	http.Handle("/categories", categoriesHandler)
	http.Handle("/sections", sectionsHandler)
	http.Handle("/groups", groupsHandler)
	http.Handle("/teachers", teachersHandler)
	http.Handle("/beings", beingsHandler)
	http.Handle("/receivers", receiversHandler)
	http.Handle("/listings", listingsHandler)

	http.ListenAndServe(*listen, nil)
}
