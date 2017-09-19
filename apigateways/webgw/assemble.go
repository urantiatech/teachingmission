package main

import (
	"context"

	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
)

func AssemblePage(ctx context.Context, req *PageRequest, res *PageResponse) (*api.Page, error) {
	var p *api.Page
	p = new(api.Page)
	p.Title = res.Title

	if checkDictionarySvc(ctx) {
		p.Dictionary = res.Dictionary
	}

	if req.Banner != nil && res.Banner.Err == "" {
		p.Banner = new(api.Banner)
		p.Banner.Url = res.Banner.Url
		p.Banner.Text = res.Banner.Text
	}

	if req.Quote != nil && res.Quote.Err == "" {
		p.Quote = new(api.Quote)
		p.Quote.Text = res.Quote.Text
		p.Quote.Teacher = res.Quote.Teacher
	}

	if req.Article != nil && res.Article.Err == "" {
		p.Article = new(api.Article)
		p.Article.Id = res.Article.Id
		p.Article.Title = res.Article.Title
		p.Article.Text = res.Article.Text
	}

	if req.Listings != nil && res.Listings.Err == "" {
		p.Listings = new(api.Listings)

		// Copy Categories
		p.Listings.Categories.Count = res.Listings.Categories.Count
		for _, c := range res.Listings.Categories.Categories {
			category := api.Category{
				Id:          c.Id,
				Name:        c.Name,
				ParentId:    c.ParentId,
				Description: c.Description,
			}
			p.Listings.Categories.Categories = append(p.Listings.Categories.Categories, category)
		}

		// Copy Sections
		p.Listings.Sections.Count = res.Listings.Sections.Count
		for _, c := range res.Listings.Sections.Sections {
			section := api.Section{
				Id:          c.Id,
				Name:        c.Name,
				ParentId:    c.ParentId,
				Description: c.Description,
			}
			p.Listings.Sections.Sections = append(p.Listings.Sections.Sections, section)
		}

		// Copy Teachers
		p.Listings.Teachers.Count = res.Listings.Teachers.Count
		for _, t := range res.Listings.Teachers.Teachers {
			teacher := api.Teacher{
				Id:         t.Id,
				Name:       t.Name,
				OtherNames: t.OtherNames,
				Being:      t.BeingId,
				Profile:    t.Profile,
			}
			p.Listings.Teachers.Teachers = append(p.Listings.Teachers.Teachers, teacher)

		}

		// Copy Beings
		p.Listings.Beings.Count = res.Listings.Beings.Count
		for _, b := range res.Listings.Beings.Beings {
			being := api.Being{
				Id:          b.Id,
				Name:        b.Name,
				Description: b.Description,
			}
			p.Listings.Beings.Beings = append(p.Listings.Beings.Beings, being)
		}

		// Copy Receivers
		p.Listings.Receivers.Count = res.Listings.Receivers.Count
		for _, r := range res.Listings.Receivers.Receivers {
			receiver := api.Receiver{
				Id:      r.Id,
				Name:    r.Name,
				Profile: r.Profile,
			}
			p.Listings.Receivers.Receivers = append(p.Listings.Receivers.Receivers, receiver)
		}

	}

	if req.Translations != nil && res.Translations.Err == "" {
		p.Translations = new(api.Translations)
		p.Translations.Id = res.Translations.Id

		for _, f := range res.Translations.Fields {
			header := api.TranscriptHeader{Language: f.Language, Title: f.Title,
				Visibility: f.Visibility, Status: f.Status}
			p.Translations.Headers = append(p.Translations.Headers, header)
		}
	}

	if req.Transcript != nil && res.Transcript.Err == "" {
		p.Transcript = new(api.Transcript)
		p.Transcript.Id = res.Transcript.Id
		p.Transcript.Title = res.Transcript.Title

		for _, category := range res.Transcript.Categories {
			p.Transcript.Categories = append(p.Transcript.Categories,
				api.Category{
					Id:          category.Id,
					Name:        category.Name,
					ParentId:    "",
					Description: "",
				})
		}

		p.Transcript.Date = res.Transcript.Date
		p.Transcript.Location = res.Transcript.Location

		for _, teacher := range res.Transcript.Teachers {
			p.Transcript.Teachers = append(p.Transcript.Teachers,
				api.Teacher{
					Id:   teacher.Id,
					Name: teacher.Name,
				})
		}

		p.Transcript.Receiver = api.Person{
			Id:   res.Transcript.Receiver.Id,
			Name: res.Transcript.Receiver.Name,
		}

		for _, attendee := range res.Transcript.Attendees {
			p.Transcript.Attendees = append(p.Transcript.Attendees,
				api.Person{
					Id:   attendee.Id,
					Name: attendee.Name,
				})
		}

		// Copy Body
		p.Transcript.Body.Html = res.Transcript.Body.Html
		p.Transcript.Body.Paragraphs = res.Transcript.Body.Paragraphs

		for _, section := range res.Transcript.Body.Sections {
			p.Transcript.Body.Sections = append(p.Transcript.Body.Sections,
				api.TranscriptSection{
					TocTitle:   section.Toc,
					Heading:    section.Heading,
					Paragraphs: section.Paragraphs,
				})
		}
	}

	if req.Search != nil && res.Search.Err == "" {
		p.Search = new(api.Search)
		p.Search.Query = req.Search.Query

		p.Search.Filters.Category = req.Search.Category
		p.Search.Filters.Group = req.Search.Group
		p.Search.Filters.Teacher = req.Search.Teacher
		p.Search.Filters.Being = req.Search.Being
		p.Search.Filters.Receiver = req.Search.Receiver
		p.Search.Filters.StartDate = req.Search.StartDate
		p.Search.Filters.EndDate = req.Search.EndDate

		p.Search.Skipped = res.Search.Skip
		p.Search.Limit = res.Search.Limit
		p.Search.Total = res.Search.Total

		for _, r := range res.Search.Results {
			result := api.InternalResult{}
			result.Id = r.Id
			result.Title = r.Title

			for _, t := range r.Teachers {
				teacher := api.Teacher{}
				teacher.Id = t.Id
				teacher.Name = t.Name
				result.Teachers = append(result.Teachers, teacher)
			}

			p.Search.Results = append(p.Search.Results, result)
		}
	}

	return p, nil
}

