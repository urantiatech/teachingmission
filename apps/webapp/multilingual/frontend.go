package multilingual

const defaultLanguage = "en"

var frontend map[string]map[string]string

func init() {
	frontend = map[string]map[string]string{
		"en": {
			"title": "Teaching Mission",
		},
		"fr": {
			"title": "French Teaching Mission",
		},
		"es": {
			"title": "Spanish Teaching Mission",
		},
		"ru": {
			"title": "Russian Teaching Mission",
		},
		"hi": {
			"title": "Hindi Teaching Mission",
		},
	}
}

func Get(lang, key string) string {
	if el, ok := frontend[lang]; ok {
		return el[key]
	}
	return Get(defaultLanguage, key)
}
