package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Organizer struct {
	Widgets []Widget
}

func NewOrganizer() *Organizer {
	return &Organizer{
		Widgets: []Widget{},
	}
}

func (o *Organizer) AddWidget(w Widget) Widget {
	o.Widgets = append(o.Widgets, w)
	return w
}

func (o *Organizer) SimpleHorizontal() {
	var lastX float32 = 0.0
	for _, w := range o.Widgets {
		b := w.GetBounds()
		fmt.Println(lastX)
		w.SetBounds(rl.NewRectangle(lastX, b.Y, b.Width, b.Height))
		lastX += b.X + b.Width
	}
}
