package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/music"
)

type Sequencer struct {
	*Core
	TimeSignature music.TimeSignature
}

func NewSequencer(timeSig music.TimeSignature) *Sequencer {
	return &Sequencer{
		Core:          NewCore(rl.NewRectangle(0, 0, 30, 30)),
		TimeSignature: timeSig,
	}
}

func (s *Sequencer) GetCore() *Core {
	return s.Core
}

func (s *Sequencer) Click() {

}

func (s *Sequencer) Update() {

}

func (s *Sequencer) Draw() {
	if gui.Button(s.Bounds, gui.IconText(gui.ICON_PLAYER_PLAY, "")) {
		s.Click()
	}
}

func (s *Sequencer) Close() {

}
