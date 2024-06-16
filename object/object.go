package object

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Position struct {
	X int
	Y int
}

// IsUnder makes comparing Y values more readable, as the y-axis is inverted
func (p Position) IsUnder(y int) bool {
	return p.Y > y
}

type Velocity struct {
	X int
	Y int
}
