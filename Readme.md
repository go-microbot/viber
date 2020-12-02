# [GO-MICROBOT] Viber

## This is fully documented and easy to use go-bindings for working with the Bot [Viber](https://www.viber.com/ru/) API.

[![Coverage Status](https://coveralls.io/repos/github/go-microbot/viber/badge.svg)](https://coveralls.io/github/go-microbot/viber)

<p align="center">
  <img height="120" src="./.github/assets/gopher.png">
  <img height="80" src="./.github/assets/heart.png">
  <img height="120" src="./.github/assets/viber.png">
</p>

> Please read the [Official Viber API Documentation](https://developers.viber.com/docs/api/rest-bot-api/) before starting.

# Guides

- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Create bot token](#bot-token)
- [Callbacks](#callbacks)
  - [Setup Webhook](#webhook)
- [Example](#example)
- [Test](#test)
  - [Local testing](#run-tests-locally)
  - [Lint](#run-linter)

## Getting started

### Installation
Download the latest version of the Bot API.

```bash
go get -u github.com/go-microbot/viber
```

### Bot token
Create your own bot token. Follow the [Official guide](https://partners.viber.com/account/create-bot-account).

## Callbacks
Each callback will contain a signature on the JSON passed to the callback. The signature is HMAC with SHA256 that will use the authentication token as the key and the JSON as the value. The result will be passed as HTTP Header `X-Viber-Content-Signature` so the receiver can determine the origin of the message.

### Webhook
For more information see [Setting a Webhook](https://developers.viber.com/docs/api/rest-bot-api/#setting-a-webhook) article. 

```go
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
	}))

	// setup Webhook.
	go func() {
		whResp, err := botAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{
			URL: "https://55442d01e546.ngrok.io", // use your website URL (SSL required).
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
```

> Viber **doesn’t support** self signed certificates.

## Example
See the [examples](./examples) folder to get all available examples.

## Test

### Run tests locally
To run tests locally please specify the `TEST_BOT_TOKEN` env variable. It should contains your bot token.

Use the following command:
```bash
go test -p 1
```

Or use [Makefile](./Makefile)'s `test` command:
```bash
make test
```

### Run linter
Use the following commands:
```bash
golangci-lint cache clean
golangci-lint run --config .golangci.yml --timeout=5m
```

Or use [Makefile](./Makefile)'s `lint` command:
```bash
make lint
```
