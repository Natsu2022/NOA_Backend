package user

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

// SendOTPEmail sends an OTP to the user's email
func SendOTPEmail(email, otp string) error {
	log.Println("Sending OTP to email...")

	from := "toonglar@gmail.com"      // Replace with your email address
	password := "gbgs bzkr kdbg xvqe" // Replace with your app-specific password
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	log.Println("Sending email to:", email)

	// Dynamic content for the email
	emailData := struct {
		Name    string
		Message string
		OTP     string
	}{
		Name:    "John Doe",
		Message: "This is a dynamic message generated by Go!",
		OTP:     otp,
	}

	// Create dynamic HTML email content
	emailTemplate := `
	<html>
		<head></head>
		<body>
			<h1>Hello, Welcome to verify your Email !</h1>
			<p>{{.Message}}</p>
			<p>Your OTP is: <strong>{{.OTP}}</strong></p>
		</body>
	</html>
	`

	// Parse the template and generate HTML
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Println("Error parsing email template:", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, emailData); err != nil {
		log.Println("Error executing email template:", err)
		return err
	}

	// Set up email subject and content
	to := []string{email}
	msg := []byte("Subject: OTP Verification\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body.String())

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	} else {
		log.Println("Sent OTP to email successfully.")
		return nil
	}
}
