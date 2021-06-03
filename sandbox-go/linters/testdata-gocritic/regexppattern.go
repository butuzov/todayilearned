package testdata_gocritic

import "regexp"

// Example for the rule: regexpPattern

func _() {
	re := regexp.MustCompile(`google.com|yandex.ru`)
	re.FindAllString("yahoo.com", -1)
}
