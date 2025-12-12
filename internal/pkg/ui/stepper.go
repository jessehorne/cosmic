package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/daw"
)

type Stepper struct {
	*Core
	Widgets  map[string]Widget
	Knobs    map[string]*Knob
	SButtons []*SButton
	DAW      *daw.DAW
}

func NewStepper(d *daw.DAW) *Stepper {
	panKnob := NewKnob(5, 5, 20, 20, 50, func(int) {})
	volumeKnob := NewKnob(5, 30, 20, 20, 100, func(int) {})

	s := &Stepper{
		Core: NewCore(rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), 100)),
		Widgets: map[string]Widget{
			"pan":    panKnob,
			"volume": volumeKnob,
		},
		Knobs: map[string]*Knob{
			"pan":    panKnob,
			"volume": volumeKnob,
		},
		DAW: d,
	}
	s.PlaceSButtons()
	return s
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
	for _, sButton := range s.SButtons {
		sButton.Update()
		sButton.SetOrigin(s.Origin)
	}

	scroll := rl.GetMouseWheelMove()
	if scroll != 0 {
		if scroll < 0 {
			s.Origin.X -= 40
		} else if scroll > 0 {
			s.Origin.X += 40
			if s.Origin.X > 0 {
				s.Origin.X = 0
			}
		}
	}
}

func (s *Stepper) PlaceSButtons() {
	s.SButtons = []*SButton{}

	divided := 16 / s.DAW.TimeSig.Denominator
	count := divided * s.DAW.TimeSig.Numerator
	current := 0
	which := false
	for i := 0; i < count; i++ {
		if current == divided {
			which = !which
			current = 0
		}

		newSButton := NewSButton(false, i+1, which, divided)
		s.SButtons = append(s.SButtons, newSButton)

		current += 1
	}
}

func (s *Stepper) Draw() {
	for _, widget := range s.Widgets {
		widget.Draw()
	}
	for _, sButton := range s.SButtons {
		sButton.Draw()
	}
}

func (s *Stepper) Close() {

}
