package views

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"text/template"

	"github.com/urantiatech/teachingmission/apigateways/webgw/api"
	"github.com/urantiatech/teachingmission/apps/webapp/constants"
)

var funcMap = template.FuncMap{
	"ad":                   advertisement,
	"sectionAd":            sectionAd,
	"sum":                  sum,
	"mod":                  mod,
	"mul":                  mul,
	"div":                  div,
	"firstRecord":          firstRecord,
	"lastRecord":           lastRecord,
	"randomColor":          randomColor,
	"subdomain":            subdomain,
	"langname":             langname,
	"flag":                 flag,
	"title":                strings.Title,
	"rowStart":             rowStart,
	"vue":                  vue,
	"pagerLeftLink":        pagerLeftLink,
	"pagerRightLink":       pagerRightLink,
	"pagerIsActive":        pagerIsActive,
	"pagerAllLabels":       pagerAllLabels,
	"markedAs":             markedAs,
	"activeUri":            activeUri,
	"isHome":               isHome,
	"isNotHome":            isNotHome,
	"isSearch":             isSearch,
	"isNotSearch":          isNotSearch,
	"slug":                 slug,
	"isFilterAvailable":    isFilterAvailable,
	"isDateRangeAvailable": isDateRangeAvailable,
}

type Query struct {
	String string // Full Query string
	Q      string // Search Keywords
	C      string // Category filter
	B      string // BeingType filter
	T      string // Teacher filter
	R      string // Receiver filter
	Y1     int    // Year start
	Y2     int    // Year end
}

type Pager struct {
	Num   int
	Label string
	Uri   string
	Hide  bool
}

func sum(a, b int) int {
	return a + b
}

func mod(a, b int) int {
	return a % b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func firstRecord(skipped, limit, total int) int {
	if limit <= 0 || total <= 0 {
		return 0
	}
	return skipped + 1
}

func lastRecord(skipped, limit, total int) int {
	if limit <= 0 {
		return 0
	}

	if total-skipped >= limit {
		return skipped + limit
	}
	return skipped + total%limit
}

func markedAs() int {
	// 3 => Favorite
	// 2 => Read
	// 1 => Seen
	// 0 => Unseen
	return rand.Intn(4)
}

func activeUri(uri1 string, uri2 string) bool {
	return uri1 == uri2
}

func sectionAd(image string, url string) string {
	card := fmt.Sprintf(
		"<div class=\"card card-gap\">"+
			"<a href=\"%s\">"+
			"<div class=\"card-image\">"+
			"<img src=\"/images/sections/%s\">"+
			"</div>"+
			"</a>"+
			"</div>",
		url, image)
	return card
}

func advertisement(image string, url string) string {
	card := fmt.Sprintf(
		"<div class=\"row\">"+
			"<div class=\"col s12\">"+

			"<div class=\"card\">"+
			"<a href=\"%s\" target=\"_blank\">"+
			"<div class=\"card-image\">"+
			"<img src=\"/images/ads/%s\">"+
			"</div>"+
			"</a>"+
			"</div>"+

			"</div>"+
			"</div>",
		url, image)
	return card
}

func randomColor() string {
	i := rand.Intn(len(constants.Palette))
	return constants.Palette[i]
}

func subdomain(langcode string) string {
	if l, ok := lmap[langcode]; ok {
		return l.Subdomain
	}
	return ""
}

func langname(langcode string) string {
	if l, ok := lmap[langcode]; ok {
		return l.Name
	}
	return ""
}

func flag(langcode string) string {
	if l, ok := lmap[langcode]; ok {
		return l.Flag
	}
	return ""
}

func rowStart(x int) int {
	if x%4 == 0 {
		return 1
	} else {
		return 0
	}
}

func makeQueryString(q string, filter api.Filter) string {
	uri := "/?q="

	if q != "" {
		uri += q
	}
	if filter.Category != "" {
		uri += "&c=" + filter.Category
	}
	if filter.Being != "" {
		uri += "&b=" + filter.Being
	}
	if filter.Teacher != "" {
		uri += "&t=" + filter.Teacher
	}
	if filter.Receiver != "" {
		uri += "&r=" + filter.Receiver
	}
	if filter.StartDate != "" {
		uri += "&y1=" + filter.StartDate
	}
	if filter.EndDate != "" {
		uri += "&y2=" + filter.EndDate
	}
	return uri
}

func pagerLeftLink(q string, filter api.Filter, limit, skipped int) string {
	// fmt.Printf("pagerLeftLink: q=[%s], limit=[%d], skipped=[%d]\n", q, limit, skipped)
	p := skipped / limit
	uri := makeQueryString(q, filter)

	if limit != constants.SearchResultsLimit {
		return uri + "&limit=" + strconv.Itoa(limit) + "&p=" + strconv.Itoa(p)
	}
	return uri + "&p=" + strconv.Itoa(p)
}

func pagerRightLink(q string, filter api.Filter, limit, skipped, total int) string {
	uri := makeQueryString(q, filter)

	if skipped+limit >= total {
		return ""
	}
	p := skipped/limit + 2

	if limit != constants.SearchResultsLimit {
		return uri + "&limit=" + strconv.Itoa(limit) + "&p=" + strconv.Itoa(p)
	}
	return uri + "&p=" + strconv.Itoa(p)
}

func pagerIsActive(limit, skipped, p int) bool {
	// fmt.Printf("pagerIsActive: limit=[%d], skipped=[%d], p=[%d]\n", limit, skipped, p)
	return p == (skipped/limit + 1)
}

func pagerAllLabels(q string, filter api.Filter, limit, skipped, total int) []Pager {
	uri := makeQueryString(q, filter)

	currentPage := skipped / limit
	totalPages := total / limit
	if total%limit > 0 {
		totalPages++
	}

	startingPager := currentPage - 4
	if startingPager <= 0 {
		startingPager = 1
	}

	activePager := skipped/limit + 1

	var p []Pager
	for i := 0; i < constants.SearchPagerCount && i+startingPager <= totalPages; i++ {
		pageNum := i + startingPager
		pageLabel := strconv.Itoa(pageNum)
		var pageUri string

		if limit != constants.SearchResultsLimit {
			pageUri = uri + "&limit=" + strconv.Itoa(limit) + "&p=" + pageLabel
		} else {
			pageUri = uri + "&p=" + pageLabel
		}

		// Hide for mobile devices
		pagerGap := pageNum - activePager
		if pagerGap < 0 {
			pagerGap = -pagerGap
		}

		hideOnMobile := false
		if pagerGap > 2 {
			hideOnMobile = true
		}

		pager := Pager{pageNum, pageLabel, pageUri, hideOnMobile}
		p = append(p, pager)

		// Add ... if more pages
		if i == 9 && totalPages > i+startingPager {
			pager := Pager{-1, "...", "#!", false}
			p = append(p, pager)
		}
	}
	return p
}

func vue(s string) string {
	return "{{ " + s + " }}"
}

func isHome(p Page) bool {
	return p.URI == "/" && p.Search.Query == ""
}

func isNotHome(p Page) bool {
	return !isHome(p)
}

func isSearch(p Page) bool {
	return p.URI == "/"
}

func isNotSearch(p Page) bool {
	return !isSearch(p)
}

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

func isFilterAvailable(f api.Filter) bool {
	return f.Category != "" ||
		f.Group != "" ||
		f.Being != "" ||
		f.Teacher != "" ||
		f.Receiver != ""
}

func isDateRangeAvailable(y1, y2 string) bool {
	return y1 != "" || y2 != ""
}
