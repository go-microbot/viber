package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

const (
	botNameCtxKey         = "bot_name"
	webhookURLCtxKey      = "webhook_url"
	chatMemberIDCtxKey    = "chat_member_id"
	pictureToSendCtxKey   = "pictute_url"
	videoToSendCtxKey     = "video_url"
	videoToSendSizeCtxKey = "video_size"
	fileToSendCtxKey      = "file_url"
	fileToSendSizeCtxKey  = "file_size"
	stickerIDCtxKey       = "sticker_id"
	defaultWebhookPort    = "8443"
)

var (
	testAPI     ViberAPI
	testContext context.Context
)

type Testable interface {
	Test(ctx context.Context, t *testing.T) context.Context
}

// TestDataKey represents test context data key.
type TestDataKey interface{}

type ngrokConfig struct {
	Tunnels []struct {
		PublicURL string `json:"public_url"`
	} `json:"tunnels"`
}

func TestMain(m *testing.M) {
	// create test API instance.
	testAPI = NewViberAPI(os.Getenv("TEST_BOT_TOKEN"))
	rand.Seed(time.Now().Unix())

	// run ngrock server.
	publicURL, err := runNgrockServer(defaultWebhookPort)
	if err != nil {
		panic(err)
	}
	testContext = context.WithValue(context.Background(), TestDataKey(webhookURLCtxKey), publicURL)

	// run webhook local server.
	go func() {
		err := runLocalWebhookServer(defaultWebhookPort)
		if err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

func Test_NewViberAPI(t *testing.T) {
	tAPI := NewViberAPI("123")
	require.NotNil(t, tAPI)
	require.Equal(t, "123", tAPI.token)
	require.NotNil(t, tAPI.client)
	require.Equal(t, baseURL, tAPI.url)
}

func TestViberAPI_Integration(t *testing.T) {
	// prepare context data.
	wd, err := os.Getwd()
	require.NoError(t, err)
	cfgMap, err := parseTestConfig(path.Join(wd, "./test_data/config.json"))
	require.NoError(t, err)
	ctx := testContext
	for k, v := range cfgMap {
		if val, ok := v.(float64); ok {
			v = int64(val)
		}
		ctx = context.WithValue(ctx, TestDataKey(k), v)
	}

	testCases := []struct {
		name        string
		testHandler Testable
	}{
		{
			name:        "getAccountInfo",
			testHandler: getAccountInfo{},
		},
		{
			name:        "removeWebhook",
			testHandler: removeWebhook{},
		},
		{
			name:        "setWebhook",
			testHandler: setWebhook{},
		},
		{
			name:        "sendTextMessage",
			testHandler: sendTextMessage{},
		},
		{
			name:        "sendPictureMessage",
			testHandler: sendPictureMessage{},
		},
		{
			name:        "sendVideoMessage",
			testHandler: sendVideoMessage{},
		},
		{
			name:        "sendFileMessage",
			testHandler: sendFileMessage{},
		},
		{
			name:        "sendContactMessage",
			testHandler: sendContactMessage{},
		},
		{
			name:        "sendLocationMessage",
			testHandler: sendLocationMessage{},
		},
		{
			name:        "sendURLMessage",
			testHandler: sendURLMessage{},
		},
		{
			name:        "sendStickerMessage",
			testHandler: sendStickerMessage{},
		},
		{
			name:        "sendRichMediaMessage",
			testHandler: sendRichMediaMessage{},
		},
		{
			name:        "sendMessageWithKeyboard",
			testHandler: sendMessageWithKeyboard{},
		},
		{
			name:        "sendBroadcastMessage",
			testHandler: sendBroadcastMessage{},
		},
		{
			name:        "getUserDetails",
			testHandler: getUserDetails{},
		},
	}
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctx = tc.testHandler.Test(ctx, t)
			time.Sleep(2 * time.Second)
		})
	}
}

func parseTestConfig(cfgPath string) (map[string]interface{}, error) {
	var cfgMap map[string]interface{}

	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &cfgMap); err != nil {
		return nil, err
	}

	return cfgMap, nil
}

func runNgrockServer(port string) (string, error) {
	cmd := exec.Command("ngrok", "http", port)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("could not run ngrok: %v", err)
	}

	time.Sleep(5 * time.Second)

	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		return "", fmt.Errorf("could not get ngrok configuration: %v", err)
	}

	var cfg ngrokConfig
	if err := json.NewDecoder(resp.Body).Decode(&cfg); err != nil {
		return "", fmt.Errorf("could not parse ngrok configuration: %v", err)
	}

	for i := range cfg.Tunnels {
		if strings.HasPrefix(cfg.Tunnels[i].PublicURL, "https") {
			return cfg.Tunnels[i].PublicURL, nil
		}
	}

	return "", errors.New("could not find public https URL")
}

func runLocalWebhookServer(port string) error {
	baseHandler := func(w http.ResponseWriter, req *http.Request) {
		var event models.CallbackEvent
		if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
			log.Printf("could not decode request body: %v", err)
			return
		}
		switch event.Event {
		case models.EventTypeWebhook:
			w.WriteHeader(http.StatusOK)
		default:
			log.Printf("unknown event type: %s", event.Event)
		}
	}
	http.HandleFunc("/", baseHandler)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
	return err
}
