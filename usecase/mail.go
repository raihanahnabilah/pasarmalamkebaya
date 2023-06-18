package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"pasarmalamkebaya/dto"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailUsecase interface {
	SendEmailVerification(input dto.RegisterEmailVerification)
}

type mailUsecase struct {
}

func NewMailUsecase() MailUsecase {
	return &mailUsecase{}
}

func (u *mailUsecase) SendEmailVerification(input dto.RegisterEmailVerification) {
	// From who
	from := mail.NewEmail(os.Getenv("EMAIL_SENDER"), os.Getenv("EMAIL_SENDER"))

	fmt.Println(input.Email)
	fmt.Println(input.Subject)
	fmt.Println(input.VerificationCode)

	// Subject
	subject := input.Subject

	// Send to Email:
	to := mail.NewEmail(input.Email, input.Email)

	// Content
	htmlContent, err := ParseTemplate("/Users/hanafatinah/Documents/pasarmalamkebaya/templates/verification.html", input)
	fmt.Println(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(htmlContent)
	// plainTextContent := "and easy to do anywhere, even with Go"
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	// Send email
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		fmt.Println(err)
	} else if response.StatusCode != 200 {
		fmt.Println(response)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
		fmt.Printf("Email is sent to %s successfully!", input.Email)
	}
}

func ParseTemplate(templateFile string, emailData dto.RegisterEmailVerification) (string, error) {

	// Parsed the template
	parsedTemplate, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Error here: ", err)
	}

	// Buffer to store the template
	buf := new(bytes.Buffer)
	err = parsedTemplate.Execute(buf, emailData)
	if err != nil {
		log.Println("Error executing template: ", err)
	}

	// Return the HTML content
	return buf.String(), nil
}
