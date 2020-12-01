package bot

import (
	"encoding/json"
	"net/http"

	"github.com/go-microbot/viber/models"
)

// StrategyWebhook represents object to receive icoming events
// using webhook (https://developers.viber.com/docs/api/rest-bot-api/#setting-a-webhook).
type StrategyWebhook struct {
	eventsChan chan models.CallbackEvent
	errorsChan chan error
	doneChan   chan struct{}
	cfg        WebhookConfig
}

// WebhookConfig represents webhook configuration.
type WebhookConfig struct {
	ServeURL string
}

type webHookHandler struct {
	events chan models.CallbackEvent
	errs   chan error
}

// NewWebhookStrategy returns new instance of the StrategyWebhook.
func NewWebhookStrategy(cfg WebhookConfig) StrategyWebhook {
	return StrategyWebhook{
		cfg:        cfg,
		eventsChan: make(chan models.CallbackEvent, defaultUpdatesLimit),
		doneChan:   make(chan struct{}),
		errorsChan: make(chan error),
	}
}

// Callbacks returns events channel.
func (wh StrategyWebhook) Callbacks() chan models.CallbackEvent {
	return wh.eventsChan
}

// Errors returns errors channel.
func (wh StrategyWebhook) Errors() chan error {
	return wh.errorsChan
}

// Listen starts webhook listening.
func (wh StrategyWebhook) Listen() {
	go wh.listen()

	<-wh.doneChan
}

// Stop stops webhook listening.
func (wh StrategyWebhook) Stop() {
	wh.doneChan <- struct{}{}
	close(wh.errorsChan)
	close(wh.eventsChan)
}

func (wh StrategyWebhook) listen() {
	if err := http.ListenAndServe(wh.cfg.ServeURL, &webHookHandler{
		events: wh.eventsChan,
		errs:   wh.errorsChan,
	}); err != nil {
		panic(err)
	}
}

func (h *webHookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var event models.CallbackEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.errs <- err
		return
	}

	h.events <- event

	w.WriteHeader(http.StatusOK)
}
