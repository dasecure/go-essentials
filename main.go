package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// ch <- 1234
	start := time.Now()
	go func() {
		ch <- 5678
	}()
	val := <-ch
	elapsed := time.Since(start)
	fmt.Println(val, elapsed)

	fmt.Println("=====")

	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		// val := <-ch
		fmt.Printf("Receiving %d\n", i)
	}

	ch2 := make(chan int, 1)
	ch2 <- 123
	val2 := <-ch2
	fmt.Println(val2)
}
