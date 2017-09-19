package main

import (
	"context"
	"errors"

	"github.com/urantiatech/teachingmission/microservices/banner/api"
)

type BannerService interface {
	Banner(context.Context, api.BannerRequest) (api.BannerResponse, error)
}

type bannerService struct{}

func (bannerService) Banner(_ context.Context, req api.BannerRequest) (api.BannerResponse, error) {
	lang := req.Language
	var image string
	if lang != "" {
		image = "/images/banners/tm-banner-" + lang + ".png"
	} else {
		image = "/images/banners/tm-banner.png"
	}

	banner := api.BannerResponse{image, "", ""}
	return banner, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for BannerService.
type ServiceMiddleware func(BannerService) BannerService
