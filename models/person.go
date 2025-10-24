package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string
	LastName   string
	Birthdate  string
	Mail       string
	FiscalCode string
}

func (p *Person) Age() (int, error) {
	splits := strings.Split(p.Birthdate, "/")
	if len(splits) < 3 {
		return 0, fmt.Errorf("invalid date fmt")
	}

	yy := splits[2]
	intYear, err := strconv.Atoi(yy)

	if err != nil {
		return 0, err
	}

	return time.Now().Year() - intYear, nil
}
