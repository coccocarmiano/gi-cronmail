package utils

import (
	"context"
	"cronmail/models"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetSheetsService(c models.Config) (*sheets.Service, error) {
	credentialsJson := []byte(c.ServiceAccount)
	sheets, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(credentialsJson))
	if err != nil {
		return nil, err
	}
	return sheets, nil
}

func GetData(svc *sheets.Service, cfg models.Config) ([]models.Person, error) {
	cols := "C3:M"
	rng := fmt.Sprintf("%s!$%s", cfg.SheetName, cols)
	res, err := svc.Spreadsheets.Values.Get(cfg.SheetID, rng).Do()

	if err != nil {
		return nil, err
	}

	persons := make([]models.Person, len(res.Values))
	cnt := 0
	for _, row := range res.Values {
		if len(row) < 11 {
			continue
		}

		firstName, _ := row[0].(string)
		lastName, _ := row[1].(string)
		birhtDate, _ := row[2].(string)
		mail, _ := row[9].(string)
		fc, _ := row[10].(string)

		p := models.Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthdate:  birhtDate,
			Mail:       mail,
			FiscalCode: fc,
		}
		persons[cnt] = p
		cnt++
	}

	persons = persons[:cnt]
	return persons, nil
}
