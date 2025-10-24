package models

import (
	"fmt"
	"os"
)

type Config struct {
	ServiceAccount string
	SheetID        string
	SheetName      string
	SendAs         string

	GmailSender      string
	GmailAppPassword string
}

func GetConfig() (Config, error) {
	serviceAccount := os.Getenv("SERVICE_ACCOUNT")
	sheetId := os.Getenv("SHEET_ID")
	sheetName := os.Getenv("SHEET_NAME")
	sendAs := os.Getenv("SEND_AS")

	if serviceAccount == "" {
		return Config{}, fmt.Errorf("SERVICE_ACCOUNT undefined")
	}

	if sheetId == "" {
		return Config{}, fmt.Errorf("SHEET_ID undefined")
	}

	if sheetName == "" {
		return Config{}, fmt.Errorf("SHEET_NAME undefined")
	}

	if sendAs == "" {
		return Config{}, fmt.Errorf("SEND_AS undefined")
	}

	gmailSender := os.Getenv("GMAIL_SENDER")
	gmailAppPassword := os.Getenv("GMAIL_APP_PASSWORD")

	return Config{
		ServiceAccount:   serviceAccount,
		SheetID:          sheetId,
		SheetName:        sheetName,
		SendAs:           sendAs,
		GmailSender:      gmailSender,
		GmailAppPassword: gmailAppPassword,
	}, nil
}
