package main

import (
	"context"
	"fmt"
	"net/url"

	httptransport "github.com/go-kit/kit/transport/http"
	listingapi "github.com/urantiatech/teachingmission/microservices/listing/api"
)

type TeachersMap map[string]string

func getTeachers(ctx context.Context, langcode string, being string) *TeachersMap {
	teachers := make(TeachersMap)

	// Create a listing svc request for teachers
	req := listingapi.TeachersRequest{Language: langcode, BeingId: being}

	tgt, err := url.Parse("http://localhost:9012/teachers")
	if err != nil {
		fmt.Println(err.Error())
	}
	endPoint := httptransport.NewClient("POST", tgt,
		encodeRequest, decodeTeachersResponse).Endpoint()

	resp, err := endPoint(ctx, req)
	if err != nil {
		// Return empty teacher map
		return &teachers
	}

	// Copy teachers
	for _, t := range resp.(listingapi.TeachersResponse).Teachers {
		teachers[t.Id] = t.Name
	}

	return &teachers
}
