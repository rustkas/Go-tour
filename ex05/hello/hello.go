package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	fetch("https://www.google.com/")
	fmt.Println(time.Since(start))
}

func fetch(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Printf("Fetched %v bytes from %s\n", len(body), url)
}
