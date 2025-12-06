package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Widget interface {
	GetBounds() rl.Rectangle
	Update()
	Draw()
}
