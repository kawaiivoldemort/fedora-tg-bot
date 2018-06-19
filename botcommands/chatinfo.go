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

import "time"

// Chat holds Chat metadata to respond to chat commands
type Chat_info struct {
	Chat_id            int64
	Last_report_time   time.Time
	Last_reminder_time time.Time
}
