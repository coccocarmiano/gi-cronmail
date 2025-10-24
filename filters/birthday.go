package filters

import (
	"cronmail/models"
	"time"
)

func GetTodayBirthdayFilter() Filter {
	now := time.Now()
	day := now.Day()
	month := int(now.Month())

	return BirthdayFilter{
		currentDay:   day,
		currentMonth: month,
	}
}

type BirthdayFilter struct {
	currentDay   int
	currentMonth int
}

func (b BirthdayFilter) Check(p models.Person) bool {
	bd, err := time.Parse("02/01/2006", p.Birthdate)
	if err != nil {
		return false
	}

	if bd.Day() == b.currentDay && int(bd.Month()) == b.currentMonth {
		return true
	}

	return false
}
