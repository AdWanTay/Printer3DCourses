package notifications

import (
	"Printer3DCourses/internal/bot"
	"Printer3DCourses/internal/config"
	"fmt"
	tele "gopkg.in/telebot.v4"
)

func HandleUserSubmission(cfg *config.Config) {
	text := fmt.Sprintf("🔔 Новый пользователь:\n👤 %s %s\n📧 %s", "", "", "")
	adminId := cfg.Bot.AdminId
	bot.Bot.Send(tele.ChatID(adminId), text)
}
