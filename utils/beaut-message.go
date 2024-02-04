package utils

import "fmt"

// printColoredMessage prints a message in the specified color using ANSI escape codes
func PrintColoredMessage(message string, colorCode string) {
	// Print the message in the specified color
	fmt.Printf("\x1b[%sm%s\x1b[0m\n", colorCode, message)
}
