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
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kylesmartin/fetch/background"
	"github.com/kylesmartin/fetch/object"
	"github.com/kylesmartin/fetch/player"
	"github.com/kylesmartin/fetch/settings"
)

type Game struct {
	objects     []object.Object
	initialized bool
}

func (g *Game) Init() {
	background := &background.Background{}
	player := &player.Player{
		Position: object.Position{
			X: 50 * settings.Unit,
			Y: settings.GroundY * settings.Unit,
		},
	}

	// Objects are stored in rendering order
	g.objects = []object.Object{background, player}

	g.initialized = true
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	if !g.initialized {
		g.Init()
	}

	for _, object := range g.objects {
		err := object.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	for _, object := range g.objects {
		object.Draw(screen)
	}

	// Log instructions and ticks per second for debugging (TODO: Remove)
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
