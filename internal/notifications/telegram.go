package notifications

import (
	"Printer3DCourses/internal/bot"
	"Printer3DCourses/internal/config"
	"fmt"
	tele "gopkg.in/telebot.v4"
)

func HandleUserSubmission(cfg *config.Config, name, email, phoneNumber string) {
	text := fmt.Sprintf("🔔 Новая заявка на покупку стартового набора:\n👤 %s %s\n📧 %s", name, email, phoneNumber)
	adminId := cfg.Bot.AdminId
	_, err := bot.Bot.Send(tele.ChatID(adminId), text)
	if err != nil {
		return
	}
}
