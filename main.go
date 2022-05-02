package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case res := <-ch1:
		fmt.Println("ch1:", res)
	case res := <-ch2:
		fmt.Println("ch2:", res)
	}

	fmt.Println("--------------------")

	out := make(chan float64)
	go func() {
		time.Sleep(20 * time.Millisecond)
		out <- math.Pi
	}()

	select {
	case res := <-out:
		fmt.Println("out:", res)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}

}
