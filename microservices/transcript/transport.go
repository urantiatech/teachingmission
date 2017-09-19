package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	listingapi "github.com/urantiatech/teachingmission/microservices/listing/api"
	"github.com/urantiatech/teachingmission/microservices/transcript/api"
)

func makeTranscriptEndpoint(svc TranscriptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TranscriptRequest)
		return svc.Transcript(ctx, req)
	}
}

func makeTranslationsEndpoint(svc TranscriptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TranslationsRequest)
		return svc.Translations(ctx, req)
	}
}

func makeSearchEndpoint(svc TranscriptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SearchRequest)
		return svc.Search(ctx, req)
	}
}

func decodeTranscriptRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TranscriptRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTranslationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TranslationsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeTeachersResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf listingapi.TeachersResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
