package lander_test

import (
	"bytes"
	"lander"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(got)) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDisplayState(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}

	g, err := lander.NewGame(lander.WithOutput(buf))
	if err != nil {
		t.Fatal(err)
	}
 
	want := "Altitude: 100.0 Velocity: -100.0 Thrust: 0.0\n"
	
	g.DisplayState()
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestReadThrust(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString("10\n")
	output := &bytes.Buffer{}

	g, err := lander.NewGame(
		lander.WithInput(input),
		lander.WithOutput(output),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 10.0
	
	g.ReadThrust()
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

	buf := &bytes.Buffer{}
	g, err := lander.NewGame(lander.WithOutput(buf))
	if err != nil {
		t.Fatal(err)
	}
	g.State.Velocity = 9.1
	want := "Safe landing (9.1m/s): you win!\n"
	
	g.DisplayResult()
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestHardLandingDisplayResultSuccess(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	g, err := lander.NewGame(lander.WithOutput(buf))
	if err != nil {
		t.Fatal(err)
	}
	g.State.Velocity = 150.0
	want := "Landed too hard (150.0m/s)! Try again\n"
	
	g.DisplayResult()
	got := buf.String()
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestPlayGame(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString("10\n")
	output := &bytes.Buffer{}
	
	lander.PlayGame(
		lander.WithInput(input),
		lander.WithOutput(output),
	)

	got := output.String()
	want := "Land safely at a vertical speed of 10m/s or less.Altitude: 100.0 Velocity: -100.0 Thrust: 0.0\nEnter thrust setting (m/s/s): Landed too hard (-100.0m/s)! Try again\n"
	
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}