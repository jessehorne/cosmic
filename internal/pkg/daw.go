package pkg

const (
	minBPM = 1   // minimum beats per minute
	maxBPM = 500 // max beats per minute
)

type DAW struct {
	bpm int
}

func NewDAW() *DAW {
	return &DAW{
		bpm: 80,
	}
}

func (d *DAW) BPM() int {
	return d.bpm
}

func (d *DAW) SetBPM(bpm int) {
	if bpm < minBPM || bpm > maxBPM {
		return
	}
	d.bpm = bpm
}
