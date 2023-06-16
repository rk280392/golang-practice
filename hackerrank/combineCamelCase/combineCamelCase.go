package combineCamelCase

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CombineCamelCase(words string) string {
	parts := strings.Split(words, " ")
	for i, _ := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(parts[i])
		} else {
			parts[i] = cases.Title(language.Und).String(parts[i])
		}
	}
	return strings.Join(parts, "")
}
