package api

import (
	transcriptapi "github.com/urantiatech/teachingmission/microservices/transcript/api"
)

type IndexRequest struct {
	Language   string                       `json:"language"`
	Bucket     string                       `json:"bucket"`
	Id         string                       `json:"id"`
	Transcript transcriptapi.TranscriptBody `json:"object"`
}

type IndexResponse struct {
	Object interface{} `json:"object"`
	Err    string      `json:"err,omitempty"`
}
