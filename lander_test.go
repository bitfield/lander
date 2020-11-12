package lander_test

import (
	"lander"
	"testing"
)

func TestInitGame(t *testing.T) {
	wantState := lander.State{
		Altitude:  100.0,
		Velocity:  -100.0,
		Gravity:  -10.0,
		Thrust:  0.0,
	}
	lander.InitGame()
	if wantState != lander.GameState {
		t.Errorf("want %+v, got %+v", wantState, lander.GameState)
	}
}