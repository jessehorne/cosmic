package pkg

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/ui"
)

const (
	minBPM = 1   // minimum beats per minute
	maxBPM = 500 // max beats per minute
)

type (
	DAW struct {
		BPM              int
		Playing          bool
		PlayingSong      bool // playing pattern if false, song if true
		MetronomeToggled bool // if the metronome should play during song

		Metronome     metronome
		PlayButton    button
		PauseButton   button
		StopButton    button
		AddStepButton button

		Widgets []ui.Widget

		PlayTime    float32 // current play time in seconds
		BeatCounter float32
		CurrentBeat int
	}

	metronome interface {
		Tick()
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

func NewDAW() *DAW {
	d := &DAW{
		BPM: 80,
	}

	playButton := ui.NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_PLAY,
		func() {
			d.Play()
		},
	)
	pauseButton := ui.NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_PAUSE,
		func() {
			d.Pause()
		},
	)
	stopButton := ui.NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_PLAYER_STOP,
		func() {
			d.Stop()
		},
	)
	bpmCounter := ui.NewBpm(d)

	which := ui.NewWhich(d)
	m := ui.NewMetronome(d)
	
	addStepButton := ui.NewButton(
		rl.NewRectangle(0, 0, 30, 30),
		gui.ICON_STEP_INTO,
		func() {
			// TODO
		},
	)

	hOrganizer := ui.NewOrganizer()

	var widgets []ui.Widget
	widgets = append(widgets, hOrganizer.AddWidget(playButton))
	widgets = append(widgets, hOrganizer.AddWidget(pauseButton))
	widgets = append(widgets, hOrganizer.AddWidget(stopButton))
	widgets = append(widgets, hOrganizer.AddWidget(bpmCounter))
	widgets = append(widgets, hOrganizer.AddWidget(which))
	widgets = append(widgets, hOrganizer.AddWidget(m))
	widgets = append(widgets, hOrganizer.AddWidget(addStepButton))
	d.Widgets = widgets

	d.Metronome = m
	d.PlayButton = playButton
	d.PauseButton = pauseButton
	d.StopButton = stopButton
	d.AddStepButton = addStepButton

	hOrganizer.SimpleHorizontal()

	return d
}

func (d *DAW) Close() {
	for _, w := range d.Widgets {
		w.Close()
	}
}

func (d *DAW) Update() {
	for _, w := range d.Widgets {
		w.Update()
		w.Draw()

		gui.SetState(gui.STATE_NORMAL)
	}

	if d.Playing {
		if d.PlayTime == 0 {
			d.TickMetronome()
		}

		// increment play time, beat counter and current beat
		d.PlayTime += rl.GetFrameTime()
		d.BeatCounter += rl.GetFrameTime()

		// tick the metronome once per beat
		if d.BeatCounter >= 60.0/float32(d.BPM) {
			d.TickMetronome()
			d.BeatCounter = 0
		}

		beatSize := float32(60 / d.BPM)
		d.CurrentBeat = int(d.PlayTime / beatSize)
	}

	// check if user hit space bar
	if rl.IsKeyPressed(rl.KeySpace) {
		if d.Playing {
			if d.StopButton != nil {
				d.StopButton.Click()
			}
			d.PlayButton.SetToggle(false)
		} else {
			if d.PlayButton != nil {
				d.PlayButton.Click()
			}
		}
	}
}

func (d *DAW) GetBPM() int {
	return d.BPM
}

func (d *DAW) SetBPM(bpm int) {
	if bpm < minBPM || bpm > maxBPM {
		return
	}
	d.BPM = bpm
}

func (d *DAW) Play() {
	d.Playing = true

	if d.PauseButton != nil {
		d.PauseButton.SetToggle(false)
	}

	if d.PlayButton != nil {
		d.PlayButton.SetToggle(true)
	}

	fmt.Println("playing")
}

func (d *DAW) Pause() {
	if d.Playing {
		if d.PauseButton != nil {
			d.PauseButton.SetToggle(true)
		}
		if d.PlayButton != nil {
			d.PlayButton.SetToggle(false)
		}
	}

	d.Playing = false
	fmt.Println("paused")
}

func (d *DAW) Stop() {
	d.Playing = false
	d.ResetTime()
	d.ResetMetronome()

	if d.PauseButton != nil {
		d.PauseButton.SetToggle(false)
	}

	if d.PlayButton != nil {
		d.PlayButton.SetToggle(false)
	}
	fmt.Println("stopped")
}

func (d *DAW) ResetTime() {
	d.PlayTime = 0
	d.BeatCounter = 0
	fmt.Println("reset time...TODO")
}

func (d *DAW) TogglePlayingSong() {
	d.PlayingSong = !d.PlayingSong
	fmt.Println("toggled between pattern and song", d.PlayingSong)
}

func (d *DAW) ToggleMetronome() {
	d.MetronomeToggled = !d.MetronomeToggled
	fmt.Println("toggled metronome: ", d.MetronomeToggled)
}

func (d *DAW) TickMetronome() {
	if d.Metronome == nil {
		return
	}
	d.Metronome.Tick()
}

func (d *DAW) ResetMetronome() {
	if d.Metronome == nil {
		return
	}
	d.Metronome.Reset()
}
