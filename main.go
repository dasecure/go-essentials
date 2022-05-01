package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	out <- fmt.Sprintf("%s: %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org/",
		"https://golang.org/doc/",
		"https://golang.org/pkg/",
		"https://golang.org/cmd/",
	}
	// create channels
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}
