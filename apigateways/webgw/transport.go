package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	articleapi "github.com/urantiatech/teachingmission/microservices/article/api"
	bannerapi "github.com/urantiatech/teachingmission/microservices/banner/api"
	dictionaryapi "github.com/urantiatech/teachingmission/microservices/dictionary/api"
	listingapi "github.com/urantiatech/teachingmission/microservices/listing/api"
	quoteapi "github.com/urantiatech/teachingmission/microservices/quote/api"
	transcriptapi "github.com/urantiatech/teachingmission/microservices/transcript/api"
)

func makeArticleEndpoint(svc WebApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ArticleRequest)
		v, err := svc.ArticlePage(ctx, req)
		return v, err
	}
}

func makeListingsEndpoint(svc WebApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ListingsRequest)
		v, err := svc.ListingsPage(ctx, req)
		return v, err
	}
}

func makeTranscriptEndpoint(svc WebApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TranscriptRequest)
		v, err := svc.TranscriptPage(ctx, req)
		return v, err
	}
}

func makeSearchEndpoint(svc WebApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SearchRequest)
		v, err := svc.SearchPage(ctx, req)
		return v, err
	}
}

func decodeTranscriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.TranscriptRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	// request.Language = r.Header.Get("Language-Code")
	return request, nil
}

func decodeArticleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	// request.Language = r.Header.Get("Language-Code")
	return request, nil
}

func decodeListingsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ListingsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	// request.Language = r.Header.Get("Language-Code")
	return request, nil
}

func decodePageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request PageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	// request.Language = r.Header.Get("Language-Code")
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeBannerResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf bannerapi.BannerResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeQuoteResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf quoteapi.QuoteResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeArticleResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf articleapi.ArticleResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeListingsResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf listingapi.ListingsResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeTooltipResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf dictionaryapi.TooltipResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeTaptargetResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf dictionaryapi.TaptargetResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeTranslationsResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf transcriptapi.TranslationsResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeTranscriptResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf transcriptapi.TranscriptResponse
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeSearchResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf transcriptapi.SearchResults
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}
