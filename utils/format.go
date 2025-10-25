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
	if len(s) < 1 {
		return s
	}

	bts := []rune(strings.ToLower(s))
	bts[0] = unicode.ToUpper(bts[0])
	for i := 1; i < len(bts); i++ {
		if !unicode.IsLetter(bts[i]) && bts[i] != ' ' {
			bts[i-1] = unicode.ToUpper(bts[i-1])
		}

		if unicode.IsSpace(bts[i-1]) {
			bts[i] = unicode.ToUpper(bts[i])
		}
	}

	return string(bts)
}

func Format(p *models.Person) *models.Person {
	firstName, lastName := p.FirstName, p.LastName
	firstName, lastName = Capitalize(firstName), Capitalize(lastName)
	p.FirstName = firstName
	p.LastName = lastName
	return p
}
