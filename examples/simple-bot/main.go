package main

import (
	"context"
	"fmt"

	"github.com/go-microbot/viber/api"
	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/bot"
	"github.com/go-microbot/viber/models"
)

const token = "<PASTE_YOUR_TOKEN_HERE>"

func main() {
	// init Bot API with token.
	botAPI := api.NewViberAPI(token)

	// create Bot instance.
	myBot := bot.NewViberBot(&botAPI)

	// start listening.
	go myBot.WaitForUpdates(bot.NewWebhookStrategy(bot.WebhookConfig{
		ServeURL: "localhost:8443", // server to catch Viber requests.
		// if you want to validate each callback signature add these parameters as well.
		// More info: https://developers.viber.com/docs/api/rest-bot-api/#callbacks.
		VerifySignature: true,
		SignatureKey:    token,
	}))

	// setup Webhook.
	go func() {
		whResp, err := botAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{
			URL: "https://03322284e668.ngrok.io", // use your website URL (SSL required).
		})
		if err != nil {
			panic(err)
		}
		if whResp.Status != models.ResponseStatusCodeOK {
			panic(fmt.Sprintf("request to set webhook returned unexpected status: %d - %s", whResp.Status, whResp.StatusMessage))
		}
	}()

	// listen Bot's events.
	events, errs := myBot.Callbacks()
	for {
		select {
		case event, ok := <-events:
			if !ok {
				fmt.Println("events channel closed")
				return
			}

			switch event.Event {
			case models.EventTypeWebhook:
				fmt.Println("webhook successfully installed")
			case models.EventTypeMessage:
				// send "hello" message.
				_, err := myBot.API().SendTextMessage(context.Background(), apiModels.SendTextMessageRequest{
					GeneralMessageRequest: apiModels.GeneralMessageRequest{
						Receiver: event.Sender.ID,
						Type:     models.MessageTypeText,
						Sender: apiModels.MessageSender{
							Name: "Greeting bot",
						},
					},
					Text: fmt.Sprintf("Hello, %s!", event.Sender.Name),
				})
				if err != nil {
					fmt.Printf("could not send message to user: %v", err)
				}
			}
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
