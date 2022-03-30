package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("안녕하세요, %v. 환영합니다!", name)
	return message
}
