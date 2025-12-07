package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type switcher interface {
	TogglePlayingSong()
	Stop()
}

type Which struct {
	X        int32
	Y        int32
	W        int32
	H        int32
	Bounds   rl.Rectangle
	Toggled  bool
	switcher switcher
}

func NewWhich(s switcher) *Which {
	w := &Which{
		W:        60,
		H:        30,
		switcher: s,
	}
	w.UpdateBounds()
	return w
}

func (w *Which) GetBounds() rl.Rectangle {
	return w.Bounds
}

func (w *Which) SetBounds(r rl.Rectangle) {
	w.Bounds = r
	w.X = int32(w.Bounds.X)
	w.Y = int32(w.Bounds.Y)
	w.W = int32(w.Bounds.Width)
	w.H = int32(w.Bounds.Height)
}

func (w *Which) UpdateBounds() {
	w.Bounds = rl.NewRectangle(float32(w.X), float32(w.Y), float32(w.W), float32(w.H))
}

func (w *Which) Click() {
	fmt.Println("toggled")
	w.Toggled = !w.Toggled
	w.switcher.TogglePlayingSong()
	w.switcher.Stop()
}

func (w *Which) Update() {
	w.UpdateBounds()

	isMouseOver := rl.CheckCollisionPointRec(rl.GetMousePosition(), w.Bounds)
	isLeftClick := rl.IsMouseButtonPressed(rl.MouseLeftButton)

	if isMouseOver && isLeftClick {
		w.Click()
	}
}

func (w *Which) Draw() {
	// first circle
	rad := int32(w.H / 2)
	centerY := int32(w.Y) + rad
	startX := int32(w.X) + rad
	rl.DrawCircle(startX, centerY, float32(rad), rl.Gray)

	// second circle
	startX2 := int32(w.X+w.W) - rad
	rl.DrawCircle(startX2, centerY, float32(rad), rl.Gray)

	// draw background rectangle
	rectWidth := startX2 - startX
	rl.DrawRectangle(startX, int32(w.Y), rectWidth, int32(w.H), rl.Gray)

	if w.Toggled {
		// draw toggle button on right side with background rect showing toggled
		rl.DrawCircle(startX+2, centerY, float32(rad-2), rl.SkyBlue)

		rl.DrawRectangle(startX+2, int32(w.Y+2), rectWidth-4, int32(w.H-4),
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
