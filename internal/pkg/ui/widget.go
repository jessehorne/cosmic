package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Widget interface {
	GetBounds() rl.Rectangle
	SetBounds(r rl.Rectangle)
	Update()
	Draw()
	Close()
}
