package main

import (
	"context"
	"errors"
	"time"

	"github.com/urantiatech/teachingmission/microservices/transcript/api"
	"github.com/urantiatech/teachingmission/microservices/transcript/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dateFormat = "2006-01-02"

type TranscriptService interface {
	Transcript(context.Context, api.TranscriptRequest) (api.TranscriptResponse, error)
	Translations(context.Context, api.TranslationsRequest) (api.TranslationsResponse, error)
	Search(context.Context, api.SearchRequest) (api.SearchResults, error)
}

type transcriptService struct{}

func (transcriptService) Transcript(ctx context.Context, req api.TranscriptRequest) (api.TranscriptResponse, error) {
	var response api.TranscriptResponse
	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(db.Name).C(db.Collection)

	var t db.Transcript
	err = c.Find(bson.M{"language": req.Language, "id": req.Id}).One(&t)

	if err != nil {
		response.Err = err.Error()
		return response, nil
	}
	// Convert db.Transcript to api.TranscriptResposne
	response = copyTranscript(ctx, &t)
	return response, nil
}

func copyTranscript(ctx context.Context, t *db.Transcript) api.TranscriptResponse {
	var r api.TranscriptResponse
	r.Id = t.Id
	r.Slug = t.Slug
	r.TrSessionId = t.TrSessionId
	r.Title = t.Title
	r.Visibility = t.Visibility
	r.Status = t.Status
	r.Language = t.Language
	r.Date = t.Date
	r.Location = t.Location

	for _, teacher := range t.Teachers {
		r.Teachers = append(r.Teachers, api.Teacher{teacher.Id, teacher.Name})
	}

	r.Receiver = api.Person{t.Receiver.Id, t.Receiver.Name}
	r.Translator = api.Person{t.Translator.Id, t.Translator.Name}

	for _, attendee := range t.Attendees {
		r.Attendees = append(r.Attendees, api.Person{attendee.Id, attendee.Name})
	}

	r.Group.Id = t.Group.Id
	r.Group.Name = t.Group.Name

	for _, category := range t.Categories {
		r.Categories = append(r.Categories,
			api.Category{category.Id, category.Name})
	}

	r.ImportDomain = t.ImportDomain
	r.ImportUrl = t.ImportUrl
	r.ImportDate = t.ImportDate
	r.LastUpdate = t.LastUpdate

	r.Body.Html = t.Body.Html
	r.Body.Paragraphs = t.Body.Paragraphs

	for _, section := range t.Body.Sections {
		var s api.TranscriptSection
		s.Toc = section.Toc
		s.Heading = section.Heading
		s.Paragraphs = section.Paragraphs
		r.Body.Sections = append(r.Body.Sections, s)
	}

	return r
}

func (transcriptService) Translations(ctx context.Context, req api.TranslationsRequest) (api.TranslationsResponse, error) {
	var response api.TranslationsResponse
	session, err := mgo.Dial("localhost")
	defer session.Close()

	response.Id = req.Id
	c := session.DB(db.Name).C(db.Collection)

	var allTranscripts []db.Transcript
	err = c.Find(bson.M{"id": req.Id}).All(&allTranscripts)
	if err != nil || len(allTranscripts) == 0 {
		response.Err = NotFound.Error()
		return response, nil
	}

	for _, t := range allTranscripts {
		fields := api.TranslationsHeaderFields{Language: t.Language, Title: t.Title, Visibility: t.Visibility, Status: t.Status}
		response.Fields = append(response.Fields, fields)
	}
	return response, nil
}

func (transcriptService) Search(ctx context.Context, query api.SearchRequest) (api.SearchResults, error) {
	var results api.SearchResults

	session, err := mgo.Dial("localhost")
	defer session.Close()

	c := session.DB(db.Name).C(db.Collection)

	var transcripts []db.Transcript

	filters := make(map[string]interface{})

	// Add Language filter
	filters["language"] = query.Language

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

	// Add Category filter
	if query.Category != "" {
		filters["categories.id"] = query.Category
	}

	// Add Group filter
	if query.Group != "" {
		filters["group.id"] = query.Group
	}

	// Add Teacher filter
	if query.Teacher != "" {
		filters["teachers.id"] = query.Teacher
	}

	// Add Receiver filter
	if query.Receiver != "" {
		filters["receiver.id"] = query.Receiver
	}

	// Add Date range filter
	var startDate time.Time
	if query.StartDate != "" {
		startDate, err = time.Parse(dateFormat, query.StartDate)
		if err != nil {
			results.Err = err.Error()
			return results, nil
		}
		// Reduce by 1 day to include startDate
		startDate = startDate.Add(-time.Hour * 24)
	}

	var endDate time.Time
	if query.EndDate != "" {
		endDate, err = time.Parse(dateFormat, query.EndDate)
		if err != nil {
			results.Err = err.Error()
			return results, nil
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

	// Add stats
	results.Skip = query.Skip
	results.Limit = query.Limit
	results.Total, err = c.Find(filters).Count()
	if err != nil {
		results.Total = 0
		results.Err = err.Error()
		return results, nil
	}

	// Add Being filter
	if query.Being != "" {
		// Apply Skip and Limit at the time of applying being filter
		err = c.Find(filters).All(&transcripts)

		// Get all Teachers from listing service.
		teachers := *getTeachers(ctx, query.Language, query.Being)

		var tmp []db.Transcript
		// Filter transcripts where teacher.BeingId != query.Being
		for _, t := range transcripts {
			for _, tchr := range t.Teachers {
				if _, ok := teachers[tchr.Id]; ok {
					tmp = append(tmp, t)
					break
				}
			}
		}

		// Calculate Total Results
		results.Total = len(tmp)
		if results.Total == 0 {
			transcripts = nil
		} else {
			last := results.Skip + results.Limit
			if last > results.Total {
				last = results.Total
			}
			transcripts = tmp[results.Skip:last]
		}
	} else {
		// No Being filter set.
		err = c.Find(filters).Skip(query.Skip).Limit(query.Limit).All(&transcripts)
	}

	if len(transcripts) == 0 || err != nil {
		results.Err = NotFound.Error()
		return results, nil
	}

	results.Results, err = copySearchResults(ctx, transcripts)
	if err != nil {
		results.Total = 0
		results.Err = err.Error()
	}
	return results, nil

}

func copySearchResults(_ context.Context, transcripts []db.Transcript) ([]api.SearchRecord, error) {
	var records []api.SearchRecord
	for _, t := range transcripts {
		var r api.SearchRecord
		r.Id = t.Id
		r.Visibility = t.Visibility
		r.Status = t.Status
		r.Title = t.Title
		r.Date = t.Date

		for _, teacher := range t.Teachers {
			r.Teachers = append(r.Teachers, api.Teacher{teacher.Id, teacher.Name})
		}

		r.Receiver = api.Person{t.Receiver.Id, t.Receiver.Name}

		for _, category := range t.Categories {
			r.Categories = append(r.Categories,
				api.Category{category.Id, category.Name})
		}

		r.Group.Id = t.Group.Id
		r.Group.Name = t.Group.Name

		records = append(records, r)
	}
	return records, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for TranscriptService.
type ServiceMiddleware func(TranscriptService) TranscriptService
