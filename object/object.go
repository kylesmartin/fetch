package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Position struct {
	X float64
	Y float64
}

// IsUnder makes comparing Y values more readable, as the y-axis is inverted
func (p Position) IsUnder(y float64) bool {
	return p.Y > y
}

type Velocity struct {
	X float64
	Y float64
}
