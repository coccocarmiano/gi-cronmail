package filters

import "cronmail/models"

/* Valid Email Check */

func GetEmailFilter() Filter {
	return EmailFilter{}
}

type EmailFilter struct{}

func (e EmailFilter) Check(p models.Person) bool {
	// TODO:
	return true
}
