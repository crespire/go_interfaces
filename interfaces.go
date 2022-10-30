package main

import (
	"fmt"
)

// Describes the interface for the engine type
type engine interface {
	turnOn() bool
	turnOff() bool
	currentRPM() string
	currentTorque() string
}

// A struct for a carEngine
type carEngine struct {
	horsePower, maxRPM, maxTorque int
}

// A struct that embeds
type car struct {
	carEngine
}

func main() {
	fmt.Println("Playing around with interfaces, structs and embedded structs")
}
