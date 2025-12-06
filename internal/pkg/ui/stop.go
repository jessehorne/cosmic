package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type stopper interface {
	Stop()
}

type Stop struct {
	X       int
	Y       int
	W       int
	H       int
	Bounds  rl.Rectangle
	stopper stopper
}

func NewStop(s stopper) *Stop {
	return &Stop{
		X:       60,
		Y:       0,
		W:       30,
		H:       30,
		stopper: s,
	}
}

func (s *Stop) GetBounds() rl.Rectangle {
	return s.Bounds
}

func (s *Stop) UpdateBounds() {
	s.Bounds = rl.NewRectangle(float32(s.X), float32(s.Y), float32(s.W), float32(s.H))
}

func (s *Stop) Click() {
	s.stopper.Stop()
}

func (s *Stop) Update() {
	s.UpdateBounds()
}

func (s *Stop) Draw() {
	if gui.Button(s.Bounds, gui.IconText(gui.ICON_PLAYER_STOP, "")) {
		s.Click()
	}
}
