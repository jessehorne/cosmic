package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Widget interface {
	GetCore() *Core
	SetOrigin(o rl.Vector2)
	Update()
	Draw()
	Close()
}
