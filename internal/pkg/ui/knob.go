package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Knob struct {
	X      int32
	Y      int32
	W      int32
	H      int32
	Value  int // should always be 1-100
	Bounds rl.Rectangle
}

func NewKnob(x, y, w, h, value int32) *Knob {
	k := &Knob{
		W:     w,
		H:     h,
		Value: int(value),
	}
	k.UpdateBounds()
	return k
}

func (k *Knob) GetValue() int {
	return k.Value
}

func (k *Knob) SetValue(i int) {
	if i < 1 || i > 100 {
		return
	}
	k.Value = i
}

func (k *Knob) GetBounds() rl.Rectangle {
	return k.Bounds
}

func (k *Knob) SetBounds(r rl.Rectangle) {
	k.Bounds = r
	k.X = int32(k.Bounds.X)
	k.Y = int32(k.Bounds.Y)
	k.W = int32(k.Bounds.Width)
	k.H = int32(k.Bounds.Height)
}

func (k *Knob) UpdateBounds() {
	k.Bounds = rl.NewRectangle(float32(k.X), float32(k.Y), float32(k.W), float32(k.H))
}

func (k *Knob) Click() {

}

func (k *Knob) Update() {
	k.UpdateBounds()
}

func (k *Knob) Draw() {
	fmt.Println("test")
	// draw background square
	rl.DrawRectangle(k.X, k.Y, k.W, k.H, rl.Gray)

	// draw knob circle
	centerX := k.X + k.H/2
	centerY := k.Y + k.H/2
	rl.DrawCircle(centerX, centerY, float32(k.W-4), rl.White)
}

func (k *Knob) Close() {

}
