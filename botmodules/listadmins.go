package botmodules

import (
	"gopkg.in/telegram-bot-api.v4"
)

// ListAdmins lists the admins of a group or supergroup in a telegram message
func List_admins(bot *tgbotapi.BotAPI, chat_id int64) ([]tgbotapi.ChatMember, error) {
	var chat_config = tgbotapi.ChatConfig { ChatID: chat_id }
	return bot.GetChatAdministrators(chat_config)
}