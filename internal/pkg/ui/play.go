package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type player interface {
	Play()
}

type Play struct {
	X       int
	Y       int
	W       int
	H       int
	Toggled bool
	Bounds  rl.Rectangle
	player  player
}

func NewPlay(p player) *Play {
	return &Play{
		X:      0,
		Y:      0,
		W:      30,
		H:      30,
		player: p,
	}
}

func (p *Play) GetBounds() rl.Rectangle {
	return p.Bounds
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
