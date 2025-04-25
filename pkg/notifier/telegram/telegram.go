package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-monitoring/config"
	"net/http"
)

type TelegramNotifier struct {
	Config *configs.TelegramConfig
}

func NewTelegramNotifier(config *configs.TelegramConfig) *TelegramNotifier {
	return &TelegramNotifier{
		Config: config,
	}
}

func (service *TelegramNotifier) SendAlert(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", service.Config.BOT_TOKEN)
	body, _ := json.Marshal(map[string]string{
		"chat_id": service.Config.CHAT_ID,
		"text":    message,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
