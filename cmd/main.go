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
	altitude := 100.0
	velocity := -100.0
	gravity := -10.0
	thrust := 0.0
	for altitude > 0 {
		fmt.Println("Altitude: ", altitude, "Velocity: ", velocity, "Thrust:", thrust)
		fmt.Print("Enter thrust setting (m/s/s): ")
		fmt.Scanln(&thrust)
		altitude += velocity
		velocity += gravity
		velocity += thrust
	}
	if math.Abs(velocity) <= 10 {
		fmt.Printf("Safe landing (%.1fm/s): you win!\n", velocity)
	} else {
		fmt.Printf("Landed too hard (%.1fm/s)! Try again\n", velocity)
	}
}