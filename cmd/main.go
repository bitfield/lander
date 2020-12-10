package main

import (
	"fmt"
	"lander"
	"log"
	"math"
	"os"
)

func main() {
	fmt.Println("Land safely at a vertical speed of 10m/s or less.")
	g, err := lander.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	for g.State.Altitude > 0 {
		g.DisplayState(os.Stdout)
		g.ReadThrust(os.Stdin, os.Stdout)
		g.UpdateWorld()
	}
	if math.Abs(g.State.Velocity) <= 10 {
		fmt.Printf("Safe landing (%.1fm/s): you win!\n", g.State.Velocity)
	} else {
		fmt.Printf("Landed too hard (%.1fm/s)! Try again\n", g.State.Velocity)
	}
}