/*
func HomePage(ctx context.Context) *Page {
	var p *Page = new(Page)
	//lang, err := ContextVal(ctx, "lang")
	p.Title = "Home Page Title"
	p.Meta = &Meta{Keywords: "Teaching Mission", Description: "Page meta description"}
	p.User = &User{Username: "sangeetk", Email: "sangeet.kumar@gmail.com", Role: "Admin"}
	p.Banner = &Banner{Url: "/images/banner1.jpg", Text: "This is banner text"}
	p.Quote = &Quote{Teacher: "Thought Adjuster", Text: "I AM the love of the Creator indwelling your mind. I AM close to you always and I know everything about you—your every thought, your most secret of desires, and your imperfections. I love you unconditionally, for I AM your true love."}
	p.Article = &Article{
		Title: "Welcome to Teaching Mission",
		Text: `<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum. </p>

<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum. </p>

<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum. </p>

<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>`,
	}
	return p
}
*/

/*
func SearchPage(ctx context.Context, keywords string) *Page {
	var p *Page = new(Page)
	//lang, err := ContextVal(ctx, "lang")
	p.Title = "Search Page Title"
	p.Meta = &Meta{Keywords: "Teaching Mission", Description: "Page meta description"}
	p.User = &User{Username: "sangeetk", Email: "sangeet.kumar@gmail.com", Role: "Admin"}
	p.Banner = &Banner{Url: "/images/banner2.jpg"}
	//p.Quote = &Quote{Teacher: "Teacher Ophelius", Text: "God has given us the opportunity to explore and co-create the universes through mind and memory, for He has provided the canvas of love and the primary colors of truth, beauty, and goodness to paint our masterpiece."}
	p.Search = &Search{
		Keywords: keywords,
		Filters:  Filter{Category: Category{Id: "love", Name: "Love"}, Teacher: "Monjoronson", Receiver: "Chris Maurus", Team: "11:11 Progress Group"},
		Results: []InternalResult{
			{Title: "Transcript 1", Teacher: "Teacher Ophelius, Christ Michael", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 2", Teacher: "Monjoronson", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 3 This is another big title", Teacher: "Unknown", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 4", Teacher: "Thought Adjuster", Uri: "/transcript", Excerpts: "... quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 5", Uri: "/transcript", Excerpts: "... Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 6", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 7", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 8", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 9", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 10", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 11", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
			{Title: "Transcript 20", Uri: "/transcript", Excerpts: "... Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."},
		},
		ExternalResults: []ExternalResult{
			{Title: "A Lesson on Spirit Communication", Url: "http://correctingtime.dev/ct-messages/lesson-spirit-communication"},
			{Title: "Prepare yourselves now for your Next Life", Url: "http://correctingtime.dev/ct-messages/prepare-yourselves-now-your-next-life"},
			{Title: "We have entered a Turning Point", Url: "http://correctingtime.dev/ct-messages/we-have-entered-turning-point"},
			{Title: "For Ye are Gods", Url: "http://correctingtime.dev/ct-messages/%E2%80%9C-ye-are-gods%E2%80%9D"},
		},
	}
	return p
}
*/

