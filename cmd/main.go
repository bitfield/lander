package main

import (
	"fmt"
	"lander"
	"log"
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

	g.DisplayResult(os.Stdout)
}
