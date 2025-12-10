package daw

import (
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

		PlayTime    float32 // current play time in seconds
		BeatCounter float32
		CurrentBeat int

		tickMetronomeCallback  func()
		resetMetronomeCallback func()
	}
)

func NewDAW(tmCallback func(), rmCallback func()) *DAW {
	d := &DAW{
		BPM:                    80,
		tickMetronomeCallback:  tmCallback,
		resetMetronomeCallback: rmCallback,
	}

	return d
}

func (d *DAW) Close() {

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
}

func (d *DAW) Pause() {
	d.Playing = false
}

func (d *DAW) Stop() {
	d.Playing = false
	d.ResetTime()
	d.resetMetronomeCallback()
}

func (d *DAW) ResetTime() {
	d.PlayTime = 0
	d.BeatCounter = 0
}

func (d *DAW) TogglePlayingSong() {
	d.PlayingSong = !d.PlayingSong
}

func (d *DAW) ToggleMetronome() {
	d.MetronomeToggled = !d.MetronomeToggled
}

func (d *DAW) TickMetronome() {
	d.tickMetronomeCallback()
}
