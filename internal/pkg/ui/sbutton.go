package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/color"
)

const (
	sButtonLeftPadding = 10
	sButtonTopPadding  = 10
	sButtonSpacing     = 5
	sButtonW           = 25
	sButtonH           = 40
)

type SButton struct {
	*Core
	Toggled bool
	Beat    int
	Which   bool // if false, one color, if true, the other. just for visual aid
}

func NewSButton(toggled bool, beat int, which bool, divided int) *SButton {
	w, h := sButtonW-divided, sButtonH
	x := sButtonLeftPadding + beat*w + sButtonSpacing*beat
	y := sButtonTopPadding
	return &SButton{
		Core:    NewCore(rl.NewRectangle(float32(x), float32(y), float32(w), float32(h))),
		Toggled: toggled,
		Beat:    beat,
		Which:   which,
	}
}

func (s *SButton) GetCore() *Core {
	return s.Core
}

func (s *SButton) Click() {

}

func (s *SButton) Update() {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), s.GetAdjustedBounds()) {
			s.Toggled = true
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), s.GetAdjustedBounds()) {
			s.Toggled = false
		}
	}
}

func (s *SButton) Draw() {
	x, y, w, h := s.UnpackInt32()

	c := color.DarkBlue
	if s.Which {
		c = color.LightBlue
	}

	if s.Toggled {
		c = rl.White
	}

	rl.DrawRectangle(x, y, w, h, c)
}

func (s *SButton) Close() {

}
