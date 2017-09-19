package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/microservices/quote/api"
)

func makeQuoteEndpoint(svc QuoteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.QuoteRequest)
		return svc.Quote(ctx, req)
	}
}

func decodeQuoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.QuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
