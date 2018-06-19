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
package botmodules

import (
	"gopkg.in/telegram-bot-api.v4"
)

// ListAdmins lists the admins of a group or supergroup in a telegram message
func List_admins(bot *tgbotapi.BotAPI, chat_id int64) ([]tgbotapi.ChatMember, error) {
	var chat_config = tgbotapi.ChatConfig{ChatID: chat_id}
	return bot.GetChatAdministrators(chat_config)
}
