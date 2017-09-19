package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/microservices/listing/api"
)

func makeCategoriesEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.CategoriesRequest)
		return svc.Categories(ctx, req)
	}
}

func makeSectionsEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SectionsRequest)
		return svc.Sections(ctx, req)
	}
}

func makeGroupsEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.GroupsRequest)
		return svc.Groups(ctx, req)
	}
}

func makeTeachersEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TeachersRequest)
		return svc.Teachers(ctx, req)
	}
}

func makeBeingsEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.BeingsRequest)
		return svc.Beings(ctx, req)
	}
}

func makeReceiversEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ReceiversRequest)
		return svc.Receivers(ctx, req)
	}
}

func makeListingsEndpoint(svc ListingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ListingsRequest)
		return svc.Listings(ctx, req)
	}
}

func decodeCategoriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.CategoriesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSectionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SectionsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.GroupsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTeachersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TeachersRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeBeingsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.BeingsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeReceiversRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ReceiversRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeListingsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ListingsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
