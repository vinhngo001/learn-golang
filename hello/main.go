package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	message := greetings.Hello("A bú");
	fmt.Println(message);
}