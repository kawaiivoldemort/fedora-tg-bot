package log

import (
	"strings"
	"time"
	"fmt"
)

// COLORIZE

// Human Readable color constants
const (
	// color_blue Blue Shell Color
	color_blue = "\x1b[34;1m"
	// color_clear Clears the last set color.
	color_clear = "\x1b[0m"
	// color_cyan Unknown color
	color_cyan = "\x1b[36;1m"
	// color_green Green Shell Color
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

var (
	// The array describing the exact Replacements for Shell colors
	replace_map = []string{
		"{{shell_blue}}",		color_blue,
		"{{s_b}}",				color_blue,
		"{{shell_clear}}",		color_clear,
		"{{s_cl}}",				color_clear,
		"{{shell_cyan}}",		color_cyan,
		"{{s_c}}",				color_cyan,
		"{{shell_green}}",		color_green,
		"{{s_g}}",				color_green,
		"{{shell_magenta}}",	color_magenta,
		"{{s_m}}",				color_magenta,
		"{{shell_red}}",		color_red,
		"{{s_r}}",				color_red,
		"{{shell_white}}",		color_white,
		"{{s_w}}",				color_white,
		"{{shell_yellow}}",		color_yellow,
		"{{s_y}}",				color_yellow,
		"{{shell_grey}}",		color_grey,
		"{{s_gy}}",				color_grey,
	}
	// Color Map Replacer
	color_replace = strings.NewReplacer(replace_map...)
)

// Applies the color mapper to colorize the string
func colorize_string(input string) string {
	return color_replace.Replace(input)
}

// TIME STRING

func get_time_string() string {
	time_string := time.Now().Format("{{s_w}}[{{s_cl}}" +
		"{{s_m}}02.01{{s_cl}}"	+
		"{{s_gy}}.2006{{s_cl}}"	+
		" {{s_g}}15:04{{s_cl}}"	+
		"{{s_gy}}:05{{s_cl}}"	+
		"{{s_w}}]{{s_cl}}")
	return time_string
}

// LOGGER INTERFACE

// Log logs the output message
func Log(input string) {
	var log_str = fmt.Sprintf("%s: %s", get_time_string(), input)
	fmt.Println(colorize_string(log_str))
}

// Err logs the error message
func Err(input string, details string) {
	var log_str = fmt.Sprintf("%s: {{s_r}}ERROR:{{s_cl}} %s", get_time_string(), input)
	if(details != "") {
		log_str += " " + details
	}
	fmt.Println(colorize_string(log_str))
}