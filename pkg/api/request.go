package api

type HomeRequest struct {
	Id string `json:"id"`
}

type SearchRequest struct {
	Language   string `bson:"language"`
	Visibility string `bson:"visibility"`
	Status     string `bson:"status"`
	Query      string `bson:"query"`
	Teacher    string `bson:"teacher"`
	Receiver   string `bson:"receiver"`
	Group      string `bson:"group"`
	Category   string `bson:"category"`
	StartDate  string `bson:"startdate"`
	EndDate    string `bson:"enddate"`
}

type PageRequest struct {
	Banner     *BannerRequest     `bson:"banner"`
	Quote      *QuoteRequest      `bson:"quote"`
	Home       *HomeRequest       `bson:"home"`
	Transcript *TranscriptRequest `bson:"transcript"`
	Search     *SearchRequest     `bson:"search"`
}
