package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type player struct {
	images     [8]*ebiten.Image
	currentImg int
	curPos     pos
}

func newPlayer(y, x int) *player {
	p := &player{}
	p.loadImages()
	p.curPos = pos{y, x}
	return p
}

func (p *player) loadImages() {
	for i := 0; i < 8; i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.PlayerImages[i]))
		handleError(err)
		p.images[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
}

func (p *player) image() *ebiten.Image {
	return p.images[p.currentImg]
}

func (p *player) draw(screen *ebiten.Image) {
	x := float64(p.curPos.x * stageBlocSize)
	y := float64(p.curPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(p.image(), op)
}
