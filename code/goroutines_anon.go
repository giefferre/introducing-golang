package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(6 * time.Second)
}

// START1 OMIT
func main() {
	rand.Seed(time.Now().UTC().UnixNano()) //OMIT
	go func(phrase string) {
		randomTime := time.Duration(rand.Intn(5000)) * time.Millisecond
		time.Sleep(randomTime)

		fmt.Println(phrase)
	}("Once upon a time...")

	pause() // waits for a while
	fmt.Println("Quit!")
}

// STOP1 OMIT
