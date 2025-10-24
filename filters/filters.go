package filters

import (
	   "cronmail/models")

type Filter interface {
	Check(models.Person) bool
}

type Filterer interface {
	All([]models.Person) []models.Person
	Apply([]models.Person, Filter) []models.Person
}

type DefaultFilterer struct {
	filters []Filter
}

func (d DefaultFilterer) All(persons []models.Person) []models.Person {
	for _, f := range d.filters {
		persons = d.Apply(persons, f)
	}

	return persons
}

func (d DefaultFilterer) Apply(persons []models.Person, f Filter) []models.Person {
	filtered := make([]models.Person, 0)

	for _, p := range persons {
		if f.Check(p) {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func GetDefaultFilterer() DefaultFilterer {
	return DefaultFilterer{
		filters: []Filter{
			GetTodayBirthdayFilter(),
			GetEmailFilter(),
		},
	}
}
