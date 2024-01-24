package utils

import (
	"github.com/go-mail/mail"
)

// Send Verification Email
func SendVerificationEmail(email string, token string, userID string) error {

	m := mail.NewMessage()
	m.SetHeader("From", "admin@admin.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify your email")

	m.SetBody("text/html", "Hello <b>Fis</b> </br> This is your verification link: <a href='"+SERVER_URL+"/verify/?id="+userID+"&token="+token+"'>Click here</a>")

	d := mail.NewDialer(EMAIL_HOST, 587, EMAIL_USERNAME, EMAIL_PASSWORD)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Send Reset Password Email
func SendResetPasswordEmail(email string, token string) error {
	m := mail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset your password")

	m.SetBody("text/html", "Hello <b>Fis</b> !")

	d := mail.NewDialer("smtp.mailtrap.io", 587, "b1d945008ecdb6", "5726370061bde7")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Send Order Confirmation Email
func SendOrderConfirmationEmail(email string, token string) error {
	m := mail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Order Confirmation")

	m.SetBody("text/html", "Hello <b>Fis</b> !")

	d := mail.NewDialer("smtp.mailtrap.io", 587, "b1d945008ecdb6", "5726370061bde7")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
