package lander

import (
	"fmt"
	"io"
)

type Game struct {
	Altitude, Velocity, Gravity, Thrust float64
}

func NewGame() (Game, error) {
	return Game{
		Altitude: 100.0,
		Velocity: -100.0,
		Gravity:  -10.0,
		Thrust:   0.0,
	}, nil
}

func (g Game) DisplayState(w io.Writer) {
	fmt.Fprintf(w, "Altitude: %.1f Velocity: %.1f Thrust: %.1f\n", g.Altitude, g.Velocity, g.Thrust)
}

func (g *Game) ReadThrust(input io.Reader, output io.Writer) {
	fmt.Fprint(output, "Enter thrust setting (m/s/s): ")
	fmt.Fscanln(input, &g.Thrust)
}