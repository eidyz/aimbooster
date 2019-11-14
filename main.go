package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	throttle "github.com/boz/go-throttle"
	"github.com/eidyz/aimbooster/core/target"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

//ResX ---
const ResX = 800

// ResY ---
const ResY = 600

// Score variable
var Score = 0
var targets = []target.Target{}

var addTarget = throttle.ThrottleFunc(time.Duration(500)*time.Millisecond, false, func() {
	targets = append(targets, target.Init())
})

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	err := screen.Fill(color.White)
	if err != nil {
		log.Panic(err)
	}

	if len(targets) > 0 {
		for i := 0; i < len(targets); i++ {
			targets[i].Draw(screen)
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				targets[i].CheckHit(screen)
			}
			targets[i].Pulse()

			if targets[i].Size <= 0 || targets[i].Clicked {
				if targets[i].Clicked {
					Score++
				} else if targets[i].Size <= 0 {
					Score--
				}
				targets = append(targets[:i], targets[i+1:]...)
				i--
			}
		}
	}
	ebitenutil.DebugPrint(screen, func() string {
		return fmt.Sprintf("%s%d", "Your Score: ", Score)
	}())

	addTarget.Trigger()

	return nil
}

func main() {
	targets = append(targets, target.Init())
	if err := ebiten.Run(update, ResX, ResY, 1, "Go Aimbooster"); err != nil {
		log.Fatal(err)
	}
}
