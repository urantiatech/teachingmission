package api

type Definition struct {
	Term  string `json:"term"`
	Value string `json:"value"`
}

type Dictionary struct {
	Definitions []Definition `json:"definitions"`
}