/*
func TranscriptPage(ctx context.Context) *Page {
	var p *Page = new(Page)
	//lang, err := ContextVal(ctx, "lang")
	p.Title = "Transcript Page Title"
	p.Meta = &Meta{Keywords: "Teaching Mission", Description: "Page meta description"}
	p.User = &User{Username: "sangeetk", Email: "sangeet.kumar@gmail.com", Role: "Admin"}
	p.Banner = &Banner{Url: "/images/banner1.jpg", Text: "This is banner text"}
	p.Quote = &Quote{Teacher: "Teacher Ophelius", Text: "Each one of you on the 11:11 list is known to us, and we ask each of you to pay close attention to the inner life and the inner voice as you will be guided through the many changes that are now taking place and which are about to occur in the near future as the Correcting Time mandates are executed.You are to be the anchors of light and a voice of hope in this world as these many changes come upon you. Stay grounded my friends, and listen for direction from on high, for we are with you and around you at all times ready to serve those who align themselves with the forces of light and those who do the Will of our paradise Father."}
	p.Transcript = &Transcript{
		Id:       "a-lesson-on-love",
		Title:    "A Lesson on Love",
		Category: Category{Id: "love", Name: "Love"},
		Date:     "2009-04-17",
		Location: "Michigan, US of A",
		Teacher:  []Teacher{{Id: "teacher-ophelius", Name: "Teacher Ophelius", Type: "Mansion World Teacher"}},
		Receiver: Person{Id: "chris-maurus", Name: "Chris Maurus"},
		Body: []Section{
			{
				Paragraph: []string{
					`Today’s lesson is on love. There is much to say about love, for it is the very essence of all things, and it truly could not be described in its entirety in words, songs, or the arts, yet it is the one thing that inspires truth, beauty, and goodness. Love is as infinite as the Creator himself—it is the only real thing of value, the very construct of all reality and relationships between personalities. Love is all that truly matters, for everything else will pass away.`,
					`What is it that we are truly searching for today my beloved? What is it that we busy our selves with, going to and fro, here and there, learning, working, playing, creating, building our careers? We do all these things day in and day out, and the years pass by very quickly, but then do we stop for a moment and take inventory of our past? For when we do, what things boil up to the surface—the things of love in our life? They are those precious moments that we cherish—your first kiss; those intimate relationships with individuals that touched you on a deep and profound level; the kindness of strangers that helped you when you were in need; holding your newborn child for the first time; on and on goes the mind as it recalls the memory of the highest value meanings and feelings in your life.`,
					`These are the things of love that will survive eternity—these are survival value soul building moments and they become a part of your morontia self. Everything else will pass away. So you see how important these things are and yet they make up a very small amount of your living moments in life. The illusion of the material world and the pursuit of material things distracts us from the most important things in life—the pursuit of love; the giving and the receiving of love; to make a connection with someone on a soul level; to link hearts with your brothers and sisters in unconditional love.`,
					`Step outside of yourself right now and take inventory of the important things of love in your life. Think about how you can make a difference and become one of those cherished memories in the life of one of your brothers or sisters. How beautiful it is to be apart of the soul memory of another human being and live forever in their eternal memory. Reach out and touch someone today my beloved. This is the very meaning of what the master meant when he said, “Save your treasures in heaven, for there also will be your heart.”`,
					"Good day",
					"The Circle of Seven",
				},
			},
		},
	}
	return p
}
*/
