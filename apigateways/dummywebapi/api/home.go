package api

import (
	"context"
)

func HomePage(ctx *context.Context) (article Article) {
	article.Title = "This is sample title"
	article.Body = "Tis is article body"
	return
}

func TranscriptPage(ctx *context.Context) (article Article) {
	article.Title = "This is sample transcript"
	article.Body = "Tis is transcript body"
	return
}
