package pkg

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"

	"app/config"
)

type EmailPayload struct {
	Recipients []string
	Body       string
	Subject    string
	BCC        []string
}

type response string
type trackId string

func SendEmail(payload EmailPayload) (response, trackId, error) {
	recipients := payload.Recipients
	body := payload.Body
	subject := payload.Subject
	bcc := payload.BCC

	mg := mailgun.NewMailgun(
		config.AppConfig.MAILGUN_DOMAIN,
		config.AppConfig.MAILGUN_PRIVATE_API_KEY,
	)
	mg.SetAPIBase(config.AppConfig.MAILGUN_API_BASE)

	message := mg.NewMessage(config.AppConfig.MAILGUN_SENDER, subject, "", recipients...)
	for _, item := range bcc {
		message.AddBCC(item)
	}
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return "", "", err
	}

	return response(resp), trackId(id), nil
}
