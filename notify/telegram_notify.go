package notify

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type TelegramNotify struct {
	BotToken          string `json:"token"`
	ChatId            string `json:"chatId"` //Not mandatory field
}

func (telegramNotify TelegramNotify) GetClientName() string {
	return "Telegram"
}

func (telegramNotify TelegramNotify) Initialize() error {

	if len(strings.TrimSpace(telegramNotify.BotToken)) == 0 {
		return errors.New("Telegram: token is a required field")
	}

	if len(strings.TrimSpace(telegramNotify.ChatId)) == 0 {
		return errors.New("Telegram: chatId is a required field")
	}

	return nil
}

func (telegramNotify TelegramNotify) SendResponseTimeNotification(responseTimeNotification ResponseTimeNotification) error {

	message := getMessageFromResponseTimeNotification(responseTimeNotification)
	return telegramNotify.sendChat(message)

}

func (telegramNotify TelegramNotify) SendErrorNotification(errorNotification ErrorNotification) error {

	message := getMessageFromErrorNotification(errorNotification)
	return telegramNotify.sendChat(message)
}

func (telegramNotify TelegramNotify) sendChat(message string) error {
	req, err := http.NewRequest("GET", "https://api.telegram.org/" + telegramNotify.BotToken + "/sendMessage", nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("chat_id", telegramNotify.ChatId)
	q.Add("text", message)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Telegram : Send notifaction failed. Response code " + strconv.Itoa(resp.StatusCode))
	}

	return nil

}
