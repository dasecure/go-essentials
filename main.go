// Open a text file and write into it

package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("./hellow", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
