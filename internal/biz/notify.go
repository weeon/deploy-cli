package biz

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	TelegramToken  = env("TELEGRAM_TOKEN")
	TelegramChatID = env("TELEGRAM_CHATID")
)

func TelegramNotify(msg string) {
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	cid, err := strconv.Atoi(TelegramChatID)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := tgbotapi.NewMessage(int64(cid), msg)
	m.ParseMode = "markdown"

	_, err = bot.Send(m)
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err)
		return
	}
}
