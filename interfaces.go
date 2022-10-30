package main

import (
	"fmt"
)

// Describes the interface for the engine type
// Interfaces describe methods required to satisfy
type engine interface {
	turnOn()
	turnOff()
	engineOn() bool
	getCurrentRPM() int
	getCurrentTorque() int
}

// A struct for a carEngine
type CarEngine struct {
	horsePower, maxRPM, maxTorque, currentRPM, currentTorque int
	isOn                                                     bool
}

func (c CarEngine) turnOn() {
	c.isOn = true
}

func (c CarEngine) turnOff() {
	c.isOn = false
}

func (c CarEngine) engineOn() bool {
	return c.isOn
}

func (c CarEngine) getCurrentRPM() int {
	return c.currentRPM
}

func (c CarEngine) getCurrentTorque() int {
	return c.currentTorque
}

// A struct that embeds
type car struct {
	CarEngine
	name string
}

func main() {
	fmt.Println("Playing around with interfaces, structs and embedded structs")
	civicEngine := CarEngine{horsePower: 300, maxRPM: 3000, maxTorque: 581}
	civic := car{civicEngine, "Birdy"}
	fmt.Println(civic.name)
	fmt.Println(civic.engineOn())
}
