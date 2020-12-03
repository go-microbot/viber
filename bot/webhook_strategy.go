package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-microbot/viber/models"
)

const contentSignatureHeader = "X-Viber-Content-Signature"

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
	ServeURL        string
	VerifySignature bool
	SignatureKey    string
}

type webHookHandler struct {
	events          chan models.CallbackEvent
	errs            chan error
	verifySignature bool
	key             string
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
		events:          wh.eventsChan,
		errs:            wh.errorsChan,
		verifySignature: wh.cfg.VerifySignature,
		key:             wh.cfg.SignatureKey,
	}); err != nil {
		panic(err)
	}
}

func (h *webHookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.errs <- fmt.Errorf("could not read callback body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if h.verifySignature {
		signature := r.Header.Get(contentSignatureHeader)
		if err := validateMAC(body, signature, h.key); err != nil {
			h.errs <- fmt.Errorf("could not verify callback signature: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	var event models.CallbackEvent
	if err := json.Unmarshal(body, &event); err != nil {
		h.errs <- fmt.Errorf("could not parse callback body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h.events <- event

	w.WriteHeader(http.StatusOK)
}

func validateMAC(message []byte, signature, key string) error {
	decodedSignature, err := hex.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("could not decode signature: %v", err)
	}

	mac := hmac.New(sha256.New, []byte(key))
	_, err = mac.Write(message)
	if err != nil {
		return fmt.Errorf("could not write mac message with provided key: %v", err)
	}

	expectedMAC := mac.Sum(nil)
	if !hmac.Equal(decodedSignature, expectedMAC) {
		return fmt.Errorf("mac is invalid")
	}

	return nil
}
