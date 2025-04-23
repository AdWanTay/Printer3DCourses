package internal

import (
	"Printer3DCourses/internal/bot"
	"Printer3DCourses/internal/config"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	tele "gopkg.in/telebot.v4"
	"time"
)

func TgApp(cfg *config.Config) error {
	pref := tele.Settings{
		Token:  cfg.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
		return err
	}

	b.Handle("/start", startHandler(cfg.Bot.AdminId))
	bot.Bot = b

	b.Start()
	return nil
}

func startHandler(adminId int64) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Sender().ID == adminId {
			keyboard := &tele.ReplyMarkup{ResizeKeyboard: true}
			return c.Send("Вы администратор", keyboard)
		}
		return nil
	}
}

func HandleUserSubmission(cfg *config.Config) {
	text := fmt.Sprintf("🔔 Новый пользователь:\n👤 %s %s\n📧 %s", "", "", "")
	adminId := cfg.Bot.AdminId
	bot.Bot.Send(tele.ChatID(adminId), text)
}
