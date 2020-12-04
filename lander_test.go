package lander_test

import (
	"bytes"
	"lander"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	want := lander.Game{
		Altitude: 100.0,
		Velocity: -100.0,
		Gravity:  -10.0,
		Thrust:   0.0,
	}
	got, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %+v, got %+v", want, got)
	}
}

func TestDisplayState(t *testing.T) {
	t.Parallel()
	// don't care about error
	g, _ := lander.NewGame()
	want := "Altitude: 100.0 Velocity: -100.0 Thrust: 0.0\n"
	buf := bytes.Buffer{}
	g.DisplayState(&buf)
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestReadThrust(t *testing.T) {
	t.Parallel()
	// don't care about error
	g, _ := lander.NewGame()
	want := 10.0
	input := bytes.NewBufferString("10.0\n")
	output := &bytes.Buffer{}
	g.ReadThrust(input, output)
	// TODO: check output
	got := g.Thrust
	if want != got {
		t.Errorf("Wanted %.1f, got %.1f", want, got)
	}
}