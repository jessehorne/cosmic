package pkg

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
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

		Metronome  metronome
		PlayButton button
		StopButton button

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
	}
)

func NewDAW() *DAW {
	return &DAW{
		BPM: 80,
	}
}

func (d *DAW) SetMetronome(m metronome) {
	d.Metronome = m
}

func (d *DAW) SetPlayButton(pb button) {
	d.PlayButton = pb
}

func (d *DAW) SetStopButton(sb button) {
	d.StopButton = sb
}

func (d *DAW) Update() {
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
	fmt.Println("playing")
}

func (d *DAW) Pause() {
	d.Playing = false
	fmt.Println("paused")
}

func (d *DAW) Stop() {
	d.Playing = false
	d.ResetTime()
	d.ResetMetronome()
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
