package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Core struct {
	Bounds rl.Rectangle
	Origin rl.Vector2
}

func NewCore(b rl.Rectangle) *Core {
	return &Core{
		Bounds: b,
		Origin: rl.NewVector2(0, 0),
	}
}

func (c *Core) UnpackInt32() (int32, int32, int32, int32) {
	return int32(c.Origin.X + c.Bounds.X), int32(c.Origin.Y + c.Bounds.Y),
		int32(c.Bounds.Width), int32(c.Bounds.Height)
}

func (c *Core) GetAdjustedPosition() rl.Vector2 {
	return rl.NewVector2(c.Origin.X+c.Bounds.X, c.Origin.Y+c.Bounds.Y)
}

func (c *Core) GetAdjustedBounds() rl.Rectangle {
	return rl.NewRectangle(
		c.Origin.X+c.Bounds.X, c.Origin.Y+c.Bounds.Y,
		c.Bounds.Width, c.Bounds.Height)
}

func (c *Core) GetBounds() rl.Rectangle {
	return c.Bounds
}

func (c *Core) SetBounds(r rl.Rectangle) {
	c.Bounds = r
}

func (c *Core) SetOrigin(r rl.Vector2) {
	c.Origin = r
}
