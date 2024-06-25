package tennis

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kylesmartin/fetch/object"
	"github.com/kylesmartin/fetch/resources"
	"github.com/kylesmartin/fetch/settings"
)

var (
	ball *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.TennisBall_png))
	if err != nil {
		panic(err)
	}
	ball = ebiten.NewImageFromImage(img)
}

// Ball is thrown for the player to catch
type Ball struct {
	Position object.Position
	Velocity object.Velocity
}

// Throw initiates the flight path of the ball with a random start position and velocity
func (t *Ball) Throw() {
	// TODO: Randomize parameters
	t.Position.X = 50 * settings.Unit
	t.Position.Y = settings.GroundY * settings.Unit
	t.Velocity.X = 5 * settings.Unit
	t.Velocity.Y = -10 * settings.Unit
}

// Update implements Object.Update
func (t *Ball) Update() {
	/*
		// Hold position if at ground level
		groundLevel := float64(settings.GroundY * settings.Unit)
		fmt.Println(t.Position.Y)
		if t.Position.Y >= groundLevel {
			t.Position.Y = groundLevel
			return
		}

		// Translate position according to velocity
		t.Position.X += t.Velocity.X
		t.Position.Y += t.Velocity.Y

		// Gravity should act upon the ball
		maxDownwardVelocity := float64(20 * settings.Unit)
		t.Velocity.Y += 5
		if t.Velocity.Y > maxDownwardVelocity {
			t.Velocity.Y = maxDownwardVelocity
		}
	*/
}

// Draw implements Object.Draw
func (p *Ball) Draw(screen *ebiten.Image) {
	s := ball
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(p.Position.X/settings.Unit, p.Position.Y/settings.Unit)
	screen.DrawImage(s, op)
}
