package player

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/kylesmartin/fetch/object"
	"github.com/kylesmartin/fetch/settings"
)

var (
	leftSprite  *ebiten.Image
	rightSprite *ebiten.Image
	idleSprite  *ebiten.Image
)

const (
	moveSpeed        = 8
	maxFallSpeed     = 20
	deltaFallSpeed   = 12
	jumpMultiplier   = 0.75
	fallMultiplier   = 2.0
	initialJumpSpeed = 15
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)
}

// Player is the main object controlled by users
type Player struct {
	Position   object.Position
	Velocity   object.Velocity
	IsJumpHeld bool
}

// Update implements Object.Update
func (p *Player) Update() {
	// Adjust velocities based on inputs
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Velocity.X = -1 * moveSpeed * settings.Unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Velocity.X = moveSpeed * settings.Unit
	} else if p.Velocity.X > 0 {
		p.Velocity.X -= moveSpeed // No settings.Unit scaling to cause gradual slowdown
	} else if p.Velocity.X < 0 {
		p.Velocity.X += moveSpeed
	}

	// Trigger jump
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.IsJumpHeld {
		p.tryJump()
		p.IsJumpHeld = true
	}
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.IsJumpHeld = false
	}

	// Translate position according to velocity
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

	// Reset position to ground level if below it
	groundLevel := float64(settings.GroundY * settings.Unit)
	if p.Position.Y > groundLevel {
		p.Position.Y = groundLevel
	}

	// Jump modification mechanics:
	// If going up while holding the jump, reduce the rate of velocity change
	// If going up without holding the jump, use the base rate of velocity change
	// If going down, increase the rate of velocity change
	if p.Velocity.Y < maxFallSpeed*settings.Unit {
		if p.Velocity.Y < 0 && p.IsJumpHeld {
			p.Velocity.Y += deltaFallSpeed * jumpMultiplier
		} else if p.Velocity.Y < 0 && !p.IsJumpHeld {
			p.Velocity.Y += deltaFallSpeed
		} else if p.Velocity.Y >= 0 {
			p.Velocity.Y += deltaFallSpeed * fallMultiplier
		}
	}

	// Keep player within screen bounds
	// TODO: Base this on settings.ScreenWidth and settings.ScreenHeight and sprite dimensions
	if p.Position.X < 0 {
		p.Position.X = 0
	}
	if p.Position.X > 14000 {
		p.Position.X = 14000
	}
}

// Draw implements Object.Draw
func (p *Player) Draw(screen *ebiten.Image) {
	s := idleSprite
	switch {
	case p.Velocity.X > 0:
		s = rightSprite
	case p.Velocity.X < 0:
		s = leftSprite
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(p.Position.X/settings.Unit, p.Position.Y/settings.Unit)
	screen.DrawImage(s, op)
}

// tryJump triggers a jump if the player is on the ground
func (p *Player) tryJump() {
	if p.Position.Y == settings.GroundY*settings.Unit {
		p.Velocity.Y = -1 * initialJumpSpeed * settings.Unit
	}
}
