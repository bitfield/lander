package lander_test

import (
	"bytes"
	"lander"
	"testing"
)

func TestInitGame(t *testing.T) {
	t.Parallel()

	wantState := lander.State{
		Altitude: 100.0,
		Velocity: -100.0,
		Gravity:  -10.0,
		Thrust:   0.0,
	}
	lander.InitGame()
	if wantState != lander.GameState {
		t.Errorf("want %+v, got %+v", wantState, lander.GameState)
	}
}

func TestDisplayState(t *testing.T) {
	t.Parallel()

	lander.InitGame()
	want := "Altitude: 100.0 Velocity: -100.0 Thrust: 0.0\n"
	buf := bytes.Buffer{}
	lander.DisplayState(&buf)
	got := buf.String()

	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}
