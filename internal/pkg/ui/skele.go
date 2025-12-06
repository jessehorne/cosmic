package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type doer interface {
	Do()
}

type Skele struct {
	X      int
	Y      int
	W      int
	H      int
	Bounds rl.Rectangle
	doer   doer
}

func NewSkele(p doer) *Skele {
	return &Skele{
		X:    0,
		Y:    0,
		W:    30,
		H:    30,
		doer: p,
	}
}

func (s *Skele) GetBounds() rl.Rectangle {
	return s.Bounds
}

func (s *Skele) UpdateBounds() {
	s.Bounds = rl.NewRectangle(float32(s.X), float32(s.Y), float32(s.W), float32(s.H))
}

func (s *Skele) Click() {
	s.doer.Do()
}

func (s *Skele) Update() {
	s.UpdateBounds()
}

func (s *Skele) Draw() {
	if gui.Button(s.Bounds, gui.IconText(gui.ICON_PLAYER_PLAY, "")) {
		s.Click()
	}
}

func (s *Skele) Close() {

}
