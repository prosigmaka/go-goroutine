package main

import (
	"fmt"
	"sync"
	// "time"
)

func print(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}

/* 	add - ada berapa go routine
   	wait - menunggu
   	done - selesai */

func main() {
	var wg sync.WaitGroup

	// wg.Add(1)
	// atau
	// wg.Add(2)
	go print("Hallo", &wg)
	// wg.Add(1)
	go print("Salam", &wg)

	wg.Wait()
	// time.Sleep(100 * time.Millisecond)
}
