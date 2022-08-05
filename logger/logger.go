package logger

import "log"

type Logger struct {
}

func Error(message string) {
	printWithLevel("Error", message)
}

func printWithLevel(level, message string) {
	log.Printf("%v: %v", level, message)
}
