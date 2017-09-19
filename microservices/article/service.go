package main

import (
	"context"
	"errors"

	"github.com/urantiatech/teachingmission/microservices/article/api"
	"github.com/urantiatech/teachingmission/microservices/article/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ArticleService interface {
	Article(context.Context, api.ArticleRequest) (api.ArticleResponse, error)
}

type articleService struct{}

func (articleService) Article(_ context.Context, req api.ArticleRequest) (api.ArticleResponse, error) {

	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(db.Name).C(db.Collection)

	var q db.Article
	err = c.Find(bson.M{"language": req.Language, "id": req.Id}).One(&q)

	var article api.ArticleResponse
	if err != nil {
		article.Err = err.Error()
		return article, nil
	}

	article.Id = q.Id
	article.Title = q.Title
	article.Text = q.Text

	return article, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for ArticleService.
type ServiceMiddleware func(ArticleService) ArticleService
