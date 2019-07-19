package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type fruitManager struct {
	fruits  [3]*ebiten.Image
	x, y    float64
	count   int
	curr    int
	alpha   float64
	show    bool
	visible bool
}

func newFruitManager(x, y float64) *fruitManager {
	fm := &fruitManager{
		x:     x,
		y:     y,
		alpha: 0,
		count: -400,
		show:  true,
	}
	fm.loadImages()
	return fm
}

func (fm *fruitManager) loadImages() {
	for i := 0; i < len(fm.fruits); i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.FruitImages[i]))
		handleError(err)
		fm.fruits[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
}

func (fm *fruitManager) update() {
	if fm.show {
		fm.count++
		if fm.count >= 70 {
			fm.alpha += 0.01
			if fm.alpha > 1 {
				fm.alpha = 1
			}
		}
	} else {
		fm.count--
		fm.alpha -= 0.01
		if fm.alpha < 0 {
			fm.alpha = 0
		}
	}

	if fm.alpha >= 0.1 {
		fm.visible = true
	} else {
		fm.visible = false
	}

	if fm.count == 400 {
		fm.show = false
	} else if fm.count <= -500 && !fm.show {
		fm.show = true
		fm.curr = (fm.curr + 1) % len(fm.fruits)
	}
}

func (fm *fruitManager) draw(screen *ebiten.Image) {
	fm.update()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(fm.x, fm.y)
	op.ColorM.Scale(1, 1, 1, fm.alpha)
	screen.DrawImage(fm.fruits[fm.curr], op)
}
