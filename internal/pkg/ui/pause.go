package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type pauser interface {
	Pause()
}

type Pause struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Toggled bool
	Bounds  rl.Rectangle
	pauser  pauser
}

func NewPause(p pauser) *Pause {
	pa := &Pause{
		W:      30,
		H:      30,
		pauser: p,
	}
	pa.UpdateBounds()
	return pa
}

func (p *Pause) GetBounds() rl.Rectangle {
	return p.Bounds
}

func (p *Pause) SetBounds(r rl.Rectangle) {
	p.Bounds = r
	p.X = int32(p.Bounds.X)
	p.Y = int32(p.Bounds.Y)
	p.W = int32(p.Bounds.Width)
	p.H = int32(p.Bounds.Height)
}

func (p *Pause) UpdateBounds() {
	p.Bounds = rl.NewRectangle(float32(p.X), float32(p.Y), float32(p.W), float32(p.H))
}

func (p *Pause) Click() {
	p.pauser.Pause()
}

func (p *Pause) SetToggle(w bool) {
	p.Toggled = w
}

func (p *Pause) Update() {
	p.UpdateBounds()
}

func (p *Pause) Draw() {
	if p.Toggled {
		gui.SetState(gui.STATE_FOCUSED)
	}
	if gui.Button(p.Bounds, gui.IconText(gui.ICON_PLAYER_PAUSE, "")) {
		p.Click()
	}
}

func (p *Pause) Close() {

}
