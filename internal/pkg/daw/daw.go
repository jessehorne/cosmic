package daw

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg/music"
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

		Volume int

		TimeSig *music.TimeSignature

		tickMetronomeCallback  func(w bool)
		resetMetronomeCallback func()
	}
)

func NewDAW(tmCallback func(w bool), rmCallback func()) *DAW {
	d := &DAW{
		BPM:                    80,
		TimeSig:                music.NewTimeSignature(),
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
			d.Tick(true)
		}

		beatSize := 60 / float32(d.BPM)
		d.CurrentBeat = int(d.PlayTime / beatSize)

		if d.CurrentBeat >= d.TimeSig.Numerator {
			d.PlayTime = 0
			d.BeatCounter = 0
			d.CurrentBeat = 0
			return
		}

		// tick the metronome once per beat
		if d.BeatCounter >= 60.0/float32(d.BPM) {
			d.Tick(false)
			d.BeatCounter = 0
		}

		// increment play time, beat counter and current beat
		d.PlayTime += rl.GetFrameTime()
		d.BeatCounter += rl.GetFrameTime()
	}
}

func (d *DAW) UpdateTimeSignature(ts *music.TimeSignature) {
	d.TimeSig = ts
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

func (d *DAW) Tick(w bool) {
	d.TickMetronome(w)
}

func (d *DAW) TickMetronome(w bool) {
	d.tickMetronomeCallback(w)
}
