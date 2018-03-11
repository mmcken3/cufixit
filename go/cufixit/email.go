package cufixit

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"

	"github.com/pkg/errors"
)

// Email is a struct used for sending go emails.
type Email struct {
	UserName    string
	Password    string
	Server      string
	Port        string
	SendTo      []string
	FromAddress string
	Feedback
}

// SendEmail sends an email to email address e.
func (se *Email) SendEmail() error {
	var b bytes.Buffer

	b.Write([]byte("To: "))

	for _, email := range se.SendTo {
		b.Write([]byte(email + ", "))
	}

	b.Write([]byte("\r\nSubject: CU Fix It Request "))
	b.Write([]byte(time.Now().Format("Jan-01-06 03:04 PM") + "\r\n"))
	b.Write([]byte("\r\nType: "))
	b.Write([]byte((se.Type.Type) + "\n"))
	b.Write([]byte("Building: "))
	b.Write([]byte((se.Building.Name) + "\n"))
	b.Write([]byte("Description: "))
	b.Write([]byte((se.Description) + "\n"))
	b.Write([]byte("Reported By: "))
	b.Write([]byte((se.UserName) + "\n"))
	b.Write([]byte("Contact: "))
	b.Write([]byte((se.PhoneNumber) + "\n"))
	b.Write([]byte("ImageURL: "))
	b.Write([]byte((se.ImageURL)))
	b.Write([]byte("\n\n"))
	fmt.Println(b.String())

	//msg := []byte("\r\nSubject: CU Fix It Request\r\nMessage Content Here")

	// Set up authentication information
	auth := smtp.PlainAuth("", se.UserName, se.Password, se.Server)

	msg := b.Bytes()
	err := smtp.SendMail(se.Server+":"+se.Port, auth, se.FromAddress, se.SendTo, msg)
	if err != nil {
		return errors.Wrapf(err, "Failed when sending email.")
	}
	return nil
}
