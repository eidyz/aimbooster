package target

import (
	"image/color"

	"github.com/eidyz/aimbooster/util"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Target ---
type Target struct {
	X       float64
	Y       float64
	Size    float64
	Speed   int
	Color   color.Color
	Grow    bool
	Clicked bool
}

// Centerize --
func Centerize(xIn, size float64) float64 {
	x := (xIn - (size / 2) + 100)
	return x
}

// New creates new Target
func New(x, y, size float64, speed int) (target Target) {
	target.X = x
	target.Y = y
	target.Size = size
	target.Speed = speed
	target.Color = color.RGBA{100, 200, 200, 0xff}
	target.Grow = true
	return
}

// NewRandom ---
func NewRandom() (target Target) {
	w, h := ebiten.ScreenSizeInFullscreen()
	var size float64 = 1
	x := float64(util.RandInt(0, w-150))
	y := float64(util.RandInt(0, h-150))
	return New(x, y, size, 500)
}

// Draw target
func (target *Target) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, Centerize(target.X, target.Size), Centerize(target.Y, target.Size), target.Size, target.Size, target.Color)
}

// CheckHit checks if target was clicked
func (target *Target) CheckHit(dst *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	if float64(mouseX) >= Centerize(target.X, target.Size) &&
		float64(mouseX) <= Centerize(target.X, target.Size)+target.Size &&
		float64(mouseY) >= Centerize(target.Y, target.Size) &&
		float64(mouseY) <= Centerize(target.Y, target.Size)+target.Size {
		target.Clicked = true
	}
}

// Pulse target
func (target *Target) Pulse() {
	if target.Grow {
		if target.Size >= 100 {
			target.Grow = false
		}
		target.Size++
	} else {
		if target.Size <= 0 {
			target.Grow = true
		}
		target.Size--
	}
}

// Init new random Target
func Init() (target Target) {
	target = NewRandom()
	target.Pulse()
	return
}
