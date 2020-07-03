package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	var ch = make(chan int)
	go kirimKabar(ch)
	terimaKabar(ch)

}

func kirimKabar(ch chan int) {
	for i := 0; true; i++ {
		ch <- i
		// Jeda sebelum data terakhir di terima
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func terimaKabar(ch chan int) {
loop:
	for {
		select {
		case <-ch:
			fmt.Println("Apa Kabar Teman Teman")
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout, tidak ada aktifitas dalam 5 detik")
			break loop
		}
	}
}
