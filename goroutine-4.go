package main

import (
	"fmt"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {

	var messages = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
