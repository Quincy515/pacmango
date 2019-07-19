package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type player struct {
	images                 [8]*ebiten.Image
	currentImg             int
	prvPos, curPos, nxtPos pos
	speed                  int
	stepsLength            pos
	steps                  int
	dir                    input // direction
}

func newPlayer(y, x int) *player {
	p := &player{}
	p.loadImages()
	p.curPos = pos{y, x}
	p.prvPos = pos{y, x}
	p.nxtPos = pos{y, x}
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
	x := float64(p.curPos.x*stageBlocSize + p.stepsLength.x) // new
	y := float64(p.curPos.y*stageBlocSize + p.stepsLength.y) // new
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(p.image(), op)
}

func (p *player) move(m [][]elem, dir input) {
	// not moving and no direction
	if !p.isMoving() && dir == 0 {
		return
	}
	// new direction
	if !p.isMoving() && dir != 0 {
		if !canMove(m, addPosDir(dir, p.curPos)) {
			return
		}
		p.updateDir(dir)
	}
	// adjust the speed
	if p.steps <= 1 || p.steps >= 6 {
		p.speed = 4
	} else {
		p.speed = 5
	}
	// move (update the coordinates)
	switch p.dir {
	case up:
		p.stepsLength.y -= p.speed
	case right:
		p.stepsLength.x += p.speed
	case down:
		p.stepsLength.y += p.speed
	case left:
		p.stepsLength.x -= p.speed
	}

	if p.steps > 5 {
		p.updateImage(false)
	} else {
		p.updateImage(true)
	}

	p.steps++

	if p.steps >= 7 {
		p.endMove()
	}
}

func (p *player) isMoving() bool {
	if p.steps > 0 {
		return true
	}
	return false
}

func (p *player) updateDir(d input) {
	p.stepsLength = pos{0, 0}
	p.dir = d
	p.nxtPos = addPosDir(d, p.curPos)
	p.prvPos = p.curPos
}

func (p *player) endMove() {
	p.curPos = p.nxtPos
	p.stepsLength = pos{0, 0}
	p.steps = 0
}

func (p *player) updateImage(openMouth bool) {
	switch p.dir {
	case up:
		if openMouth {
			p.currentImg = 7
		} else {
			p.currentImg = 6
		}
	case right:
		if openMouth {
			p.currentImg = 1
		} else {
			p.currentImg = 0
		}
	case down:
		if openMouth {
			p.currentImg = 3
		} else {
			p.currentImg = 2
		}
	case left:
		if openMouth {
			p.currentImg = 5
		} else {
			p.currentImg = 4
		}
	}
}
