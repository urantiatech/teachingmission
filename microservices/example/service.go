package main

import (
	"context"
	"errors"
	"strings"
)

type ExampleService interface {
	Example(context.Context, string) (string, error)
}

type exampleService struct{}

func (exampleService) Example(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for ExampleService.
type ServiceMiddleware func(ExampleService) ExampleService
