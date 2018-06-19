package logging

import (
	"strings"
	"time"
	"fmt"
)

// COLORIZE

// Human Readable color constants
const (
	// colorBlue Blue Shell Color
	colorBlue = "\x1b[34;1m"
	// colorClear Clears the last set color.
	colorClear = "\x1b[0m"
	// colorCyan Unknown color
	colorCyan = "\x1b[36;1m"
	// colorGreen Green Shell Color
	colorGreen = "\x1b[32;1m"
	//colorMagenta Magenta Shell Color
	colorMagenta = "\x1b[35;1m"
	//colorRed Red Shell Color
	colorRed = "\x1b[31;1m"
	//colorWhite White Shell Color
	colorWhite = "\x1b[37;1m"
	//colorYellow Yelow Shell Color
	colorYellow = "\x1b[33;1m"
	//ColorZero Unknown Color
	colorGrey = "\x1b[30;1m"
)

var (
	// The array describing the exact Replacements for Shell colors
	replaceMap = []string{
		"{{shell_blue}}",		colorBlue,
		"{{s_b}}",				colorBlue,
		"{{shell_clear}}",		colorClear,
		"{{s_cl}}",				colorClear,
		"{{shell_cyan}}",		colorCyan,
		"{{s_c}}",				colorCyan,
		"{{shell_green}}",		colorGreen,
		"{{s_g}}",				colorGreen,
		"{{shell_magenta}}",	colorMagenta,
		"{{s_m}}",				colorMagenta,
		"{{shell_red}}",		colorRed,
		"{{s_r}}",				colorRed,
		"{{shell_white}}",		colorWhite,
		"{{s_w}}",				colorWhite,
		"{{shell_yellow}}",		colorYellow,
		"{{s_y}}",				colorYellow,
		"{{shell_grey}}",		colorGrey,
		"{{s_gy}}",				colorGrey,
	}
	// Color Map Replacer
	colorReplace = strings.NewReplacer(replaceMap...)
)

// Applies the color mapper to colorize the string
func colorizeString(input string) string {
	return colorReplace.Replace(input)
}

// TIME STRING

func getTimeString() string {
	timeString := time.Now().Format("{{s_w}}[{{s_cl}}" +
		"{{s_m}}02.01{{s_cl}}"	+
		"{{s_gy}}.2006{{s_cl}}"	+
		" {{s_g}}15:04{{s_cl}}"	+
		"{{s_gy}}:05{{s_cl}}"	+
		"{{s_w}}]{{s_cl}}")
	return timeString
}

// LOGGER INTERFACE

// Log logs the output message
func Log(input string) {
	var logStr = fmt.Sprintf("%s: %s", getTimeString(), input)
	fmt.Println(colorizeString(logStr))
}

// Err logs the error message
func Err(input string, details string) {
	var logStr = fmt.Sprintf("%s: {{s_r}}ERROR:{{s_cl}} %s", getTimeString(), input)
	if(details != "") {
		logStr += " " + details
	}
	fmt.Println(colorizeString(logStr))
}