package lander

import (
	"fmt"
	"io"
)

type State struct {
	Altitude, Velocity, Gravity, Thrust float64
}

type Game struct {
	State State
}

func NewGame() (Game, error) {
	return Game{
		State: State{
			Altitude: 100.0,
			Velocity: -100.0,
			Gravity:  -10.0,
			Thrust:   0.0,
		},
	}, nil
}

func (g Game) DisplayState(w io.Writer) {
	fmt.Fprintf(w, "Altitude: %.1f Velocity: %.1f Thrust: %.1f\n", g.State.Altitude, g.State.Velocity, g.State.Thrust)
}

func (g *Game) ReadThrust(input io.Reader, output io.Writer) {
	fmt.Fprint(output, "Enter thrust setting (m/s/s): ")
	fmt.Fscanln(input, &g.State.Thrust)
}

func (g *Game) UpdateWorld() {
	g.State.Altitude += g.State.Velocity
	g.State.Velocity += g.State.Gravity
	g.State.Velocity += g.State.Thrust
}