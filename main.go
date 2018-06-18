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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	start_time         time.Time
	last_report_time   time.Time
	last_reminder_time time.Time
	ps                 = print_shell_string
	// The array describing the exact Replacements for Shell colors
	replace_map = []string{
		"{{shell_blue}}", color_blue,
		"{{s_b}}", color_blue,
		"{{shell_clear}}", color_clear,
		"{{s_cl}}", color_clear,
		"{{shell_cyan}}", color_cyan,
		"{{s_c}}", color_cyan,
		"{{shell_green}}", color_green,
		"{{s_g}}", color_green,
		"{{shell_magenta}}", color_magenta,
		"{{s_m}}", color_magenta,
		"{{shell_red}}", color_red,
		"{{s_r}}", color_red,
		"{{shell_white}}", color_white,
		"{{s_w}}", color_white,
		"{{shell_yellow}}", color_yellow,
		"{{s_y}}", color_yellow,
		"{{shell_grey}}", color_grey,
		"{{s_gy}}", color_grey,
	}

	color_replace = strings.NewReplacer(replace_map...)
)

const (
	//color_blue Blue Shell Color
	color_blue = "\x1b[34;1m"
	//color_clear Clears the last set color.
	color_clear = "\x1b[0m"
	//ColorC Unknown color
	color_cyan = "\x1b[36;1m"
	//color_green Green Shell Color
	color_green = "\x1b[32;1m"
	//color_magenta Magenta Shell Color
	color_magenta = "\x1b[35;1m"
	//color_red Red Shell Color
	color_red = "\x1b[31;1m"
	//color_white White Shell Color
	color_white = "\x1b[37;1m"
	//color_yellow Yelow Shell Color
	color_yellow = "\x1b[33;1m"
	//ColorZero Unknown Color
	color_grey = "\x1b[30;1m"
)

func ready_shell_string(input string) string {
	return color_replace.Replace(input)
}

//print_shell_string will print the given line, with replacement of colors
func print_shell_string(input string) {
	shell_content := append_time(ready_shell_string(input))
	fmt.Println(shell_content)
}

func append_time(input string) string {
	time_string := ready_shell_string(time.Now().Format("{{s_w}}[{{s_cl}}" +
		"{{s_m}}02.01{{s_cl}}" +
		"{{s_gy}}.2006{{s_cl}}" +
		" {{s_g}}15:04{{s_cl}}" +
		"{{s_gy}}:05{{s_cl}}" +
		"{{s_w}}]{{s_cl}}"))
	return time_string + " " + input
}

func uptime() time.Duration {
	return time.Since(start_time)
}

func main() {
	start_time = time.Now()

	env_err := godotenv.Load()
	if env_err != nil {
		ps("{{s_r}}FAILED TO LOAD THE DOTENV!{{s_cl}}")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_BOT_KEY"))
	if err != nil {
		ps("{{s_r}}FAILED TO BOOT!{{s_cl}}")
		ps("{{s_r}}" + err.Error() + "{{s_cl}}")
		os.Exit(2)
	}

	bot.Debug = false

	ps("{{s_g}}Authorized on Account:{{s_cl}} {{s_r}}" + bot.Self.UserName + "{{s_cl}}")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		ps("{{s_r}}ERROR! {{s_cl}} " + err.Error())
	}

	for update := range updates {
		ps("{{s_gy}}Starting Goroutine for a new UpdatePackage{{s_cl}}")
		go handle(update, bot)
	}
}

func handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	if update.Message != nil && update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "report":
			if update.Message.Chat.ID == -1001038814893 {
				one_hour := -(time.Minute * time.Duration(60))
				check_time := time.Now().Add(one_hour)
				if last_report_time.IsZero() || check_time.After(last_report_time) {
					last_report_time = time.Now()
					last_reminder_time = time.Time{}
					msg.Text = "@jflory @Kohane @sesivany @ignatenkobrain @bexelbie @michalrud @AnXh3L0 @gwindor @buredoRUNofthecyborg @x3mboy"
					bot.Send(msg)
				} else if last_reminder_time.IsZero() || check_time.After(last_reminder_time) {
					last_reminder_time = time.Now()
					msg.Text = "This Bot will only allow this command to be used every Hour!"
					bot.Send(msg)
				}
			}
		}
	}
}
