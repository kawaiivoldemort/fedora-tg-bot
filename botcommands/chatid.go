package botcommands

import (
	"fmt"

	"../botmodules"
	"../log"
	"gopkg.in/telegram-bot-api.v4"
)

// Get_chatid is a dev function for admins to get Chat IDs for their bots
func Get_chatid(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatid := update.Message.Chat.ID
	admins, err := botmodules.List_admins(bot, chatid)
	if err != nil {
		log.Err("FAILED TO GET CHAT ADMINS", err.Error())
		return
	}
	// Make an Output of the Admins
	for _, admin := range admins {
		if admin.User.ID == update.Message.From.ID {
			msg := tgbotapi.NewMessage(chatid, fmt.Sprintf("%d", chatid))
			if update.Message.ReplyToMessage != nil {
				msg.BaseChat.ReplyMarkup = update.Message.ReplyToMessage.MessageID
			}
			_, err := bot.Send(msg)
			if err != nil {
				log.Err("FAILED TO SEND MESSAGE", err.Error())
			}
			break
		}
	}
}
