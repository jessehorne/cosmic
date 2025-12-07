package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Knob struct {
	*Core
	Value int // should always be 1-100
}

func NewKnob(x, y, w, h float32, value int32) *Knob {
	return &Knob{
		Core:  NewCore(rl.NewRectangle(x, y, w, h)),
		Value: int(value),
	}
}

func (k *Knob) GetCore() *Core {
	return k.Core
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

func (k *Knob) Click() {

}

func (k *Knob) Update() {

}

func (k *Knob) Draw() {
	x, y, w, h := k.UnpackInt32()
	// draw background square
	rl.DrawRectangle(x, y, w, h, rl.Gray)

	// draw knob circle
	centerX := x + w/2
	centerY := y + y/2
	rl.DrawCircle(centerX, centerY, float32(w-4), rl.White)
}

func (k *Knob) Close() {

}
