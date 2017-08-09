package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/pkg/database/transcript"
)

func makeTranscriptEndpoint(svc TranscriptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transcriptRequest)
		t, err := svc.Transcript(ctx, req.Id)
		if err != nil {
			return transcriptResponse{t, err.Error()}, nil
		}
		return transcriptResponse{t, ""}, nil
	}
}

func makeSearchEndpoint(svc TranscriptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)
		results, err := svc.Search(ctx, req)
		if err != nil {
			return searchResponse{results, err.Error()}, nil
		}
		return searchResponse{results, ""}, nil
	}
}

func decodeTranscriptRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transcriptRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request searchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type transcriptRequest struct {
	Id string `json:"id"`
}

type transcriptResponse struct {
	T   transcript.Transcript `json:"transcript"`
	Err string                `json:"err,omitempty"`
}

type searchRequest struct {
	Language   string `bson:"language"`
	Visibility string `bson:"visibility"`
	Status     string `bson:"status"`
	Query      string `bson:"query"`
	Teacher    string `bson:"teacher"`
	Receiver   string `bson:"receiver"`
	Group      string `bson:"group"`
	Category   string `bson:"category"`
	StartDate  string `bson:"startdate"`
	EndDate    string `bson:"enddate"`
}

type searchResponse struct {
	R   []transcript.SearchRecord `json:"results"`
	Err string                    `json:"err,omitempty"`
}
