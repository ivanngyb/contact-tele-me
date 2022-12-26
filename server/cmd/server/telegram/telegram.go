package telegram

import (
	"errors"
	"server/logger"
	"time"

	tele "gopkg.in/telebot.v3"
)

type TeleBot struct {
	teleBot *tele.Bot
	chatID  int64
}

var T *TeleBot

func NewBot(token string, chatID int64) error {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	logger.L.Info().Msg("Starting telegram bot")
	b, err := tele.NewBot(pref)
	if err != nil {
		logger.L.Err(err).Msg("Failed to start telegram bot")
		return err
	}

	T = &TeleBot{
		teleBot: b,
		chatID:  chatID,
	}

	return nil
}

func (t *TeleBot) SendMessage(message string) error {
	if t.teleBot == nil {
		return errors.New("telebot not started")
	}

	_, err := t.teleBot.Send(&tele.Chat{ID: t.chatID, Type: tele.ChatPrivate}, message)
	if err != nil {
		logger.L.Err(err).Msg("Failed to send telegram message")
		return err
	}

	return nil
}
