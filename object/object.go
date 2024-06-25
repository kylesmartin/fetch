package object

import "github.com/hajimehoshi/ebiten/v2"

// Objects are spawned and managed by the game
type Object interface {
	Update()
	Draw(screen *ebiten.Image)
}

// Position tracks the location of an object on the screen
type Position struct {
	X float64
	Y float64
}

// Velocity tracks how an object should move on the screen
type Velocity struct {
	X float64
	Y float64
}
