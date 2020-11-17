package lander

import (
	"fmt"
	"io"
)

type State struct {
	Altitude, Velocity, Gravity, Thrust float64
}

var GameState State

// InitGame ...
func InitGame() {
	GameState = State{
		Altitude: 100.0,
		Velocity: -100.0,
		Gravity:  -10.0,
		Thrust:   0.0,
	}
}

func DisplayState(w io.Writer) {
	fmt.Fprintf(w, "Altitude: %.1f Velocity: %.1f Thrust: %.1f\n", GameState.Altitude, GameState.Velocity, GameState.Thrust)
}
