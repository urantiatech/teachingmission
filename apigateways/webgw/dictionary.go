package main

import (
	"context"
	"fmt"
	"net/url"

	httptransport "github.com/go-kit/kit/transport/http"
	dictionaryapi "github.com/urantiatech/teachingmission/microservices/dictionary/api"
)

// Wrapper for easy swapping between tooltips & taptargets
func AddDefinitions(ctx context.Context, text string, language string, c chan dictionaryapi.Definition) string {
	return AddTaptargets(ctx, text, language, c)
}

func AddTooltips(ctx context.Context, text string, language string, _ chan dictionaryapi.Definition) string {
	var req dictionaryapi.TooltipRequest
	req.Language = language
	req.Text = text
	req.Duplicate = false

	tgt, err := url.Parse("http://localhost:9010/tooltip")
	if err != nil {
		fmt.Println(err.Error())
	}
	endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeTooltipResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return text
	} else {
		response := resp.(dictionaryapi.TooltipResponse)
		return response.Text
	}
}

func AddTaptargets(ctx context.Context, text string, language string, c chan dictionaryapi.Definition) string {
	var req dictionaryapi.TaptargetRequest
	req.Language = language
	req.Tag = "ub-term"
	req.Css = "ub-term"
	req.Text = text
	req.Value = false
	req.Duplicate = false

	tgt, err := url.Parse("http://localhost:9010/taptarget")
	if err != nil {
		fmt.Println(err.Error())
	}
	endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeTaptargetResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return text
	} else {
		response := resp.(dictionaryapi.TaptargetResponse)
		// Send all definitions to channel

		for _, definition := range response.Dictionary.Definitions {
			c <- definition
		}
		return response.Text
	}
}

func checkDictionarySvc(ctx context.Context) bool {
	var req dictionaryapi.TooltipRequest
	req.Language = "en"
	req.Text = "test"
	req.Duplicate = false

	tgt, err := url.Parse("http://localhost:9010/tooltip")
	if err != nil {
		fmt.Println(err.Error())
	}
	endPoint := httptransport.NewClient("POST", tgt, encodeRequest, decodeTooltipResponse).Endpoint()
	_, err = endPoint(ctx, req)
	if err != nil {
		return false
	}
	return true
}
