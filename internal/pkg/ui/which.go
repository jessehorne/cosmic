package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type switcher interface {
	TogglePlayingSong()
	Stop()
}

type Which struct {
	Core
	Toggled  bool
	switcher switcher
}

func NewWhich(s switcher) *Which {
	return &Which{
		Core:     NewCore(rl.NewRectangle(0, 0, 60, 30)),
		switcher: s,
	}
}

func (w *Which) GetBounds() rl.Rectangle {
	return w.Core.Bounds
}

func (w *Which) SetBounds(r rl.Rectangle) {
	w.Core.Bounds = r
}

func (w *Which) Click() {
	w.Toggled = !w.Toggled
	w.switcher.TogglePlayingSong()
	w.switcher.Stop()
}

func (w *Which) Update() {
	isMouseOver := rl.CheckCollisionPointRec(rl.GetMousePosition(), w.Bounds)
	isLeftClick := rl.IsMouseButtonPressed(rl.MouseLeftButton)

	if isMouseOver && isLeftClick {
		w.Click()
	}
}

func (w *Which) Draw() {
	x, y, width, h := w.Core.UnpackInt32()

	// first circle
	rad := h / 2
	centerY := y + rad
	startX := x + rad
	rl.DrawCircle(startX, centerY, float32(rad), rl.Gray)

	// second circle
	startX2 := x + width - rad
	rl.DrawCircle(startX2, centerY, float32(rad), rl.Gray)

	// draw background rectangle
	rectWidth := startX2 - startX
	rl.DrawRectangle(startX, y, rectWidth, h, rl.Gray)

	if w.Toggled {
		// draw toggle button on right side with background rect showing toggled
		rl.DrawCircle(startX+2, centerY, float32(rad-2), rl.SkyBlue)

		rl.DrawRectangle(startX+2, y+2, rectWidth-4, h-4,
			rl.SkyBlue)

		rl.DrawCircle(startX2-2, centerY, float32(rad-2), rl.White)

		rl.DrawText("S", startX2-5, centerY-5, 14, rl.Black)
	} else {
		// draw toggle button on left side
		rl.DrawCircle(startX+2, centerY, float32(rad-2), rl.White)
		rl.DrawText("P", startX-2, centerY-5, 14, rl.Black)
	}
}

func (w *Which) Close() {

}
