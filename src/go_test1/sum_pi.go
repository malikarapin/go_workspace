package main

import (
	"fmt"
	"math"
)

func partialSum(c chan<- float64, kStart int, kOffset int, amount int) {
	// Using the Bailey–Borwein–Plouffe formula

	var sum float64
	for k := float64(kStart); k < float64(kStart+kOffset*amount); k += float64(kOffset) {
		sum += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
	}

	c <- sum
}

func main() {
	var pi float64

	const workers int = 5
	const kOffset int = workers
	const amount int = 10

	doneChannel := make(chan bool)

	for i := 0; i < workers; i++ {
		go (func(kStart int) {
			c := make(chan float64)
			go partialSum(c, kStart, kOffset, amount)
			pi += <-c
			doneChannel <- true
		})(i)
	}

	for i := 0; i < workers; i++ {
		<-doneChannel
	}

	fmt.Println("Calculated π:", pi)
	fmt.Println("Difference between math.Pi and calculated:", pi-math.Pi)
}
