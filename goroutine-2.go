package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

}
func printText() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("text", i)
	}

}

func main() {
	start := time.Now() // start time

	printNumber()

	printText()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println(time.Since(start)) // lamanya waktu

}
