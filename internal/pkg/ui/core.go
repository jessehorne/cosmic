package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Core struct {
	Bounds rl.Rectangle
}

func NewCore(b rl.Rectangle) Core {
	return Core{
		Bounds: b,
	}
}

func (c Core) UnpackInt32() (int32, int32, int32, int32) {
	return int32(c.Bounds.X), int32(c.Bounds.Y),
		int32(c.Bounds.Width), int32(c.Bounds.Height)
}
