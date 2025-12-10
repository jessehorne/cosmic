package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/color"
	"github.com/jessehorne/cosmic/internal/pkg/daw"
)

type Which struct {
	*Core
	Toggled bool // if toggled is true, Song, if false, Pattern
	daw     *daw.DAW
	font    *rl.Font
}

func NewWhich(d *daw.DAW, f *rl.Font) *Which {
	return &Which{
		Core: NewCore(rl.NewRectangle(0, 0, 60, 30)),
		daw:  d,
		font: f,
	}
}

func (w *Which) GetCore() *Core {
	return w.Core
}

func (w *Which) Click() {
	w.Toggled = !w.Toggled
	w.daw.TogglePlayingSong()
	w.daw.Stop()
}

func (w *Which) Update() {
	isMouseOver := rl.CheckCollisionPointRec(rl.GetMousePosition(), w.Bounds)
	isLeftClick := rl.IsMouseButtonPressed(rl.MouseLeftButton)

	if isMouseOver && isLeftClick {
		w.Click()
	}
}

func (w *Which) Draw() {
	x, y, width, height := w.Core.UnpackInt32()
	pad := int32(3)
	buttonSize := height - pad*2

	// draw background rectangle
	rl.DrawRectangleLinesEx(w.Core.GetAdjustedBounds(), 1, color.LightestBlue)

	if w.Toggled {
		rl.DrawRectangle(x+pad, y+pad, width-pad*2, height-pad*2, color.DarkBlue)
		rl.DrawRectangle(x+width-buttonSize-pad, y+height-buttonSize-pad, buttonSize, buttonSize, rl.White)
		rl.DrawTextEx(*w.font, "S", rl.NewVector2(float32(x+width-buttonSize-pad)+7, float32(y+height-buttonSize-pad)), 24, 1, rl.Black)
	} else {
		// draw toggle button on left side
		rl.DrawRectangle(x+pad, y+pad, buttonSize, buttonSize, rl.White)
		rl.DrawTextEx(*w.font, "P", rl.NewVector2(float32(x+pad)+7, float32(y+pad)), 24, 1, rl.Black)
	}
}

func (w *Which) Close() {

}
