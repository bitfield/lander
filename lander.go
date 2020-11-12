package lander

type State struct {
	Altitude, Velocity, Gravity, Thrust float64
}

var GameState State

func InitGame() {}