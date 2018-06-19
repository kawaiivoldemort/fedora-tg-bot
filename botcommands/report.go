package botcommands

import (
	"fmt"
	"strings"
	"time"

	"../botmodules"
	"../log"
	"gopkg.in/telegram-bot-api.v4"
)

func get_chat(chats []Chat_info, chat_id int64) *Chat_info {
	for _, val := range chats {
		if chat_id == val.Chat_id {
			return &val
		}
	}
	return nil
}

// Report notifies an occurence in a group to admins
func Report(update tgbotapi.Update, bot *tgbotapi.BotAPI, chats []Chat_info) {
	if update.Message.Chat != nil {
		var chat_info = get_chat(chats, update.Message.Chat.ID)
		if chat_info != nil {
			// Do a time check so you dont spam the admins
			hour := -(time.Minute * time.Duration(60))
			check_time := time.Now().Add(hour)
			if chat_info.Last_report_time.IsZero() || check_time.After(chat_info.Last_report_time) {
				chat_info.Last_report_time = time.Now()
				chat_info.Last_reminder_time = time.Time{}
				// Get Chat Administrators
				message := update.Message
				chat := message.Chat
				admins, err := botmodules.List_admins(bot, chat.ID)
				if err != nil {
					log.Err("FAILED TO GET CHAT ADMINS", err.Error())
					return
				}
				// Make an Output of the Admins
				var print_buffer strings.Builder
				for _, admin := range admins {
					if admin.User.UserName != "" {
						print_buffer.WriteString(fmt.Sprintf("@%s  ", admin.User.UserName))
					} else {
						var last_name string
						if admin.User.LastName != "" {
							last_name = " " + admin.User.LastName
						} else {
							last_name = ""
						}
						print_buffer.WriteString(fmt.Sprintf("[%s %s](tg://user?id=%d)  ", admin.User.FirstName, last_name, admin.User.ID))
					}
				}
				// Print the Admins
				response := tgbotapi.NewMessage(chat.ID, print_buffer.String())
				if message.ReplyToMessage != nil {
					response.BaseChat.ReplyToMessageID = message.ReplyToMessage.MessageID
				}
				_, err = bot.Send(response)
				if err != nil {
					log.Err("FAILED TO SEND RESPONSE", err.Error())
				}
			} else if chat_info.Last_reminder_time.IsZero() || check_time.After(chat_info.Last_reminder_time) {
				// Print no pester Message
				chat_info.Last_reminder_time = time.Now()
				response := tgbotapi.NewMessage(update.Message.Chat.ID, "This Bot will only allow this command to be used every Hour!")
				if update.Message.ReplyToMessage != nil {
					response.BaseChat.ReplyToMessageID = update.Message.MessageID
				}
				bot.Send(response)
			}
		}
	}
}
