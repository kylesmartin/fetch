// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/kylesmartin/fetch/object"
	"github.com/kylesmartin/fetch/player"
	"github.com/kylesmartin/fetch/settings"
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

type Game struct {
	player *player.Player
}

func (g *Game) Update() error {
	if g.player == nil {
		g.player = &player.Player{
			Position: object.Position{
				X: 50 * settings.Unit,
				Y: settings.GroundY * settings.Unit,
			},
		}
	}

	return nil
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	//

	// Draws Background Image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)

	// Draws the Gopher.
	g.player.Draw(screen)

	// Show the message.
	msg := fmt.Sprintf("TPS: %0.2f\nPress the space key to jump.", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.ScreenWidth, settings.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(settings.ScreenWidth, settings.ScreenHeight)
	ebiten.SetWindowTitle("Fetch!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
