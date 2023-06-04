package tgBotApi

import (
	"github.com/YelKar/tgBotApi/utils"
	"github.com/YelKar/tgBotApi/utils/errors"
)

func (bot *Bot) Polling() <-chan struct{} {
	if len(bot.MessageHandlers) == 0 {
		panic(errors.NoHandlers)
	}
	stop := make(chan struct{})
	bot.stopChannel = stop
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				update, err := bot.Get()
				if err == errors.Nil {
					bot.processUpdate(update)
				} else if err.Level == errors.HIGH {
					panic(err)
				}
			}
		}
	}()
	return stop
}

func (bot *Bot) StopPolling() {
	bot.GetUpdates()
	bot.stopChannel <- struct{}{}
}

func (bot *Bot) WaitPolling() {
	<-bot.stopChannel
}

func (bot *Bot) processUpdate(update utils.Update) {
	if update.Message != nil {
		msg := update.Message
		for _, handler := range bot.MessageHandlers {
			if (handler.MessageType&msg.Type() != 0 || handler.MessageType == 0) && handler.Filter(msg) {
				handler.Handler(bot, msg)
				break
			}
		}
	}
	if update.CallbackQuery != nil {
		callback := update.CallbackQuery
		for _, handler := range bot.CallbackHandlers {
			if handler.Filter(callback) {
				handler.Handler(bot, callback)
				break
			}
		}
	}
}

func (bot *Bot) LastUpdate(set ...int) int {
	if len(set) > 0 && lastUpdate < set[0] {
		lastUpdate = set[0]
	}
	return lastUpdate
}
