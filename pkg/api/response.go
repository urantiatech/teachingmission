package api

import (
	"github.com/urantiatech/teachingmission/pkg/database/quote"
	"github.com/urantiatech/teachingmission/pkg/database/transcript"
)

type HomeResponse struct {
	Id string `json:"id"`
}

type PageResponse struct {
	Banner     *BannerResponse     `bson:"banner"`
	Quote      *QuoteResponse      `bson:"quote"`
	Home       *HomeResponse       `bson:"home"`
	Transcript *TranscriptResponse `bson:"transcript"`
	Search     *SearchResponse     `bson:"search"`
	Err        string              `json:"err,omitempty"`
}
