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
package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"./botcommands"
	"./log"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	start_time  time.Time
	bot_api_key string
	chat_infos  []botcommands.Chat_info
)

func uptime() time.Duration {
	return time.Since(start_time)
}

func main() {
	start_time = time.Now()

	err := godotenv.Load(".env")
	if err != nil {
		log.Err("FAILED TO LOAD THE DOTENV!", "")
		os.Exit(1)
	}
	bot_api_key = os.Getenv("TELEGRAM_API_BOT_KEY")
	chat_ids := strings.Split(os.Getenv("CHATIDS"), ",")
	chat_infos = make([]botcommands.Chat_info, len(chat_ids))
	for i, chat_id := range chat_ids {
		n, err := strconv.ParseInt(chat_id, 10, 64)
		if err != nil {
			log.Err("ERROR IN EVN, INVALID CHAT ID", chat_id)
			continue
		}
		chat_infos[i] = botcommands.Chat_info{Chat_id: n}
	}

	bot, err := tgbotapi.NewBotAPI(bot_api_key)
	if err != nil {
		log.Err("FAILED TO BOOT!", err.Error())
		os.Exit(2)
	}

	bot.Debug = false

	log.Log("{{s_g}}Authorized on Account:{{s_cl}} {{s_r}}" + bot.Self.UserName + "{{s_cl}}")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Err("FAILED TO GET UPDATE", err.Error())
	}

	for update := range updates {
		log.Log("{{s_gy}}Starting Goroutine for a new UpdatePackage{{s_cl}}")
		go handle(update, bot)
	}
}

func handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.IsCommand() {

		switch update.Message.Command() {
		case "report":
			botcommands.Report(update, bot, chat_infos)
		case "chatid":
			botcommands.Get_chatid(update, bot)
		}
	}
}
