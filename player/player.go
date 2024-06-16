package player

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	Position object.Position
	Velocity object.Velocity
}

// Update implements Object.Update
func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Velocity.X = -4 * settings.Unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Velocity.X = 4 * settings.Unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.tryJump()
	}

	// Translate position according to velocity
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

	groundLevel := settings.GroundY * settings.Unit

	// Reset position to ground level if below it
	if p.Position.IsUnder(groundLevel) {
		p.Position.Y = groundLevel
	}

	// Set horizontal velocity
	if p.Velocity.X > 0 {
		p.Velocity.X -= 4
	} else if p.Velocity.X < 0 {
		p.Velocity.X += 4
	}

	// Add gravity effect
	// TODO: Better understand what this is doing
	if p.Velocity.Y < 20*settings.Unit {
		p.Velocity.Y += 8
	}

	// Keep player within screen bounds
	// TODO: Base this on settings.ScreenWidth and settings.ScreenHeight
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
	op.GeoM.Scale(0.5, 0.5) // Scales the image down by 50%
	op.GeoM.Translate(float64(p.Position.X)/settings.Unit, float64(p.Position.Y)/settings.Unit)
	screen.DrawImage(s, op)
}

func (p *Player) tryJump() {
	if p.Position.Y == settings.GroundY*settings.Unit {
		p.Velocity.Y = -10 * settings.Unit
	}
}
