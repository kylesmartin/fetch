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

func init() {
	// Preload images
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

type Player struct {
	Position   object.Position
	Velocity   object.Velocity
	IsJumpHeld bool
}

// Update implements Object.Update
func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Velocity.X = -8 * settings.Unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Velocity.X = 8 * settings.Unit
	}
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

	groundLevel := float64(settings.GroundY * settings.Unit)

	// Reset position to ground level if below it
	if p.Position.IsUnder(groundLevel) {
		p.Position.Y = groundLevel
	}

	// Set horizontal velocity
	if p.Velocity.X > 0 {
		p.Velocity.X -= 8
	} else if p.Velocity.X < 0 {
		p.Velocity.X += 8
	}

	// Jump mechanics:
	// - If going up while holding the jump, reduce the rate of velocity change
	// - If going up without holding the jump, use the base rate of velocity change
	// - If going down, increase the rate of velocity change
	maxDownwardVelocity := float64(20 * settings.Unit)
	baseVelocityChange := float64(12)
	upwardsMultiplier := 0.75
	downwardsMultiplier := 2.0
	if p.Velocity.Y < maxDownwardVelocity {
		if p.Velocity.Y < 0 && p.IsJumpHeld {
			p.Velocity.Y += baseVelocityChange * upwardsMultiplier
		} else if p.Velocity.Y < 0 && !p.IsJumpHeld {
			p.Velocity.Y += baseVelocityChange
		} else if p.Velocity.Y >= 0 {
			p.Velocity.Y += baseVelocityChange * downwardsMultiplier
		}
	}

	// Keep player within screen bounds
	// TODO: Base this on settings.ScreenWidth and settings.ScreenHeight
	if p.Position.X < 0 {
		p.Position.X = 0
	}
	if p.Position.X > 14000 {
		p.Position.X = 14000
	}

	return nil
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
	op.GeoM.Scale(0.5, 0.5) // Scales the image down by 50%
	op.GeoM.Translate(float64(p.Position.X)/settings.Unit, float64(p.Position.Y)/settings.Unit)
	screen.DrawImage(s, op)
}

func (p *Player) tryJump() {
	if p.Position.Y == settings.GroundY*settings.Unit {
		p.Velocity.Y = -15 * settings.Unit
	}
}
