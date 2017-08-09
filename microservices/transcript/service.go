package main

import (
	"context"
	"errors"
	"time"

	"github.com/urantiatech/teachingmission/pkg/database"
	"github.com/urantiatech/teachingmission/pkg/database/transcript"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dateFormat = "2006-01-02"

type TranscriptService interface {
	Transcript(context.Context, string) (transcript.Transcript, error)
	Search(context.Context, searchRequest) ([]transcript.SearchRecord, error)
}

type transcriptService struct{}

func (transcriptService) Transcript(_ context.Context, id string) (transcript.Transcript, error) {

	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(database.Database).C(transcript.Collection)

	var t transcript.Transcript
	err = c.Find(bson.M{"id": id}).One(&t)
	return t, err
}

func (transcriptService) Search(ctx context.Context, query searchRequest) ([]transcript.SearchRecord, error) {

	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(database.Database).C(transcript.Collection)

	var transcripts []transcript.Transcript

	filters := make(map[string]interface{})

	// Add Language filter
	switch query.Language {
	case "":
		filters["language"] = "en"
	default:
		filters["language"] = query.Language
	}

	// Add Visibility filter
	switch query.Visibility {
	case "any": // No filter
	case "false":
		filters["visibility"] = false
	case "true":
		filters["visibility"] = true
	default:
		filters["visibility"] = true
	}

	// Add Status filter
	switch query.Status {
	case "any": // No filter
	case "":
		filters["status"] = "live"
	default:
		filters["status"] = query.Status
	}

	// Add query filter
	if query.Query != "" {
		filters["$text"] = bson.M{"$search": query.Query}
	}

	// Add Teacher filter
	if query.Teacher != "" {
		filters["teachers.id"] = query.Teacher
	}

	// Add Receiver filter
	if query.Receiver != "" {
		filters["receiver.id"] = query.Receiver
	}

	// Add Group filter
	if query.Group != "" {
		filters["group"] = query.Group
	}

	// Add Category filter
	if query.Category != "" {
		filters["categories.id"] = query.Category
	}

	// Add Date range filter
	var startDate time.Time
	if query.StartDate != "" {
		startDate, err = time.Parse(dateFormat, query.StartDate)
		if err != nil {
			return nil, err
		}
		// Reduce by 1 day to include startDate
		startDate = startDate.Add(-time.Hour * 24)
	}

	var endDate time.Time
	if query.EndDate != "" {
		endDate, err = time.Parse(dateFormat, query.EndDate)
		if err != nil {
			return nil, err
		}
		// Increase by 1 day to include endDate
		endDate = endDate.Add(time.Hour * 24)
	}

	if query.StartDate != "" && query.EndDate != "" {
		// Date range
		filters["date"] = bson.M{"$gt": startDate, "$lt": endDate}
	} else if query.StartDate != "" {
		// Starting date to latest
		filters["date"] = bson.M{"$gt": startDate}
	} else if query.EndDate != "" {
		// Oldest to ending date
		filters["date"] = bson.M{"$lt": endDate}
	}

	err = c.Find(filters).All(&transcripts)
	if len(transcripts) == 0 || err != nil {
		return nil, NotFound
	}

	results, err := Convert(ctx, transcripts)
	return results, err
}

func Convert(_ context.Context, transcripts []transcript.Transcript) ([]transcript.SearchRecord, error) {
	var records []transcript.SearchRecord
	for _, t := range transcripts {
		var r transcript.SearchRecord
		r.Id = t.Id
		r.Language = t.Language
		r.Visibility = t.Visibility
		r.Status = t.Status
		r.Title = t.Title
		r.Date = t.Date
		r.Teachers = t.Teachers
		r.Receiver = t.Receiver
		r.Categories = t.Categories
		r.Group = t.Group
		records = append(records, r)
	}
	return records, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for TranscriptService.
type ServiceMiddleware func(TranscriptService) TranscriptService
