package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Payload struct {
	Data []EmailData `json:"data"`
}

type EmailData struct {
	To  string `json:"to"`
	Url string `json:"url"`
}

func main() {
	var argData string
	flag.StringVar(&argData, "d", "{}", "{}")
	flag.Parse()

	var allData Payload

	fmt.Println(argData)
	err := json.Unmarshal([]byte(argData), &allData)
	if err != nil {
		fmt.Println("Please provide valid json!!")
		fmt.Println(err)
	}

	client := sendgrid.NewSendClient("testkey")
	fromAddress := "test@example.com"
	fromName := "testName"

	m := mail.NewV3Mail()
	e := mail.NewEmail(fromName, fromAddress)
	m.SetFrom(e)
	m.SetTemplateID("d-17dbb63a630348068b4b8d875ac50621")

	for i := 0; i < len(allData.Data); i++ {
		fmt.Println(allData.Data[i].To)
		p := mail.NewPersonalization()
		to := mail.NewEmail("", allData.Data[i].To)

		p.SetDynamicTemplateData("campaign_url", allData.Data[i].Url)
		p.AddTos(to)
		m.AddPersonalizations(p)
	}
	fmt.Println("Sending mail...")

	response, err := client.Send(m)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
