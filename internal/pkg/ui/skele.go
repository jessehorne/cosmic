package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type doer interface {
	Do()
}

type Skele struct {
	X      int32
	Y      int32
	W      int32
	H      int32
	Bounds rl.Rectangle
	doer   doer
}

func NewSkele(p doer) *Skele {
	sk := &Skele{
		X:    0,
		Y:    0,
		W:    30,
		H:    30,
		doer: p,
	}
	sk.UpdateBounds()
	return sk
}

func (s *Skele) GetBounds() rl.Rectangle {
	return s.Bounds
}

func (s *Skele) SetBounds(r rl.Rectangle) {
	s.Bounds = r
	s.X = int32(s.Bounds.X)
	s.Y = int32(s.Bounds.Y)
	s.W = int32(s.Bounds.Width)
	s.H = int32(s.Bounds.Height)
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
