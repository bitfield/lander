package lander_test

import (
	"bytes"
	"lander"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	want := lander.Game{
		State: lander.State{
			Altitude: 100.0,
			Velocity: -100.0,
			Gravity:  -10.0,
			Thrust:   0.0,
		},
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
	g, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
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
	g, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
	want := 10.0
	input := bytes.NewBufferString("10\n")
	output := &bytes.Buffer{}
	g.ReadThrust(input, output)
	// TODO: check output
	got := g.State.Thrust
	if want != got {
		t.Errorf("Wanted %.1f, got %.1f", want, got)
	}
}

func TestUpdateWorld(t *testing.T) {
	t.Parallel()
	g, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
	want := lander.State{
		Altitude: 0.0,
		Velocity: -110.0,
		Gravity:  -10.0,
		Thrust:   0.0,
	}
	g.UpdateWorld()
	got := g.State
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSoftLandingDisplayResultSuccess(t *testing.T) {
	t.Parallel()
	g, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
	g.State.Velocity = 9.1
	want := "Safe landing (9.1m/s): you win!\n"
	buf := bytes.Buffer{}
	g.DisplayResult(&buf)
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestHardLandingDisplayResultSuccess(t *testing.T) {
	t.Parallel()
	g, err := lander.NewGame()
	if err != nil {
		t.Fatal(err)
	}
	g.State.Velocity = 150.0
	want := "Landed too hard (150.0m/s)! Try again\n"
	buf := bytes.Buffer{}
	g.DisplayResult(&buf)
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}
