package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/microservices/banner/api"
)

func makeBannerEndpoint(svc BannerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.BannerRequest)
		return svc.Banner(ctx, req)
	}
}

func decodeBannerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.BannerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
