package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/microservices/dictionary/api"
)

func makeDefinitionEndpoint(svc DefinitionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.DefinitionRequest)
		return svc.Definition(ctx, req)
	}
}

func makeTooltipEndpoint(svc DefinitionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TooltipRequest)
		return svc.Tooltip(ctx, req)
	}
}

func makeTaptargetEndpoint(svc DefinitionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TaptargetRequest)
		return svc.Taptarget(ctx, req)
	}
}

func makeTextEndpoint(svc DefinitionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TextRequest)
		return svc.Text(ctx, req)
	}
}

func decodeDefinitionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.DefinitionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTooltipRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TooltipRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTaptargetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TaptargetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTextRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TextRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
