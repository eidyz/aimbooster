package main

import (
	"log"

	"image/color"

	"github.com/eidyz/aimbooster/util"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	resX = 800
	resY = 600
)

// FPS in MS
const FPS = 144

var cyan = color.RGBA{100, 200, 200, 0xff}

// Target ---
type Target struct {
	x      float64
	y      float64
	width  float64
	heigth float64
}

// NewTarget ---
func NewTarget(x, y, width, height float64) (target Target) {
	target.x = x
	target.y = y
	target.width = width
	target.heigth = height
	return
}

// Pulse target
func (target *Target) Pulse() {
	grow := true
	util.SetInterval(func() {
		if grow {
			if target.width >= 300 || target.heigth >= 300 {
				grow = false
			}
			target.width++
			target.heigth++
		} else {
			if target.width <= 0 || target.heigth <= 0 {
				grow = true
			}
			target.width--
			target.heigth--
		}

	}, (1000 / FPS), false)
}

var target = NewTarget((resX / 2), (resY / 2), 100, 100)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.Fill(color.White)
	log.Print(target)
	ebitenutil.DrawRect(screen, target.x-(target.width/2), target.y-(target.heigth/2), target.width, target.heigth, cyan)
	return nil
}

func main() {
	target.Pulse()
	if err := ebiten.Run(update, resX, resY, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
