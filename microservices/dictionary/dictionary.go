package main

import (
	"context"
	"time"

	"github.com/urantiatech/pkg/trie"
	"github.com/urantiatech/teachingmission/microservices/dictionary/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Dictionary map[string]*trie.Trie

func init() {
	Dictionary = make(map[string]*trie.Trie)
}

func initDictionary(ctx context.Context, language string) (*trie.Trie, error) {
	// Initilize the dictionary
	dict, err := loadDictionary(ctx, language)

	// Setup auto-refresh of dictionary terms
	ticker := time.NewTicker(1 * time.Hour)
	go reloadDictionary(ctx, language, ticker)

	return dict, err
}

func loadDictionary(ctx context.Context, language string) (*trie.Trie, error) {
	var dict *trie.Trie

	// Connect to the database
	session, err := mgo.Dial("localhost")
	if err != nil {
		return dict, NotFound
	}
	defer session.Close()

	c := session.DB(db.Name).C(db.Collection)

	// Find all the terms of the given language
	var definitions []db.Definition
	err = c.Find(bson.M{"language": language}).All(&definitions)
	if len(definitions) == 0 || err != nil {
		return dict, NotFound
	}

	// Initialize a new trie and add definitions
	dict = trie.New()
	for _, d := range definitions {
		dict.Insert(d.Term, d.Definition)
	}

	// Return dictionary
	return dict, nil
}

func reloadDictionary(ctx context.Context, language string, ticker *time.Ticker) {

	for range ticker.C {
		dict, err := loadDictionary(ctx, language)
		if err == nil {
			Dictionary[language] = dict
		}
	}
}
