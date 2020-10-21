package main

import (
	"fmt"
	"math"
)

// TODO: fuel
// TODO: downrange velocity
// TODO: real time controls
// TODO: graphics!
// TODO: randomise starting position
// TODO: last-minute diversion

func main() {
	fmt.Println("Land safely at a vertical speed of 10m/s or less.")
	alt := 100.0
	velocity := -100.0
	acceleration := -10.0
	thrust := 0.0
	for alt > 0 {
		fmt.Println("Altitude: ", alt, "Velocity: ", velocity, "Thrust:", thrust)
		fmt.Print("Enter thrust setting (m/s/s): ")
		fmt.Scanln(&thrust)
		alt += velocity
		velocity += acceleration
		velocity += thrust
	}
	if math.Abs(velocity) <= 10 {
		fmt.Printf("Safe landing (%.1fm/s): you win!\n", velocity)
	} else {
		fmt.Printf("Landed too hard (%.1fm/s)! Try again\n", velocity)
	}
}