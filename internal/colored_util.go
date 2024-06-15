package internal

import "fmt"

var resetColor = "\033[0m"
var redColor = "\033[31m"
var greeColor = "\033[32m"
var yellowColor = "\033[33m"

func LogError(text string) {
	fmt.Println(redColor + text + resetColor)
}

// Print in new line with yellow color
func LogMedium(text string) {
	fmt.Println(yellowColor + text + resetColor)
}

func LogSuccess(text string) {
	fmt.Println(greeColor + text + resetColor)
}
