package main

import (
	"context"
	"errors"
	"fmt"
	_ "net/http"

	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	articleapi "github.com/urantiatech/teachingmission/microservices/article/api"
	bannerapi "github.com/urantiatech/teachingmission/microservices/banner/api"
	listingapi "github.com/urantiatech/teachingmission/microservices/listing/api"
	quoteapi "github.com/urantiatech/teachingmission/microservices/quote/api"
	transcriptapi "github.com/urantiatech/teachingmission/microservices/transcript/api"
)

// WebApi provides operations on strings.
type WebApiService interface {
	ArticlePage(context.Context, api.ArticleRequest) (*api.Page, error)
	ListingsPage(context.Context, api.ListingsRequest) (*api.Page, error)
	TranscriptPage(context.Context, api.TranscriptRequest) (*api.Page, error)
	SearchPage(context.Context, api.SearchRequest) (*api.Page, error)
}

type webapiService struct{}

func (webapiService) ArticlePage(ctx context.Context, request api.ArticleRequest) (*api.Page, error) {
	var req PageRequest
	req.Language = request.Language
	req.Banner = new(bannerapi.BannerRequest)
	req.Quote = new(quoteapi.QuoteRequest)
	req.Article = new(articleapi.ArticleRequest)
	req.Article.Id = request.Id

	req.Listings = new(listingapi.ListingsRequest)
	req.Listings.Language = request.Language
	req.Listings.Categories = true
	req.Listings.Teachers = true
	req.Listings.Beings = true
	req.Listings.Receivers = true
	req.Listings.Sections = true

	res, err := CallOtherServices(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	page, err := AssemblePage(ctx, &req, &res)
	return page, err
}

func (webapiService) ListingsPage(ctx context.Context, request api.ListingsRequest) (*api.Page, error) {
	var req PageRequest
	req.Language = request.Language

	req.Listings = new(listingapi.ListingsRequest)
	req.Listings.Language = request.Language
	req.Listings.Categories = request.Categories
	req.Listings.Teachers = request.Teachers
	req.Listings.Beings = request.Beings
	req.Listings.Receivers = request.Receivers
	req.Listings.Sections = request.Sections

	res, err := CallOtherServices(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	page, err := AssemblePage(ctx, &req, &res)
	return page, err
}

func (webapiService) TranscriptPage(ctx context.Context, request api.TranscriptRequest) (*api.Page, error) {
	var req PageRequest
	req.Language = request.Language
	req.Banner = new(bannerapi.BannerRequest)
	req.Quote = new(quoteapi.QuoteRequest)
	req.Translations = new(transcriptapi.TranslationsRequest)
	req.Translations.Id = request.Id
	req.Transcript = new(transcriptapi.TranscriptRequest)
	req.Transcript.Id = request.Id

	req.Listings = new(listingapi.ListingsRequest)
	req.Listings.Language = request.Language
	req.Listings.Categories = true
	req.Listings.Teachers = true
	req.Listings.Beings = true
	req.Listings.Receivers = true
	req.Listings.Sections = true

	res, err := CallOtherServices(ctx, req)
	if err != nil {
		return nil, err
	}

	page, err := AssemblePage(ctx, &req, &res)
	return page, err
}

func (webapiService) SearchPage(ctx context.Context, request api.SearchRequest) (*api.Page, error) {
	var req PageRequest
	req.Language = request.Language
	req.Banner = new(bannerapi.BannerRequest)
	req.Quote = new(quoteapi.QuoteRequest)

	// Copy api.SearchRequest into transcriptapi.SearchRequest
	req.Search = new(transcriptapi.SearchRequest)
	req.Search.Visibility = request.Visibility
	req.Search.Status = request.Status
	req.Search.Query = request.Query
	req.Search.Category = request.Category
	req.Search.Group = request.Group
	req.Search.Teacher = request.Teacher
	req.Search.Being = request.Being
	req.Search.Receiver = request.Receiver
	req.Search.StartDate = request.StartDate
	req.Search.EndDate = request.EndDate
	req.Search.Limit = request.Limit
	req.Search.Skip = request.Skip

	req.Listings = new(listingapi.ListingsRequest)
	req.Listings.Language = request.Language
	req.Listings.Categories = true
	req.Listings.Group = true
	req.Listings.Teachers = true
	req.Listings.Beings = true
	req.Listings.Receivers = true

	res, err := CallOtherServices(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	page, err := AssemblePage(ctx, &req, &res)
	return page, err
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for WebApiService.
type ServiceMiddleware func(WebApiService) WebApiService
