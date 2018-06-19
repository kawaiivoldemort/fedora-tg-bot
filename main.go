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
	"time"
	"./botcommand"
	"./logging"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	startTime			time.Time
	lastReportTime		time.Time
	lastReminderTime	time.Time
)

func uptime() time.Duration {
	return time.Since(startTime)
}

func main() {
	startTime = time.Now()

	err := godotenv.Load()
	if err != nil {
		logging.Err("FAILED TO LOAD THE DOTENV!", "")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_BOT_KEY"))
	if err != nil {
		logging.Err("FAILED TO BOOT!", err.Error())
		os.Exit(2)
	}

	bot.Debug = false

	logging.Log("{{s_g}}Authorized on Account:{{s_cl}} {{s_r}}" + bot.Self.UserName + "{{s_cl}}")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		logging.Err("FAILED TO GET UPDATE", err.Error())
	}

	for update := range updates {
		logging.Log("{{s_gy}}Starting Goroutine for a new UpdatePackage{{s_cl}}")
		go handle(update, bot)
	}
}

func handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	if update.Message != nil && update.Message.IsCommand() {

		switch update.Message.Command() {
		case "report":
			if update.Message.Chat.ID == -1001038814893 {
				hour := -(time.Minute * time.Duration(60))
				checkTime := time.Now().Add(hour)
				if lastReportTime.IsZero() || checkTime.After(lastReportTime) {
					lastReportTime = time.Now()
					lastReminderTime = time.Time{}
					botcommand.ListAdmins(update, bot, logging.Log, logging.Err)
				} else if lastReminderTime.IsZero() || checkTime.After(lastReminderTime) {
					lastReminderTime = time.Now()
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This Bot will only allow this command to be used every Hour!")
					if update.Message.ReplyToMessage != nil {
						msg.BaseChat.ReplyToMessageID = update.Message.MessageID
					}
					bot.Send(msg)
				}
			}
		}
	}
}
