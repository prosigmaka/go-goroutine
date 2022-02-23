package main

import (
	"fmt"
	"math"
	"time"
)

func getSisi(s chan int) {
	var luasKubus float64 = 96
	//diberikan waktu jeda 10 detik agar terlihat bahwa func getLuas menunggu channel terisi
	time.Sleep(time.Second * 10)
	// ketika proses selesai nilainya dimasukkan ke channel s
	s <- int(math.Sqrt(luasKubus / 6))
}

func getLuas(s chan int) {
	// proses menerima message dan menyiman ke variable sisi
	sisi := <-s
	fmt.Println(sisi * sisi)
	close(s)
}

func main() {
	var sisi chan int = make(chan int)

	go getSisi(sisi)
	go getLuas(sisi)

}
