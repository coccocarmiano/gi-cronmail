package main

import (
	"cronmail/filters"
	"cronmail/mail"
	"cronmail/models"
	"cronmail/utils"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		ForceColors:      true,
		QuoteEmptyFields: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Infoln("loading .env")
	err := godotenv.Load()
	if err != nil {
		logrus.Warnln("Could not load .env files -- Proceding")
	}

	logrus.Infoln("loading config")
	config, err := models.GetConfig()
	if err != nil {
		log.Fatalf("could not build config: %s", err)
	}

	logrus.Infoln("getting gmail client")
	client, err := mail.GetGmailClient(config)
	if err != nil {
		logrus.Fatalf("could not get gmail client: %s", err)
	}

	logrus.Infoln("creating gsheets service")
	svc, err := utils.GetSheetsService(config)
	if err != nil {
		logrus.Fatalf("could not get sheets services: %s", err)
	}

	logrus.Infoln("getting data")
	filterer := filters.GetDefaultFilterer()
	people, err := utils.GetData(svc, config)
	if err != nil {
		logrus.Fatalf("could not get data: %s", err)
	}

	logrus.Infoln("Filtering out data")
	people = filterer.All(people)
	logrus.Debugln("got data:")
	for _, p := range people {
		logrus.Debugln(p)
	}

	logrus.Infoln("getting formatter")
	outgoing := mail.Mail{
		Body:    mail.DefaultMail,
		Title:   "",
		ReplyTo: config.GmailSender,
		SendAs:  config.SendAs,
	}
	formatter, err := mail.GetFormatter(outgoing)
	if err != nil {
		logrus.Fatalf("could not get formatter: %s", err)
	}

	params := map[string]string{
		"Nome": "",
		"Age":  "",
	}
	ok_count, err_count := 0, 0
	for _, pp := range people {
		utils.Format(&pp)
		age, err := pp.Age()
		if err != nil {
			err_count++
			continue
		}

		outgoing.Title = fmt.Sprintf("Tanti auguri, %s!", pp.FirstName)
		params["Nome"] = pp.FirstName
		params["Age"] = fmt.Sprint(age)
		body, err := formatter.String(pp, params)
		if err != nil {
			err_count++
			continue
		}

		logrus.Debug("Sending mail to ", pp.Mail)
		err = client.Send(pp.Mail, body)
		if err != nil {
			err_count++
			continue
		}

		ok_count++
		_, _ = body, client
	}

	logrus.Infof("Run Completed.\tOK: %d\tNOT OK: %d", ok_count, err_count)
}
