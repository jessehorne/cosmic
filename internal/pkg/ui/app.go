package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/color"
	"github.com/jessehorne/cosmic/internal/pkg/daw"
	"github.com/jessehorne/cosmic/internal/pkg/music"
)

const (
	screenWidth  = 900
	screenHeight = 600
)

type (
	App struct {
		DAW *daw.DAW

		Metronome     metronome
		PlayButton    button
		PauseButton   button
		StopButton    button
		AddStepButton button
		VolumeKnob    *Knob

		Widgets  []Widget
		Steppers []*Stepper

		Font *rl.Font
	}

	metronome interface {
		Tick(w bool)
		Reset()
	}

	button interface {
		Click()
		SetToggle(w bool)
	}

	addStep interface {
		Click()
	}
)

func NewApp() *App {
	// init raylib stuff
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowHighdpi)
	rl.InitWindow(screenWidth, screenHeight, "cosmic v0.0.1")

	font := rl.LoadFont("data/spacemono.ttf")
	gui.SetFont(font)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	// gui settings
	gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_NORMAL, int64(rl.ColorToInt(color.Clear)))
	gui.SetStyle(gui.BUTTON, gui.BORDER_COLOR_NORMAL, int64(rl.ColorToInt(color.LightestBlue)))
	//gui.SetStyle(gui.BUTTON, gui.BACKGROUND_COLOR, int64(rl.ColorToInt(color.Clear)))
	gui.SetStyle(gui.BUTTON, gui.TEXT_COLOR_NORMAL, int64(rl.ColorToInt(color.LightestBlue)))
	gui.SetStyle(gui.BUTTON, gui.BORDER_WIDTH, 1)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 22)
	gui.SetStyle(gui.VALUEBOX, gui.BORDER_COLOR_NORMAL, int64(rl.ColorToInt(color.LightestBlue)))
	gui.SetStyle(gui.VALUEBOX, gui.TEXT_COLOR_NORMAL, int64(rl.ColorToInt(color.LightestBlue)))
	gui.SetStyle(gui.LABEL, gui.TEXT_COLOR_NORMAL, int64(rl.ColorToInt(color.LightestBlue)))

	a := &App{
		Font: &font,
	}

	tickMetronome := func(w bool) {
		if a.Metronome == nil {
			return
		}
		a.Metronome.Tick(w)
	}

	resetMetronome := func() {
		if a.Metronome == nil {
			return
		}
		a.Metronome.Reset()
	}

	newDAW := daw.NewDAW(tickMetronome, resetMetronome)
	a.DAW = newDAW

	hOrganizer := NewOrganizer()
	wOrganizer := NewOrganizer()

	playButton := NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_PLAY,
		func() {
			a.DAW.Play()
			if a.PauseButton != nil {
				a.PauseButton.SetToggle(false)
			}

			if a.PlayButton != nil {
				a.PlayButton.SetToggle(true)
			}
		},
	)
	pauseButton := NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_PAUSE,
		func() {
			if a.DAW.Playing {
				if a.PauseButton != nil {
					a.PauseButton.SetToggle(true)
				}
			}
			if a.PlayButton != nil {
				a.PlayButton.SetToggle(false)
			}
			a.DAW.Pause()
		},
	)
	stopButton := NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_STOP,
		func() {
			a.DAW.Stop()
			if a.PauseButton != nil {
				a.PauseButton.SetToggle(false)
			}

			if a.PlayButton != nil {
				a.PlayButton.SetToggle(false)
			}
		},
	)
	bpmCounter := NewBpm(a.DAW)

	which := NewWhich(a.DAW, a.GetFont(), func(w bool) {
		a.DAW.PlayingSong = w // false means pattern, true means song
	})
	m := NewMetronome(a.DAW)

	addStepButton := NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_STEP_INTO,
		func() {
			newStepper := NewStepper(a.DAW)

			startY := float32(50 + len(a.Steppers)*60 + 10)
			newStepper.SetOrigin(rl.NewVector2(0, startY))

			a.Widgets = append(a.Widgets, wOrganizer.AddWidget(newStepper))
			a.Steppers = append(a.Steppers, newStepper)
		},
	)

	timing := NewTiming(a.DAW, a.GetFont(), func(ts *music.TimeSignature) {
		a.DAW.UpdateTimeSignature(ts)
		for _, s := range a.Steppers {
			s.PlaceSButtons()
		}
	})

	volumeKnob := NewKnob(0, 0, 30, 30, 100, func(v int) {
		a.DAW.Volume = v
	})

	var widgets []Widget
	widgets = append(widgets, hOrganizer.AddWidget(playButton))
	widgets = append(widgets, hOrganizer.AddWidget(pauseButton))
	widgets = append(widgets, hOrganizer.AddWidget(stopButton))
	widgets = append(widgets, hOrganizer.AddWidget(bpmCounter))
	widgets = append(widgets, hOrganizer.AddWidget(which))
	widgets = append(widgets, hOrganizer.AddWidget(m))
	widgets = append(widgets, hOrganizer.AddWidget(addStepButton))
	widgets = append(widgets, hOrganizer.AddWidget(timing))
	widgets = append(widgets, hOrganizer.AddWidget(volumeKnob))
	a.Widgets = widgets

	a.Metronome = m
	a.PlayButton = playButton
	a.PauseButton = pauseButton
	a.StopButton = stopButton
	a.AddStepButton = addStepButton

	hOrganizer.SimpleHorizontal()

	return a
}

func (a *App) Run() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 0, G: 0, B: 0, A: 220})

		a.Update()
		for _, w := range a.Widgets {
			w.Update()
			w.Draw()
		}

		rl.EndDrawing()
	}

	a.DAW.Close()

	for _, w := range a.Widgets {
		w.Close()
	}

	rl.UnloadFont(*a.Font)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func (a *App) Update() {
	a.DAW.Update()

	// check if user hit space bar
	if rl.IsKeyPressed(rl.KeySpace) {
		if a.DAW.Playing {
			if a.StopButton != nil {
				a.StopButton.Click()
			}
			a.PlayButton.SetToggle(false)
		} else {
			if a.PlayButton != nil {
				a.PlayButton.Click()
			}
		}
	}
}

func (a *App) GetFont() *rl.Font {
	return a.Font
}
