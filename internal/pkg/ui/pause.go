package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type pauser interface {
	Pause()
}

type Pause struct {
	X      int
	Y      int
	W      int
	H      int
	Bounds rl.Rectangle
	pauser pauser
}

func NewPause(p pauser) *Pause {
	return &Pause{
		X:      30,
		Y:      0,
		W:      30,
		H:      30,
		pauser: p,
	}
}

func (p *Pause) GetBounds() rl.Rectangle {
	return p.Bounds
}

func (p *Pause) UpdateBounds() {
	p.Bounds = rl.NewRectangle(float32(p.X), float32(p.Y), float32(p.W), float32(p.H))
}

func (p *Pause) Click() {
	p.pauser.Pause()
}

func (p *Pause) Update() {
	p.UpdateBounds()
}

func (p *Pause) Draw() {
	if gui.Button(p.Bounds, gui.IconText(gui.ICON_PLAYER_PAUSE, "")) {
		p.Click()
	}
}
