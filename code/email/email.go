package email

import (
	"fmt"
	"io/ioutil"
	"net/mail"
	"strings"
)

type Email struct {
	MessageId string   `json:"messageID"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	FromName  string   `json:"fromName"`
	Cc        []string `json:"cc"`
	Bcc       []string `json:"bcc"`
	Body      string   `json:"body"`
}

func NewEmail(header mail.Header, body []byte) Email {
	email := Email{
		MessageId: header.Get("Message-ID"),
		Date:      header.Get("Date"),
		From:      header.Get("From"),
		To:        strings.Split(header.Get("To"), ","),
		Subject:   header.Get("Subject"),
		FromName:  header.Get("X-From"),
		Cc:        strings.Split(header.Get("X-cc"), ","),
		Bcc:       strings.Split(header.Get("X-bcc"), ","),
		Body:      string(body),
	}

	return email
}

func ParseContent(fileContent string) (Email, error) {
	reader := strings.NewReader(fileContent)
	message, err := mail.ReadMessage(reader)

	if message == nil {
		fmt.Print(err)
		return Email{}, err
	}

	header := message.Header
	body, err := ioutil.ReadAll(message.Body)

	if err != nil {
		fmt.Print(err)
		return Email{}, err
	}

	return NewEmail(header, body), nil
}
