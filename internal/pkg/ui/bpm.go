package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type bpmSetter interface {
	SetBPM(v int)
}

const (
	minBPM   = 1   // minimum beats per minute
	maxBPM   = 500 // max beats per minute
	bpmInc   = 1   // amount you can normally increment or decrement the BPM counter
	bpmMulti = 5   // amount to multiple bpmInc when user is holding shift while scrolling
)

type Bpm struct {
	X         int
	Y         int
	W         int
	H         int
	Value     int32
	Bounds    rl.Rectangle
	bpmSetter bpmSetter
}

func NewBpm(bpm bpmSetter) *Bpm {
	b := &Bpm{
		X:         90,
		Y:         0,
		W:         60,
		H:         30,
		bpmSetter: bpm,
		Value:     80,
	}
	b.UpdateBounds()
	return b
}

func (b *Bpm) UpdateBounds() {
	b.Bounds = rl.NewRectangle(float32(b.X), float32(b.Y), float32(b.W), float32(b.H))
}

func (b *Bpm) GetBounds() rl.Rectangle {
	return b.Bounds
}

func (b *Bpm) Scroll(direction int, isHoldingShift bool) {
	b.IncrementBPM(direction > 0, isHoldingShift)
}

// IncrementBPM will change the widgets BPM value and the DAWs BPM. If up is true, it will increment up, if not, then
// down. If multi is true, it will multiply the increment by bpmMulti. For example, normally the value is 1 but if multi
// is true then it will be 5.
func (b *Bpm) IncrementBPM(up, multi bool) {
	amt := bpmInc
	if multi {
		amt = amt * bpmMulti
	}
	vel := -1
	if up {
		vel = 1
	}
	oldValue := b.Value
	b.Value = b.Value + int32(amt*vel)
	if b.Value < minBPM || b.Value > maxBPM {
		b.Value = oldValue
	}
	b.bpmSetter.SetBPM(int(b.Value))
}

func (b *Bpm) Update() {
	// check if hovering widget
	mouseHovering := rl.CheckCollisionPointRec(rl.GetMousePosition(), b.GetBounds())

	if mouseHovering {
		// check if scrolling
		scroll := rl.GetMouseWheelMove()
		if scroll != 0 {
			b.Scroll(int(scroll), rl.IsKeyDown(rl.KeyLeftShift))
		}
	}

	b.UpdateBounds()
}

func (b *Bpm) Draw() {
	gui.ValueBox(
		b.Bounds,
		"",
		&b.Value,
		0,
		500,
		false,
	)
}
