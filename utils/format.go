package utils

import (
	"cronmail/models"
	"fmt"
	"strings"
	"unicode"
)

func FormatRange(sheet string, rng string) string {
	return fmt.Sprintf("%s!%s", sheet, rng)
}

func Capitalize(s string) string {
	s = strings.Trim(s, " ")
	ss := strings.Split(s, " ")
	capitalized := make([]string, len(ss), 0)
	for idx, word := range ss {
		if len(word) < 1 {
			continue
		}

		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		capitalized[idx] = string(runes)
	}

	return strings.Join(capitalized, " ")
}

func Format(p *models.Person) *models.Person {
	firstName, lastName := p.FirstName, p.LastName
	firstName, lastName = Capitalize(firstName), Capitalize(lastName)
	p.FirstName = firstName
	p.LastName = lastName
	return p
}
