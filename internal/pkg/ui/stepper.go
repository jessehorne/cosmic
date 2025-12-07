package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stepper struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Bounds  rl.Rectangle
	Widgets map[string]Widget
}

func NewStepper() *Stepper {
	panKnob := NewKnob(5, 5, 20, 20, 100)
	volumeKnob := NewKnob(5, 25, 20, 20, 100)

	s := &Stepper{
		X: 0,
		Y: 0,
		W: 30,
		H: 30,
		Widgets: map[string]Widget{
			"pan":    panKnob,
			"volume": volumeKnob,
		},
	}
	s.UpdateBounds()
	return s
}

func (s *Stepper) GetBounds() rl.Rectangle {
	return s.Bounds
}

func (s *Stepper) SetBounds(r rl.Rectangle) {
	s.Bounds = r
	s.X = int32(s.Bounds.X)
	s.Y = int32(s.Bounds.Y)
	s.W = int32(s.Bounds.Width)
	s.H = int32(s.Bounds.Height)
}

func (s *Stepper) UpdateBounds() {
	s.Bounds = rl.NewRectangle(float32(s.X), float32(s.Y), float32(s.W), float32(s.H))
}

func (s *Stepper) Click() {

}

func (s *Stepper) Update() {
	s.UpdateBounds()
	for _, widget := range s.Widgets {
		widget.Update()
	}
}

func (s *Stepper) Draw() {
	for _, widget := range s.Widgets {
		widget.Draw()
	}
}

func (s *Stepper) Close() {

}
