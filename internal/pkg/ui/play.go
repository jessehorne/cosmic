package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type player interface {
	Play()
}

type Play struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Toggled bool
	Bounds  rl.Rectangle
	player  player
}

func NewPlay(p player) *Play {
	play := &Play{
		W:      30,
		H:      30,
		player: p,
	}
	play.UpdateBounds()
	return play
}

func (p *Play) GetBounds() rl.Rectangle {
	return p.Bounds
}

func (p *Play) SetBounds(r rl.Rectangle) {
	p.Bounds = r
	p.X = int32(p.Bounds.X)
	p.Y = int32(p.Bounds.Y)
	p.W = int32(p.Bounds.Width)
	p.H = int32(p.Bounds.Height)
}

func (p *Play) UpdateBounds() {
	p.Bounds = rl.NewRectangle(float32(p.X), float32(p.Y), float32(p.W), float32(p.H))
}

func (p *Play) Click() {
	p.player.Play()
}

func (p *Play) SetToggle(w bool) {
	p.Toggled = w
}

func (p *Play) Update() {
	p.UpdateBounds()
}

func (p *Play) Draw() {
	if p.Toggled {
		gui.SetState(gui.STATE_FOCUSED)
	}
	if gui.Button(p.Bounds, gui.IconText(gui.ICON_PLAYER_PLAY, "")) {
		p.Click()
	}
}

func (p *Play) Close() {

}
