package main

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/urantiatech/pkg/trie"
	"github.com/urantiatech/teachingmission/microservices/dictionary/api"
)

func slug(title string) string {
	// Filter and conver to lowercase
	slug := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 'a' - 'A'
		case r >= 'a' && r <= 'z':
			return r
		case r >= '0' && r <= '9':
			return r
		}
		return ' '
	}

	// Convert whitespace to hyphen '-'
	str := strings.Map(slug, title)
	strarray := strings.Fields(str)
	str = strings.Join(strarray, "-")

	return str
}

func isLetter(word string) bool {
	for _, r := range word {
		return unicode.IsLetter(r)
	}
	return false
}

func insertTextTags(tag, key, val, css string, valFlag bool) (string, string) {
	var startTag, endTag string

	if valFlag {
		startTag = fmt.Sprintf("<%s class=\"%s\" title=\"%s\">", tag, css, val)
	} else {
		startTag = "<" + tag + ">"
	}
	endTag = "</" + tag + ">"
	return startTag, endTag
}

func insertTooltipTags(tag, key, val, css string, valFlag bool) (string, string) {
	var startTag, endTag string

	if valFlag {
		startTag = fmt.Sprintf("<%s class=\"%s\" title=\"%s\">", tag, css, val)
	} else {
		startTag = "<" + tag + ">"
	}
	endTag = "</" + tag + ">"
	return startTag, endTag
}

func insertTaptargetTags(tag, key, val, css string, valFlag bool) (string, string) {
	var startTag, endTag string

	// <tag onclick="$('.thought-adjuster').tapTarget('open')">Thought Adjuster</tag>

	if valFlag {
		startTag = fmt.Sprintf("<%s class=\"%s\" onclick=\"$('.%s').tapTarget('open')\" value=\"%s\">", tag, css, slug(key), val)
	} else {
		startTag = fmt.Sprintf("<%s class=\"%s\" onclick=\"$('.%s').tapTarget('open')\">", tag, css, slug(key))
	}
	endTag = "</" + tag + ">"
	return startTag, endTag
}

func insertDefinition(ctx context.Context, text, svc, tag, lang, css string,
	duplicate, valFlag bool) (string, int, *api.Dictionary) {
	// Split the text into ([]*[w])* array of words and non-words
	type word []rune
	var sentence []string
	var prev rune
	var count int
	var dict api.Dictionary

	for _, r := range text {
		if len(sentence) == 0 ||
			(unicode.IsLetter(prev) && !unicode.IsLetter(r)) ||
			(!unicode.IsLetter(prev) && unicode.IsLetter(r)) {
			// Start a new word and append to the sentence
			//w := []rune{r}
			w := string(r)
			sentence = append(sentence, w)
			prev = r
		} else {
			// Append the character to current word
			last := len(sentence) - 1
			//sentence[last] = append(sentence[last], r)
			sentence[last] += string(r)
			prev = r
		}
	}

	/*
		for _, w := range sentence {
			fmt.Print("[")
			for _, r := range w {
				fmt.Printf("%c", r)
			}
			fmt.Println("]")
		}
	*/

	// return if empty text
	if len(sentence) == 0 {
		return text, count, &dict
	}

	// Point startIndex to the first valid word
	startIndex := 0
	if !isLetter(sentence[startIndex]) {
		startIndex++
	}
	// Point endIndex to last word
	endIndex := len(sentence) - 1

	level := 0
	x := 0
	y := -1
	err := NotFound

	// Map for detecting duplicate keys and calculate counts
	dupMap := make(map[string]int)

	// Apply the tag recurssively
	for currIndex := startIndex; currIndex <= endIndex; {

		_, x, y, err = insertDefinition_r(ctx, sentence, currIndex,
			endIndex, level, x, y, Dictionary[lang])
		// fmt.Printf("curr=[%d], x=[%d], y=[%d]\n", currIndex, x, y)

		if err == nil && x <= y {
			key := ""
			for i := x; i <= y; i++ {
				key += sentence[i]
			}

			if v, ok := dupMap[strings.ToLower(key)]; duplicate || !ok {
				// Duplicates Allowed or First request
				// fmt.Printf("key=[%s]\n", key)

				term, val, err := Dictionary[lang].TermValue(key)
				if err == nil {

					// Insert the tags / definition
					// fmt.Printf("x=[%d], y=[%d], key=[%s]\n", x, y, key)
					var startTag, endTag string

					switch svc {
					case textSvc:
						startTag, endTag = insertTextTags(tag, key, val, css, valFlag)
					case taptargetSvc:
						startTag, endTag = insertTaptargetTags(tag, key, val, css, valFlag)
					case tooltipSvc:
						startTag, endTag = insertTooltipTags(tag, key, val, css, valFlag)
					}

					sentence[x] = startTag + sentence[x]
					sentence[y] = sentence[y] + endTag

					// Add term to the mini dictionary if not already added
					if !ok {
						dict.Definitions = append(dict.Definitions, api.Definition{Term: term, Value: val})
					}

					count++
					dupMap[strings.ToLower(key)] = v + 1
				}
			}

			// Skip the non-letter word
			currIndex = y + 2
		} else {
			// Jump the non-letter word
			currIndex += 2
		}
	}

	// Print Duplicate Map
	/*
		for k, v := range dupMap {
			fmt.Printf("[%s] : %d\n", k, v)
		}
	*/

	result := strings.Join(sentence, "")
	return result, count, &dict
}

// This function returns the posx and posy of the longest key in the
// original sentence and it points the current to the next word
func insertDefinition_r(ctx context.Context, sentence []string, currIndex, endIndex,
	level, x, y int, dict *trie.Trie) (int, int, int, error) {

	indent := "    "
	for i := 0; i < level; i++ {
		indent += indent
	}
	level++
	// fmt.Printf("%s============ LEVEL %d: currIndex: %d , x: %d, y: %d =============\n", indent, level, currIndex, x, y)
	// defer fmt.Printf("%s============ LEVEL %d: =============\n", indent, level)
	if currIndex > endIndex {
		return currIndex, x, y, nil
	}

	key := sentence[currIndex]

	// fmt.Printf("%skey=[%s]\n", indent, key)
	val, err := dict.Value(key)
	if err != nil {
		return currIndex, x, y, err
	}

	if level == 1 {
		x = currIndex
	}

	if len(val) > 0 {
		y = currIndex
	}

	// fmt.Printf("%sinsertDefinition_r: key=[%s], Val=[%s]\n", indent, key, val)
	// fmt.Printf("%sinsertDefinition_r: x=[%d], y=[%d]\n", indent, x, y)

	currIndex += 2 // Skip the non-letter word
	childDict, err := dict.Child(key)
	currIndex1, x1, y1, err := insertDefinition_r(ctx, sentence, currIndex, endIndex,
		level, x, y, childDict)
	if err != nil {
		return currIndex, x, y, nil
	}
	return currIndex1, x1, y1, nil
}
