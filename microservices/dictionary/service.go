package main

import (
	"context"
	"errors"

	"github.com/urantiatech/teachingmission/microservices/dictionary/api"
)

const defaultTextTag = "tm-dict"
const toolTag = "abbr"
const defaultTaptargetTag = "tm-taptarget"

const textSvc = "/text"
const tooltipSvc = "/tooltip"
const taptargetSvc = "/taptarget"

type DefinitionService interface {
	Definition(context.Context, api.DefinitionRequest) (api.DefinitionResponse, error)
	Tooltip(context.Context, api.TooltipRequest) (api.TooltipResponse, error)
	Taptarget(context.Context, api.TaptargetRequest) (api.TaptargetResponse, error)
	Text(context.Context, api.TextRequest) (api.TextResponse, error)
}

type definitionService struct{}

func (definitionService) Definition(ctx context.Context, request api.DefinitionRequest) (api.DefinitionResponse, error) {
	var definition api.DefinitionResponse
	definition.Term = request.Term

	// Load the dictionary from database if not present
	if Dictionary[request.Language] == nil {
		dictionary, err := initDictionary(ctx, request.Language)
		if err == nil {
			Dictionary[request.Language] = dictionary
		} else {
			definition.Err = NotFound.Error()
			return definition, nil
		}
	}

	// Check for the definition from dictionary
	value, err := Dictionary[request.Language].Value(request.Term)
	if err != nil {
		definition.Err = NotFound.Error()
		return definition, nil
	}

	// Return ddefinition
	definition.Definition = value
	return definition, nil
}

func (definitionService) Tooltip(ctx context.Context, request api.TooltipRequest) (api.TooltipResponse, error) {
	var resp api.TooltipResponse

	// Load the dictionary from database if not present
	if Dictionary[request.Language] == nil {
		dictionary, err := initDictionary(ctx, request.Language)
		if err == nil {
			Dictionary[request.Language] = dictionary
		}
	}

	// Add tooltips to the text
	resp.Text, resp.Count, _ = insertDefinition(ctx, request.Text, tooltipSvc, toolTag, request.Language, request.Css, request.Duplicate, true)
	return resp, nil
}

func (definitionService) Taptarget(ctx context.Context, request api.TaptargetRequest) (api.TaptargetResponse, error) {
	var resp api.TaptargetResponse
	var tag string

	if request.Tag != "" {
		tag = request.Tag
	} else {
		tag = defaultTaptargetTag
	}

	// Load the dictionary from database if not present
	if Dictionary[request.Language] == nil {
		dictionary, err := initDictionary(ctx, request.Language)
		if err == nil {
			Dictionary[request.Language] = dictionary
		}
	}

	// Add tooltips to the text
	resp.Text, resp.Count, resp.Dictionary = insertDefinition(ctx, request.Text, taptargetSvc, tag, request.Language, request.Css, request.Duplicate, request.Value)
	return resp, nil
}

func (definitionService) Text(ctx context.Context, request api.TextRequest) (api.TextResponse, error) {
	var resp api.TextResponse
	var tag string

	if request.Tag != "" {
		tag = request.Tag
	} else {
		tag = defaultTextTag
	}

	// Load the dictionary from database if not present
	if Dictionary[request.Language] == nil {
		dictionary, err := initDictionary(ctx, request.Language)
		if err == nil {
			Dictionary[request.Language] = dictionary
		}
	}

	// Add custom tags to the text
	resp.Text, resp.Count, _ = insertDefinition(ctx, request.Text, textSvc, tag, request.Language, request.Css, request.Duplicate, request.Value)
	return resp, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for DefinitionService.
type ServiceMiddleware func(DefinitionService) DefinitionService
