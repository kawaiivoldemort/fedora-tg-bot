package botcommand

import (
	"fmt"
	"strings"
	"gopkg.in/telegram-bot-api.v4"
)

// ListAdmins lists the admins of a group or supergroup in a telegram message
func ListAdmins(update tgbotapi.Update, bot *tgbotapi.BotAPI, logFn func(string), errFn func(string, string)) {
	if(update.Message.Chat != nil) {
		// Get Chat Administrators
		var message = update.Message
		var chat = message.Chat
		var chatConfig = tgbotapi.ChatConfig { ChatID: chat.ID }
		var admins, err = bot.GetChatAdministrators(chatConfig)
		if err != nil {
			errFn("FAILED TO GET CHAT ADMINS", err.Error())
		} else {
			// Make an Output of the Admins
			var printBuffer strings.Builder
			for _, admin := range admins {
				if admin.User.UserName != "" {
					printBuffer.WriteString(admin.User.UserName)
					printBuffer.WriteString(" ")
				} else {
					var lastName string
					if(admin.User.LastName != "") {
						lastName = " " + admin.User.LastName
					} else {
						lastName = ""
					}
					printBuffer.WriteString(fmt.Sprintf("[%s %s](tg://user?id=%d) ", admin.User.FirstName, lastName, admin.User.ID))
				}
			}
			// Print the Admins
			var response = tgbotapi.NewMessage(chat.ID, printBuffer.String())
			if message.ReplyToMessage != nil {
				response.BaseChat.ReplyToMessageID = message.ReplyToMessage.MessageID
			}
			_, err = bot.Send(response)
			if err != nil {
				// Do Something
			}
		}
	}
}