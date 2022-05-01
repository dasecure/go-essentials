package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type")
}

func siteConcurrent2(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println("**", url, returnType(url))
		}(url)
	}
	wg.Wait()
}

func siteConcurrent1(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ch <- returnType(url)
		}(url)
	}
	for range urls {
		fmt.Println(">>", <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func siteSerial(urls []string) {
	for _, url := range urls {
		fmt.Println(returnType(url))
	}
}

func main() {
	urls := []string{
		"https://www.golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSerial(urls)
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime elapsed: %s\n", elapsed)

	start = time.Now()
	siteConcurrent1(urls)
	elapsed = time.Since(start)
	fmt.Printf("\n\nConcurrent Time elapsed: %s\n", elapsed)
	siteConcurrent2(urls)
}
