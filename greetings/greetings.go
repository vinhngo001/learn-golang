package greetings

import (
	"fmt"
)

func Hello(name string) string{
	message := fmt.Sprintf("Hi, %v Welcome with love!", name);
	return message;
}