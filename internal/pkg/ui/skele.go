package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Skele struct {
	*Core
}

func NewSkele() *Skele {
	return &Skele{
		Core: NewCore(rl.NewRectangle(0, 0, 30, 30)),
	}
}

func (s *Skele) GetCore() *Core {
	return s.Core
}

func (s *Skele) Click() {

}

func (s *Skele) Update() {

}

func (s *Skele) Draw() {
	if gui.Button(s.Bounds, gui.IconText(gui.ICON_PLAYER_PLAY, "")) {
		s.Click()
	}
}

func (s *Skele) Close() {

}
