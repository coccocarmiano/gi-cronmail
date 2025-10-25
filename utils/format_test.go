package utils_test

import (
	"cronmail/utils"
	"testing"
)

func TestCapitalize(t *testing.T) {
	cases := map[string]string{
		"":       "",
		"Mario":  "Mario",
		"mario":  "Mario",
		"MARIO":  "Mario",
		"Mar io": "Mar Io",
		"Mar IO": "Mar Io",
	}

	for k, v := range cases {
		if utils.Capitalize(k) != v {
			t.Error("wrong capitalization")
		}
	}
}
