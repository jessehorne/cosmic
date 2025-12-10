package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Organizer struct {
	Widgets  []Widget
	PaddingX int
}

func NewOrganizer() *Organizer {
	return &Organizer{
		Widgets:  []Widget{},
		PaddingX: 5,
	}
}

func (o *Organizer) AddWidget(w Widget) Widget {
	o.Widgets = append(o.Widgets, w)
	return w
}

func (o *Organizer) SimpleHorizontal() {
	var lastX float32 = 0.0
	for _, w := range o.Widgets {
		b := w.GetCore().Bounds
		w.GetCore().SetBounds(rl.NewRectangle(lastX, b.Y, b.Width, b.Height))
		lastX += b.X + b.Width + float32(o.PaddingX)
	}
}
