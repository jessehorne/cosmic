package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	minValue = 1
	maxValue = 100
	slideMod = 0.1 // how much to increment knob value per pixel of mouse movement
)

type Knob struct {
	*Core
	Value         int // should always be 1-100
	DefaultValue  int
	IsDown        bool       // if the knob is pressed down
	ValueWhenDown int        // the value when the mouse was clicked down
	StartMousePos rl.Vector2 // start of mouse position when user clicked knob
}

func NewKnob(x, y, w, h float32, value int32) *Knob {
	return &Knob{
		Core:         NewCore(rl.NewRectangle(x, y, w, h)),
		Value:        int(value),
		DefaultValue: int(value),
	}
}

func (k *Knob) GetCore() *Core {
	return k.Core
}

func (k *Knob) SetOrigin(o rl.Vector2) {
	k.Origin = o
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
	bounds := rl.NewRectangle(
		k.Origin.X+k.Bounds.X,
		k.Origin.Y+k.Bounds.Y,
		k.Bounds.Width,
		k.Bounds.Height)
	isMouseHovering := rl.CheckCollisionPointRec(rl.GetMousePosition(), bounds)
	isLeftClick := rl.IsMouseButtonPressed(rl.MouseButtonLeft)
	isMouseDown := rl.IsMouseButtonDown(rl.MouseButtonLeft)

	if isMouseHovering && rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		k.Value = k.DefaultValue
	}

	if isMouseHovering && isLeftClick {
		k.IsDown = true
		k.StartMousePos = rl.GetMousePosition()
		k.ValueWhenDown = k.Value
	}

	if k.IsDown && !isMouseDown {
		k.IsDown = false
	}

	if k.IsDown {
		currentX := rl.GetMousePosition().X
		startX := k.StartMousePos.X
		diff := (currentX - startX) * slideMod

		d := (225 + 45) * (diff / 100)
		k.Value = k.ValueWhenDown + int(d)
		if k.Value < minValue {
			k.Value = minValue
		} else if k.Value > maxValue {
			k.Value = maxValue
		}
	}
}

func (k *Knob) Draw() {
	x, y, w, h := k.UnpackInt32()
	// draw background square
	rl.DrawRectangle(x, y, w, h, rl.Gray)

	// draw knob circle
	centerX := x + (w / 2)
	centerY := y + (h / 2)
	rl.DrawCircle(centerX, centerY, float32(w-12), rl.White)

	// draw rotated rectangle to show value
	rec := rl.NewRectangle(float32(centerX), float32(centerY), 10, 3)
	origin := rl.NewVector2(0, 1.5)
	var currentRot float32 = (225+45)*(float32(k.Value)/100) + 135
	rl.DrawRectanglePro(rec, origin, currentRot, rl.Blue)
}

func (k *Knob) Close() {

}
