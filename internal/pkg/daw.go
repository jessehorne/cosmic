package pkg

import "fmt"

const (
	minBPM = 1   // minimum beats per minute
	maxBPM = 500 // max beats per minute
)

type DAW struct {
	BPM     int
	Playing bool
}

func NewDAW() *DAW {
	return &DAW{
		BPM: 80,
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
	fmt.Println("stopped")
}

func (d *DAW) ResetTime() {
	fmt.Println("reset time...TODO")
}
