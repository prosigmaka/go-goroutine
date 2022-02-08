package main

import (
	"fmt"
)

func printSalam() {
	fmt.Println("Hallo!")
}

func main() {

	/* tanpa go-routine */
	// printSalam()
	// fmt.Println("Selamat datang!!!")

	go printSalam()

	/* waktu untuk menunggu aplikasi berhenti setelah menjaklankan go-routine */
	// time.Sleep(100 * time.Millisecond)

	fmt.Println("Selamat datang!!!")

	/* stop aplikasi dengan user klik enter */
	// var input string
	// fmt.Scanln(&input)

}
