package main

import (
	"fmt"
	"lander"
	"log"
	"math"
	"os"
)

// TODO: fuel
// TODO: downrange velocity
// TODO: real time controls
// TODO: graphics!
// TODO: randomise starting position
// TODO: last-minute diversion

func main() {
	fmt.Println("Land safely at a vertical speed of 10m/s or less.")
	g, err := lander.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	for g.Altitude > 0 {
		g.DisplayState(os.Stdout)
		g.ReadThrust(os.Stdin, os.Stdout)
		g.UpdateWorld()
		// g.Altitude += g.Velocity
		// g.Velocity += g.Gravity
		// g.Velocity += g.Thrust
	}
	if math.Abs(velocity) <= 10 {
		fmt.Printf("Safe landing (%.1fm/s): you win!\n", velocity)
	} else {
		fmt.Printf("Landed too hard (%.1fm/s)! Try again\n", velocity)
	}
}
