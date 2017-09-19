package main

import (
	"context"
	"errors"

	"github.com/urantiatech/teachingmission/microservices/indexing/api"
)

type IndexingService interface {
	Index(context.Context, api.IndexRequest) (api.IndexResponse, error)
	Search(context.Context, api.SearchRequest) (api.SearchResponse, error)
}

type indexingService struct{}

func (indexingService) Index(ctx context.Context, req api.IndexRequest) (api.IndexResponse, error) {
	var resp api.IndexResponse

	return resp, nil
}

func (indexingService) Search(ctx context.Context, req api.SearchRequest) (api.SearchResponse, error) {
	var resp api.SearchResponse

	return resp, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for IndexingService.
type ServiceMiddleware func(IndexingService) IndexingService
