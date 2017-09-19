package main

import (
	"context"
	"errors"

	"github.com/urantiatech/teachingmission/microservices/listing/api"
	"github.com/urantiatech/teachingmission/microservices/listing/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ListingService interface {
	Categories(context.Context, api.CategoriesRequest) (api.CategoriesResponse, error)
	Sections(context.Context, api.SectionsRequest) (api.SectionsResponse, error)
	Groups(context.Context, api.GroupsRequest) (api.GroupsResponse, error)
	Teachers(context.Context, api.TeachersRequest) (api.TeachersResponse, error)
	Beings(context.Context, api.BeingsRequest) (api.BeingsResponse, error)
	Receivers(context.Context, api.ReceiversRequest) (api.ReceiversResponse, error)
	Listings(context.Context, api.ListingsRequest) (api.ListingsResponse, error)
}

type listingService struct{}

func (listingService) Categories(ctx context.Context, req api.CategoriesRequest) (api.CategoriesResponse, error) {
	var categories api.CategoriesResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.Name != "" {
		filters["name"] = req.Name
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.CategoriesCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		categories.Count = 0
		categories.Err = NotFound.Error()
		return categories, nil
	}

	categories.Count = count

	// Get records from DB
	var dbCategories []db.Category
	err = collection.Find(filters).All(&dbCategories)
	for _, c := range dbCategories {
		category := api.Category{
			Id:          c.Id,
			Name:        c.Name,
			ParentId:    c.ParentId,
			Description: c.Description,
		}
		categories.Categories = append(categories.Categories, category)
	}

	return categories, nil
}

func (listingService) Sections(ctx context.Context, req api.SectionsRequest) (api.SectionsResponse, error) {
	var sections api.SectionsResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.Name != "" {
		filters["name"] = req.Name
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.SectionsCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		sections.Count = 0
		sections.Err = NotFound.Error()
		return sections, nil
	}

	sections.Count = count

	// Get records from DB
	var dbSections []db.Section
	err = collection.Find(filters).All(&dbSections)
	for _, c := range dbSections {
		section := api.Section{
			Id:          c.Id,
			Name:        c.Name,
			ParentId:    c.ParentId,
			Description: c.Description,
		}
		sections.Sections = append(sections.Sections, section)
	}

	return sections, nil
}

func (listingService) Groups(ctx context.Context, req api.GroupsRequest) (api.GroupsResponse, error) {
	var groups api.GroupsResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.Name != "" {
		filters["name"] = req.Name
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.GroupsCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		groups.Count = 0
		groups.Err = NotFound.Error()
		return groups, nil
	}

	groups.Count = count

	// Get records from DB
	var dbGroups []db.Group
	err = collection.Find(filters).All(&dbGroups)
	for _, c := range dbGroups {
		category := api.Group{
			Id:          c.Id,
			Name:        c.Name,
			Description: c.Description,
		}
		groups.Groups = append(groups.Groups, category)
	}

	return groups, nil
}

func (listingService) Teachers(ctx context.Context, req api.TeachersRequest) (api.TeachersResponse, error) {
	var teachers api.TeachersResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.BeingId != "" {
		filters["beingid"] = req.BeingId
	}

	// Search name in other names too
	if req.Name != "" {
		filters["$text"] = bson.M{"$search": req.Name}
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.TeachersCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		teachers.Count = 0
		teachers.Err = NotFound.Error()
		return teachers, nil
	}

	teachers.Count = count

	// Get records from DB
	var dbTeachers []db.Teacher
	err = collection.Find(filters).All(&dbTeachers)
	for _, c := range dbTeachers {
		teacher := api.Teacher{
			Id:         c.Id,
			Name:       c.Name,
			OtherNames: c.OtherNames,
			BeingId:    c.BeingId,
			Profile:    c.Profile,
		}
		teachers.Teachers = append(teachers.Teachers, teacher)
	}

	return teachers, nil
}

func (listingService) Beings(ctx context.Context, req api.BeingsRequest) (api.BeingsResponse, error) {
	var beings api.BeingsResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.Name != "" {
		filters["name"] = req.Name
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.BeingsCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		beings.Count = 0
		beings.Err = NotFound.Error()
		return beings, nil
	}

	beings.Count = count

	// Get records from DB
	var dbBeings []db.Being
	err = collection.Find(filters).All(&dbBeings)
	for _, c := range dbBeings {
		being := api.Being{
			Id:          c.Id,
			Name:        c.Name,
			Description: c.Description,
		}
		beings.Beings = append(beings.Beings, being)
	}

	return beings, nil
}

func (listingService) Receivers(ctx context.Context, req api.ReceiversRequest) (api.ReceiversResponse, error) {
	var receivers api.ReceiversResponse

	filters := make(map[string]interface{})
	filters["language"] = req.Language
	if req.Id != "" {
		filters["id"] = req.Id
	}
	if req.Name != "" {
		filters["name"] = req.Name
	}

	session, err := mgo.Dial("localhost")
	defer session.Close()

	collection := session.DB(db.Name).C(db.ReceiversCollection)

	// Get the count
	count, err := collection.Find(filters).Count()
	if err != nil || count == 0 {
		receivers.Count = 0
		receivers.Err = NotFound.Error()
		return receivers, nil
	}

	receivers.Count = count

	// Get records from DB
	var dbReceivers []db.Receiver
	err = collection.Find(filters).All(&dbReceivers)
	for _, c := range dbReceivers {
		teacher := api.Receiver{
			Id:      c.Id,
			Name:    c.Name,
			Profile: c.Profile,
		}
		receivers.Receivers = append(receivers.Receivers, teacher)
	}

	return receivers, nil
}

func (listingService) Listings(ctx context.Context, req api.ListingsRequest) (api.ListingsResponse, error) {
	var listings api.ListingsResponse
	svc := listingService{}

	// Request Categories listing
	if req.Categories {
		req := api.CategoriesRequest{Language: req.Language}
		listings.Categories, _ = svc.Categories(ctx, req)
	}

	// Request Sections listing
	if req.Sections {
		req := api.SectionsRequest{Language: req.Language}
		listings.Sections, _ = svc.Sections(ctx, req)
	}

	// Request Teachers listing
	if req.Teachers {
		req := api.TeachersRequest{Language: req.Language}
		listings.Teachers, _ = svc.Teachers(ctx, req)
	}

	// Request Beings listing
	if req.Beings {
		req := api.BeingsRequest{Language: req.Language}
		listings.Beings, _ = svc.Beings(ctx, req)
	}

	// Request Receivers listing
	if req.Receivers {
		req := api.ReceiversRequest{Language: req.Language}
		listings.Receivers, _ = svc.Receivers(ctx, req)
	}
	return listings, nil
}

var NotFound = errors.New("Not Found")

// ServiceMiddleware is a chainable behavior modifier for ListingService.
type ServiceMiddleware func(ListingService) ListingService
