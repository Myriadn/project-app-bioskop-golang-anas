package utils

import (
	"fmt"
	"net/smtp"

	"go.uber.org/zap"
)

type EmailService struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	Logger   *zap.Logger
}

func NewEmailService(host string, port int, username, password, from string, logger *zap.Logger) *EmailService {
	return &EmailService{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
		Logger:   logger,
	}
}

// SendOTPEmail mengirim email dengan OTP code
func (e *EmailService) SendOTPEmail(to, username, otpCode string) error {
	subject := "Verify Your Cinema Booking Account"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Segoe UI', Arial, sans-serif; line-height: 1.6; color: #233D4D; background-color: #F5FBE6; margin: 0; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 12px; overflow: hidden; border: 1px solid #e1e1e1; }
        .header { background-color: #215E61; color: #F5FBE6; padding: 30px; text-align: center; }
        .content { padding: 40px; }
        .otp-box { background-color: #F5FBE6; border: 2px solid #215E61; padding: 25px; text-align: center; margin: 30px 0; border-radius: 12px; }
        .otp-code { font-size: 36px; font-weight: bold; color: #FE7F2D; letter-spacing: 10px; margin: 10px 0; }
        .footer { text-align: center; padding: 20px; color: #233D4D; font-size: 12px; opacity: 0.7; }
        ul { padding-left: 20px; }
        li { margin-bottom: 8px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1 style="margin:0;">🎬 Cinema Booking</h1>
            <p style="margin:5px 0 0 0; opacity: 0.9;">Email Verification</p>
        </div>
        <div class="content">
            <h2 style="color: #215E61;">Hello, %s! 👋</h2>
            <p>Thank you for joining us. Please use the following code to verify your account:</p>

            <div class="otp-box">
                <p style="margin: 0; font-weight: bold; color: #215E61;">Your OTP Code</p>
                <div class="otp-code">%s</div>
                <p style="margin: 5px 0 0 0; color: #233D4D; font-size: 13px;">Valid for 10 minutes</p>
            </div>

            <p><strong>Security Tips:</strong></p>
            <ul>
                <li>Do not share this code with anyone.</li>
                <li>This code will expire shortly.</li>
                <li>If you didn't request this, please ignore this email.</li>
            </ul>
        </div>
        <div class="footer">
            <p>&copy; 2026 Cinema Booking System. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
	`, username, otpCode)

	return e.sendEmail(to, subject, body)
}

// SendWelcomeEmail mengirim email welcome setelah verifikasi
func (e *EmailService) SendWelcomeEmail(to, username string) error {
	subject := "Welcome to Cinema Booking System! 🎉"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Segoe UI', Arial, sans-serif; line-height: 1.6; color: #233D4D; background-color: #F5FBE6; margin: 0; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 12px; overflow: hidden; border: 1px solid #e1e1e1; }
        .header { background-color: #215E61; color: #F5FBE6; padding: 30px; text-align: center; }
        .content { padding: 40px; }
        .cta-button { display: inline-block; padding: 14px 28px; background-color: #FE7F2D; color: #ffffff; text-decoration: none; border-radius: 8px; font-weight: bold; margin-top: 20px; }
        .footer { text-align: center; padding: 20px; color: #233D4D; font-size: 12px; opacity: 0.7; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1 style="margin:0;">🎬 Welcome Aboard!</h1>
        </div>
        <div class="content">
            <h2 style="color: #215E61;">Hi %s, You're All Set! 🎉</h2>
            <p>Your account has been successfully verified. Now you're ready to explore the best movies in town.</p>
            <p>What can you do now?</p>
            <ul style="padding-left: 20px;">
                <li>Check out the latest <strong>Movie Premieres</strong></li>
                <li>Book your <strong>Favorite Seats</strong></li>
                <li>Get exclusive <strong>Promo Offers</strong></li>
            </ul>
            <div style="text-align: center;">
                <a href="#" class="cta-button">Start Booking Now</a>
            </div>
        </div>
        <div class="footer">
            <p>&copy; 2026 Cinema Booking System. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
	`, username)

	return e.sendEmail(to, subject, body)
}

func (e *EmailService) sendEmail(to, subject, body string) error {
	// Setup email headers
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n%s\n%s",
		e.From, to, subject, mime, body))

	// SMTP server configuration
	addr := fmt.Sprintf("%s:%d", e.Host, e.Port)
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	// Send email
	err := smtp.SendMail(addr, auth, e.From, []string{to}, msg)
	if err != nil {
		e.Logger.Error("Failed to send email",
			zap.String("to", to),
			zap.String("subject", subject),
			zap.Error(err),
		)
		return fmt.Errorf("failed to send email: %w", err)
	}

	e.Logger.Info("Email sent successfully",
		zap.String("to", to),
		zap.String("subject", subject),
	)

	return nil
}

// SendEmailAsync mengirim email secara async menggunakan goroutine
func (e *EmailService) SendEmailAsync(to, username, otpCode string) {
	go func() {
		if err := e.SendOTPEmail(to, username, otpCode); err != nil {
			e.Logger.Error("Async email send failed", zap.Error(err))
		}
	}()
}

// SendWelcomeEmailAsync mengirim welcome email secara async
func (e *EmailService) SendWelcomeEmailAsync(to, username string) {
	go func() {
		if err := e.SendWelcomeEmail(to, username); err != nil {
			e.Logger.Error("Async welcome email send failed", zap.Error(err))
		}
	}()
}
