package main

import (
	"log"

	"github.com/eidyz/aimbooster/util"
	"image/color"

	"github.com/eidyz/aimbooster/core/target"
	"github.com/hajimehoshi/ebiten"
)

//ResX ---
const ResX = 800

// ResY ---
const ResY = 600

// Score variable
var Score = 0
var targets = []target.Target{}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.Fill(color.White)
	if len(targets) > 0 {
		for i := 0; i < len(targets); i++ {
			targets[i].Draw(screen)
			targets[i].CheckHit(screen)
			targets[i].Pulse()

			if targets[i].Size <= 0 {
				targets = append(targets[:i], targets[i+1:]...)
				i--
			}

			if targets[i].Clicked {
				targets = append(targets[:i], targets[i+1:]...)
				i--
			}
		}
	}

	return nil
}

func main() {
	targets = append(targets, target.Init())
	util.SetInterval(func() {
		targets = append(targets, target.Init())
	}, 500, false)
	if err := ebiten.Run(update, ResX, ResY, 1, "Go Aimbooster"); err != nil {
		log.Fatal(err)
	}
}
