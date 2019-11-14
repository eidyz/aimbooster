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

// Misses variable
var Misses = 0
var targets = []target.Target{}

var startTime = time.Now()

var addTarget = throttle.ThrottleFunc(time.Duration(1000)*time.Millisecond, false, func() {
	targets = append(targets, target.Init())
})

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if Misses >= 3 {
		log.Fatalln("You lasted ", time.Since(startTime), " seconds")
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
				if targets[i].Size <= 0 {
					Misses++
				}
				targets = append(targets[:i], targets[i+1:]...)
				i--
			}
		}
	}

	ebitenutil.DebugPrint(screen, func() string {
		return fmt.Sprintf("%s%d", "Score: ", int(time.Since(startTime).Seconds()))
	}())

	maxTargets := int(time.Since(startTime).Milliseconds()/1000) / 13
	if len(targets) <= maxTargets {
		targets = append(targets, target.Init())
	}
	return nil
}

func main() {
	targets = append(targets, target.Init())
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, w, h, 1, "Go Aimbooster"); err != nil {
		log.Fatal(err)
	}
}
