package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stepper struct {
	*Core
	Widgets map[string]Widget
}

func NewStepper() *Stepper {
	panKnob := NewKnob(5, 5, 20, 20, 50)
	volumeKnob := NewKnob(5, 30, 20, 20, 100)

	return &Stepper{
		Core: NewCore(rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), 100)),
		Widgets: map[string]Widget{
			"pan":    panKnob,
			"volume": volumeKnob,
		},
	}
}

func (s *Stepper) GetCore() *Core {
	return s.Core
}

func (s *Stepper) Click() {

}

func (s *Stepper) Update() {
	for _, widget := range s.Widgets {
		widget.Update()
		widget.SetOrigin(s.Origin)
	}
}

func (s *Stepper) Draw() {
	//rl.PushMatrix()
	//
	//rl.Translatef(s.Core.Origin.X, s.Core.Origin.Y, 0)
	for _, widget := range s.Widgets {
		widget.Draw()
	}
	//
	//rl.PopMatrix()
}

func (s *Stepper) Close() {

}
