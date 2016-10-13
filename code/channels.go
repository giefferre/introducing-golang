package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(20 * time.Second)
}

// START1 OMIT
func beacon(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("[Beacon] Sending Ping")
		ch <- "Ping"
	}
}

func ship(ch chan string) {
	for {
		fmt.Println("[Ship] Received", <-ch)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // OMIT
	channel := make(chan string)

	go beacon(channel)
	go ship(channel)

	pause()
	fmt.Println("[Main] Quit!")
}

// STOP1 OMIT
