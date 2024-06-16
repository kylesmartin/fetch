package background

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/kylesmartin/fetch/object"
)

var (
	backgroundImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Background_png))
	if err != nil {
		panic(err)
	}
	backgroundImage = ebiten.NewImageFromImage(img)
}

type Background struct {
	Position object.Position
}

// Update implements Object.Update
func (p *Background) Update() error {
	return nil
}

// Draw implements Object.Draw
func (p *Background) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)
}
