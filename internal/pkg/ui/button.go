package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	*Core
	Toggled  bool
	Callback func()
	Icon     int32
}

func NewButton(b rl.Rectangle, icon int32, callback func()) *Button {
	return &Button{
		Core:     NewCore(b),
		Icon:     icon,
		Callback: callback,
	}
}

func (b *Button) GetCore() *Core {
	return b.Core
}

func (b *Button) Click() {
	b.Callback()
}

func (b *Button) SetToggle(w bool) {
	b.Toggled = w
}

func (b *Button) Update() {

}

func (b *Button) Draw() {
	if b.Toggled {
		gui.SetState(gui.STATE_FOCUSED)
	}
	if gui.Button(b.Bounds, gui.IconText(b.Icon, "")) {
		b.Click()
	}
	gui.SetState(gui.STATE_NORMAL)
}

func (b *Button) Close() {

}
