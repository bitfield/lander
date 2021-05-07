package lander

import (
	"fmt"
	"io"
	"math"
	"os"
)

type State struct {
	Altitude, Velocity, Gravity, Thrust float64
}

type Game struct {
	State State
	output io.Writer
	input io.Reader
}

func NewGame(opts ...option) (Game, error) {

	g := Game{
		State: State{
			Altitude: 100.0,
			Velocity: -100.0,
			Gravity:  -10.0,
			Thrust:   0.0,
		},
		output: os.Stdout,
		input: os.Stdin,
	}

	for _, o := range opts {
		if err := o(&g); err != nil {
			return Game{}, err
		}
	}

	return g, nil
}

func (g Game) DisplayState() {
	fmt.Fprintf(g.output, "Altitude: %.1f Velocity: %.1f Thrust: %.1f\n", g.State.Altitude, g.State.Velocity, g.State.Thrust)
}

func (g Game) DisplayResult() {
	if math.Abs(g.State.Velocity) <= 10 {
		fmt.Fprintf(g.output, "Safe landing (%.1fm/s): you win!\n", g.State.Velocity)
	} else {
		fmt.Fprintf(g.output, "Landed too hard (%.1fm/s)! Try again\n", g.State.Velocity)
	}
}

func (g *Game) ReadThrust() {
	fmt.Fprint(g.output, "Enter thrust setting (m/s/s): ")
	fmt.Fscanln(g.input, &g.State.Thrust)
}

func (g *Game) UpdateWorld() {
	g.State.Altitude += g.State.Velocity
	g.State.Velocity += g.State.Gravity
	g.State.Velocity += g.State.Thrust
}

func (g Game) DisplayWelcome() {
	fmt.Fprintf(g.output, "Land safely at a vertical speed of 10m/s or less.")
}

type option func(*Game) error

func WithInput(r io.Reader) option {
	return func(g *Game) error {
		g.input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(g *Game) error {
		g.output = w
		return nil
	}
}

func PlayGame(opts ...option) error {
	g, err := NewGame(opts...)
	if err != nil {
		return err
	}
	
	g.DisplayWelcome()

	for g.State.Altitude > 0 {
		g.DisplayState()
		g.ReadThrust()
		g.UpdateWorld()
	}

	g.DisplayResult()

	return nil
}
