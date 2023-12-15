package main

import (
	"fmt"
	"hash/crc32"
	
)

func main() {
	h := crc32.NewIEEE()
	fmt.Fprintf(h, "Hello, Anatolii\n")
	fmt.Printf("hash=%#x\n", h.Sum32())
}