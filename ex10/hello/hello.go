package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org/"}
	data, err := json.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
