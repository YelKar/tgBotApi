package tgBotApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YelKar/tgBotApi/utils"
	"github.com/YelKar/tgBotApi/utils/errors"
	"io"
	"log"
	"net/http"
)

func (bot *Bot) SendMessage(chatID int, msgText string, params ...interface{}) {
	msg := utils.SentMessage{ChatID: chatID, Text: msgText}

	for _, param := range params {
		switch mu := param.(type) {
		case utils.Keyboard:
			msg.ReplyMarkup.Keyboard = &mu
		case utils.InlineKeyboard:
			msg.ReplyMarkup.InlineKeyboard = &mu
		case utils.RemoveKeyboard:
			msg.ReplyMarkup.RemoveKeyboard = &mu
		}
	}

	Json, _ := json.Marshal(msg)
	param := bytes.NewReader(Json)
	resp, err := http.Post(
		fmt.Sprintf(url, bot.TOKEN, "sendMessage"),
		"application/json",
		param,
	)
	if err != nil {
		var body []byte
		_, err := resp.Body.Read(body)
		if err != nil {
			panic(err)
		}
		panic(errors.Error{
			Code:  -1,
			Text:  err.Error() + string(body),
			Level: errors.MIDDLE,
		})
	}
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
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	updateJson, _ := io.ReadAll(res.Body)
	var resp utils.TGResponse
	err := json.Unmarshal(updateJson, &resp)
	if err != nil {
		log.Println(err.Error())
	}
	resp.JSON = string(updateJson)
	return resp
}
