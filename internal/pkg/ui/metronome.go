package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type toggler interface {
	ToggleMetronome()
}

type Metronome struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Bounds  rl.Rectangle
	Toggled bool
	toggler toggler

	// Ticking decides if at the current moment, the metronome is showing a "tick"
	// a "tick" in this case is if the color of an active metronome is changed
	// from skyblue to white. The DAW will "tick" the metronome once per beat.
	// when ticking is done, the color goes back to skyblue.
	Ticking          bool // if triggered to tick for this beat (this happens once per beat)
	TickingCounter   float32
	TickSoundCounter int

	Sound1 rl.Sound
	Sound2 rl.Sound
}

func NewMetronome(t toggler) *Metronome {
	return &Metronome{
		X:       210,
		Y:       0,
		W:       30,
		H:       30,
		toggler: t,
		Sound1:  rl.LoadSound("data/metronome-1.ogg"),
		Sound2:  rl.LoadSound("data/metronome-2.ogg"),
	}
}

func (m *Metronome) GetBounds() rl.Rectangle {
	return m.Bounds
}

func (m *Metronome) UpdateBounds() {
	m.Bounds = rl.NewRectangle(float32(m.X), float32(m.Y), float32(m.W), float32(m.H))
}

func (m *Metronome) Click() {
	m.Toggled = !m.Toggled
	m.toggler.ToggleMetronome()
}

func (m *Metronome) Update() {
	m.UpdateBounds()

	// if ticking, increment ticking counter and reset to 0 if it goes over .1
	// (a tenth of a second)
	if m.Ticking {
		m.TickingCounter += rl.GetFrameTime()

		if m.TickingCounter > 0.1 {
			m.Ticking = false
			m.TickingCounter = 0
		}
	}

	isMouseHovering := rl.CheckCollisionPointRec(rl.GetMousePosition(), m.Bounds)
	isLeftClick := rl.IsMouseButtonPressed(rl.MouseButtonLeft)

	if isMouseHovering && isLeftClick {
		m.Click()
	}
}

func (m *Metronome) Draw() {
	rl.DrawRectangle(m.X, m.Y, m.W, m.H, rl.Gray)
	if m.Toggled {
		color := rl.SkyBlue
		if m.Ticking {
			color = rl.White
		}
		rl.DrawRectangle(m.X+2, m.Y+2, m.W-4, m.H-4, color)
	}
	gui.Label(
		rl.NewRectangle(float32(m.X+6), float32(m.Y), float32(m.W), float32(m.H)),
		gui.IconText(gui.ICON_CLOCK, ""))
}

func (m *Metronome) Tick() {
	m.Ticking = true

	if m.Toggled {
		if m.TickSoundCounter == 0 {
			rl.PlaySound(m.Sound2)
		} else {
			rl.PlaySound(m.Sound1)
		}
	}

	m.TickSoundCounter += 1
	if m.TickSoundCounter > 3 {
		m.TickSoundCounter = 0
	}
}

func (m *Metronome) Reset() {
	m.TickSoundCounter = 0
}

func (m *Metronome) Close() {
	rl.UnloadSound(m.Sound1)
	rl.UnloadSound(m.Sound2)
}
