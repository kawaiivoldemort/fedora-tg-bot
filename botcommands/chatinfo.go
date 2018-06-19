package botcommands

import "time"

// Chat holds Chat metadata to respond to chat commands
type Chat_info struct {
	Chat_id            int64
	Last_report_time   time.Time
	Last_reminder_time time.Time
}
