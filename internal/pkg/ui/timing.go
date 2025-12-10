package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/daw"
	"github.com/jessehorne/cosmic/internal/pkg/music"
)

type Timing struct {
	*Core
	TimeSig *music.TimeSignature
	daw     *daw.DAW
	font    *rl.Font
}

func NewTiming(d *daw.DAW, f *rl.Font) *Timing {
	return &Timing{
		Core:    NewCore(rl.NewRectangle(0, 0, 60, 30)),
		TimeSig: music.NewTimeSignature(),
		daw:     d,
		font:    f,
	}
}

func (t *Timing) GetCore() *Core {
	return t.Core
}

func (t *Timing) Click() {

}

func (t *Timing) Update() {
	// check which square mouse is in
	adjBounds := t.GetAdjustedBounds()
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), adjBounds) {
		// check if scrolling, if so increment or decrement each side accordingly
		scroll := rl.GetMouseWheelMove()
		if scroll != 0 {
			if rl.GetMousePosition().X < adjBounds.X+(adjBounds.Width/2) {
				// affecting numerator
				t.TimeSig.IncrementNumerator(int(scroll), 1)
			} else {
				// affecting denominator
				t.TimeSig.IncrementDenominator(int(scroll), 1)
			}
		}
	}
}

func (t *Timing) Draw() {
	a := t.GetAdjustedBounds()
	rl.DrawRectangle(int32(a.X), int32(a.Y), int32(a.Width), int32(a.Height), rl.Gray)

	textSize := rl.MeasureTextEx(*t.font, t.TimeSig.String(), 18, 4)
	centerX := a.X + (a.Width / 2)
	centerY := a.Y + (a.Height / 2)

	posX := centerX - (textSize.X / 2)
	posY := centerY - (textSize.Y / 2)

	rl.DrawTextEx(*t.font, t.TimeSig.String(),
		rl.NewVector2(posX, posY), 18, 4, rl.Black)
}

func (t *Timing) Close() {

}
