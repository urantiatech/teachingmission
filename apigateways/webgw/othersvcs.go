package main

import (
	"context"
	"fmt"
	"net/url"

	httptransport "github.com/go-kit/kit/transport/http"
	api "github.com/urantiatech/teachingmission/apigateways/webgw/api"
	articleapi "github.com/urantiatech/teachingmission/microservices/article/api"
	bannerapi "github.com/urantiatech/teachingmission/microservices/banner/api"
	dictionaryapi "github.com/urantiatech/teachingmission/microservices/dictionary/api"
	listingapi "github.com/urantiatech/teachingmission/microservices/listing/api"
	quoteapi "github.com/urantiatech/teachingmission/microservices/quote/api"
	transcriptapi "github.com/urantiatech/teachingmission/microservices/transcript/api"
)

type PageRequest struct {
	Language     string                             `json:"language"`
	Banner       *bannerapi.BannerRequest           `bson:"banner"`
	Quote        *quoteapi.QuoteRequest             `bson:"quote"`
	Article      *articleapi.ArticleRequest         `bson:"article"`
	Listings     *listingapi.ListingsRequest        `bson:"listings"`
	Transcript   *transcriptapi.TranscriptRequest   `bson:"transcript"`
	Translations *transcriptapi.TranslationsRequest `bson:"translations"`
	Search       *transcriptapi.SearchRequest       `bson:"search"`
}

type PageResponse struct {
	Title        string                             `json:"title"`
	Banner       bannerapi.BannerResponse           `bson:"banner"`
	Quote        quoteapi.QuoteResponse             `bson:"quote"`
	Article      articleapi.ArticleResponse         `bson:"article"`
	Listings     listingapi.ListingsResponse        `bson:"listings"`
	Transcript   transcriptapi.TranscriptResponse   `bson:"transcript"`
	Translations transcriptapi.TranslationsResponse `bson:"translations"`
	Search       transcriptapi.SearchResults        `bson:"search"`
	Dictionary   *api.Dictionary                    `json:"dictionary"`
}

func collectDefinitions(inChannel chan dictionaryapi.Definition, outChannel chan api.Definition) {
	dict := make(map[string]string)
	defer close(outChannel)

	// Read all definitions from input Channel
	for d := range inChannel {
		dict[d.Term] = d.Value
	}
	// Writing unique definitions back to the Output channel
	for term, value := range dict {
		outChannel <- api.Definition{Term: term, Value: value}
	}
}

func CallOtherServices(ctx context.Context, req PageRequest) (PageResponse, error) {
	var p PageResponse

	// Read definitions from channel
	inChannel := make(chan dictionaryapi.Definition)
	outChannel := make(chan api.Definition)
	go collectDefinitions(inChannel, outChannel)

	if req.Banner != nil {
		req.Banner.Language = req.Language
		tgt, err := url.Parse("http://localhost:9002/banner")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeBannerResponse).Endpoint()
		resp, err := endPoint(ctx, req.Banner)
		if err != nil {
			p.Banner = bannerapi.BannerResponse{"/images/banners/tm-banner.png", "", ""}
		} else {
			p.Banner = resp.(bannerapi.BannerResponse)
		}
	}

	if req.Quote != nil {
		req.Quote.Language = req.Language
		tgt, err := url.Parse("http://localhost:9004/quote")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeQuoteResponse).Endpoint()
		resp, err := endPoint(ctx, req.Quote)
		if err != nil {
			p.Quote = quoteapi.QuoteResponse{}
		} else {
			p.Quote = resp.(quoteapi.QuoteResponse)
		}
	}

	if req.Article != nil {
		req.Article.Language = req.Language
		tgt, err := url.Parse("http://localhost:9006/article")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeArticleResponse).Endpoint()
		resp, err := endPoint(ctx, req.Article)
		if err != nil {
			p.Article = articleapi.ArticleResponse{}
			p.Title = "Article Not Found"
		} else {
			p.Article = resp.(articleapi.ArticleResponse)
			p.Title = p.Article.Title
			p.Article.Text = AddDefinitions(ctx, p.Article.Text, req.Language, inChannel)
		}
	}

	if req.Listings != nil {
		req.Listings.Language = req.Language
		tgt, err := url.Parse("http://localhost:9012/listings")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt,
			encodeRequest, decodeListingsResponse).Endpoint()
		resp, err := endPoint(ctx, req.Listings)
		if err != nil {
			p.Listings = listingapi.ListingsResponse{}
		} else {
			p.Listings = resp.(listingapi.ListingsResponse)
		}
	}

	if req.Translations != nil {
		tgt, err := url.Parse("http://localhost:9008/translations")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeTranslationsResponse).Endpoint()
		resp, err := endPoint(ctx, req.Translations)
		if err != nil {
			p.Translations = transcriptapi.TranslationsResponse{}
		} else {
			p.Translations = resp.(transcriptapi.TranslationsResponse)
		}
	}

	if req.Transcript != nil {
		req.Transcript.Language = req.Language
		tgt, err := url.Parse("http://localhost:9008/transcript")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeTranscriptResponse).Endpoint()
		resp, err := endPoint(ctx, req.Transcript)
		if err != nil {
			p.Transcript = transcriptapi.TranscriptResponse{}
		} else {
			p.Transcript = resp.(transcriptapi.TranscriptResponse)
			p.Title = p.Transcript.Title

			/* This doesn't work and Unmarshal fails
			// Convert Struct to JSON
			body, err := json.Marshal(p.Transcript.Body)
			if err != nil {
				fmt.Println("Marshal Error: ", err.Error())
			}

			// []byte to string conversion
			text := string(body[:])

			// Add tooltips
			text = AddDefinitions(ctx, text, req.Language)

			// string to []byte conversion
			body = []byte(text)

			// Converto JSON to Struct
			err = json.Unmarshal(body, &p.Transcript.Body)
			if err != nil {
				fmt.Println("Unmarshal Error: ", err.Error())
			}
			*/

			p.Transcript.Body.Html = AddDefinitions(ctx, p.Transcript.Body.Html, req.Language, inChannel)

			for i, para := range p.Transcript.Body.Paragraphs {
				p.Transcript.Body.Paragraphs[i] = AddDefinitions(ctx, para, req.Language, inChannel)
			}

			for i, section := range p.Transcript.Body.Sections {
				for j, para := range section.Paragraphs {
					p.Transcript.Body.Sections[i].Paragraphs[j] = AddDefinitions(ctx, para, req.Language, inChannel)
				}
			}
		}
	}

	if req.Search != nil {
		req.Search.Language = req.Language
		tgt, err := url.Parse("http://localhost:9008/search")
		if err != nil {
			fmt.Println(err.Error())
		}
		endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeSearchResponse).Endpoint()
		resp, err := endPoint(ctx, req.Search)
		if err != nil {
			// Search service not available
			p.Search = transcriptapi.SearchResults{}
			p.Search.Limit = req.Search.Limit
		} else {
			p.Search = resp.(transcriptapi.SearchResults)
			if p.Search.Total == 0 {
				p.Search = transcriptapi.SearchResults{}
				p.Search.Limit = req.Search.Limit
			}
		}
	}

	// Close the input channel
	close(inChannel)

	p.Dictionary = new(api.Dictionary)

	// Add all definitions to the page
	for d := range outChannel {
		p.Dictionary.Definitions = append(p.Dictionary.Definitions, d)
	}

	return p, nil
}
