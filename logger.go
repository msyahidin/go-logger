package go_logger

import "fmt"

func Error(errorMessage string) {
	_ = fmt.Errorf(errorMessage)
}

func Info(message string) {
	fmt.Println(message)
}
