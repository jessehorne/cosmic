package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Core struct {
	Bounds rl.Rectangle
	Origin rl.Rectangle
}

func NewCore(b rl.Rectangle) *Core {
	return &Core{
		Bounds: b,
		Origin: rl.NewRectangle(0, 0, 0, 0),
	}
}

func (c *Core) UnpackInt32() (int32, int32, int32, int32) {
	return int32(c.Bounds.X), int32(c.Bounds.Y),
		int32(c.Bounds.Width), int32(c.Bounds.Height)
}

func (c *Core) GetBounds() rl.Rectangle {
	return c.Bounds
}

func (c *Core) SetBounds(r rl.Rectangle) {
	c.Bounds = r
}

func (c *Core) SetOrigin(r rl.Rectangle) {
	c.Origin = r
}
