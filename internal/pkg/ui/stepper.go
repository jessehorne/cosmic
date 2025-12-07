package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stepper struct {
	*Core
	Widgets map[string]Widget
}

func NewStepper() *Stepper {
	panKnob := NewKnob(5, 5, 20, 20, 100)
	volumeKnob := NewKnob(5, 25, 20, 20, 100)

	return &Stepper{
		Core: NewCore(rl.NewRectangle(0, 0, 500, 100)),
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
	}
}

func (s *Stepper) Draw() {
	for _, widget := range s.Widgets {
		widget.Draw()
	}
}

func (s *Stepper) Close() {

}
