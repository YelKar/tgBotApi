package tgBotApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YelKar/tgBotApi/utils"
	"io"
	"net/http"
)

func (bot *Bot) SendMessage(chatID int, msgText string) {
	msg := utils.SentMessage{ChatID: chatID, Text: msgText}
	Json, _ := json.Marshal(msg)
	param := bytes.NewReader(Json)
	http.Post(
		fmt.Sprintf(url, bot.TOKEN, "sendMessage"),
		"application/json",
		param,
	)
}

func (bot *Bot) GetUpdates() utils.TGResponse {
	query := utils.Query{Offset: lastUpdate + 1}
	QJson, _ := json.Marshal(query)
	param := bytes.NewReader(QJson)
	res, _ := http.Post(
		fmt.Sprintf(url, bot.TOKEN, "getUpdates"),
		"application/json",
		param,
	)
	defer res.Body.Close()

	updateJson, _ := io.ReadAll(res.Body)
	var resp utils.TGResponse
	json.Unmarshal(updateJson, &resp)
	resp.JSON = string(updateJson)
	return resp
}
