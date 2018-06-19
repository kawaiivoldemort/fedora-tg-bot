/*
   This file is part of the Fedora Telegram Report Bot.

   The Fedora Telegram Report Bot is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   The Fedora Telegram Report Bot is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with the Fedora Telegram Report Bot.  If not, see <https://www.gnu.org/licenses/>.
*/
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
