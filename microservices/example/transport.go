package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeExampleEndpoint(svc ExampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(exampleRequest)
		v, err := svc.Example(ctx, req.S)
		if err != nil {
			return exampleResponse{v, err.Error()}, nil
		}
		return exampleResponse{v, ""}, nil
	}
}

func decodeExampleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request exampleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type exampleRequest struct {
	S string `json:"s"`
}

type exampleResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
