package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AddStep struct {
	X      int32
	Y      int32
	W      int32
	H      int32
	Bounds rl.Rectangle
}

func NewAddStep() *AddStep {
	as := &AddStep{
		W: 30,
		H: 30,
	}
	as.UpdateBounds()
	return as
}

func (as *AddStep) GetBounds() rl.Rectangle {
	return as.Bounds
}

func (as *AddStep) SetBounds(r rl.Rectangle) {
	as.Bounds = r
	as.X = int32(as.Bounds.X)
	as.Y = int32(as.Bounds.Y)
	as.W = int32(as.Bounds.Width)
	as.H = int32(as.Bounds.Height)
}

func (as *AddStep) UpdateBounds() {
	as.Bounds = rl.NewRectangle(float32(as.X), float32(as.Y), float32(as.W), float32(as.H))
}

func (as *AddStep) Click() {

}

func (as *AddStep) Update() {
	as.UpdateBounds()
}

func (as *AddStep) Draw() {
	if gui.Button(as.Bounds, gui.IconText(gui.ICON_CROSS, "")) {
		as.Click()
	}
}

func (as *AddStep) SetToggle(w bool) {

}

func (as *AddStep) Close() {

}
