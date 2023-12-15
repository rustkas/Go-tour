package main

import "fmt"

func main() {
	
	fmt.Printf("Hello, %s\n", Boston)
}

func init() {
	officePlace = make(map[Office]string)
	officePlace[0] = "Cambridge, MA"
	officePlace[1] = "New York, NY"
}
type Office int

const (
	Boston Office = iota
	NewYork
)

var officePlace map[Office]string


func (o Office) String() string {
	return "Google, " + officePlace[o]
}

