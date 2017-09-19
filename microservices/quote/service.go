package main

import (
	"context"
	"errors"
	"math/rand"

	"github.com/urantiatech/teachingmission/microservices/quote/api"
	"github.com/urantiatech/teachingmission/microservices/quote/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QuoteService interface {
	Quote(context.Context, api.QuoteRequest) (api.QuoteResponse, error)
}

type quoteService struct{}

func (quoteService) Quote(_ context.Context, req api.QuoteRequest) (api.QuoteResponse, error) {
	var quote api.QuoteResponse

	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(db.Name).C(db.Collection)

	var q db.Quote
	count, err := c.Find(bson.M{"language": req.Language}).Count()
	if err != nil || count == 0 {
		quote.Err = NotFound.Error()
		return quote, nil
	}

	randomId := rand.Intn(count) + 1
	err = c.Find(bson.M{"language": req.Language, "id": randomId}).One(&q)
	if err != nil {
		quote.Err = err.Error()
		return quote, nil
	}

	quote.Id = q.Id
	quote.Language = q.Language
	quote.Teacher = q.Teacher
	quote.Text = q.Text

	return quote, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for QuoteService.
type ServiceMiddleware func(QuoteService) QuoteService
