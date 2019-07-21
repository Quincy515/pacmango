package pacman

import (
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type player struct {
	images                 [8]*ebiten.Image
	currentImg             int
	prvPos, curPos, nxtPos pos
	initialPos             pos
	speed                  int
	stepsLength            pos
	steps                  int
	dir                    input // direction
	score                  int
	ctExplosion            int
	pm                     *particleManager
	lost                   bool
}

func newPlayer(y, x int) *player {
	p := &player{}
	p.loadImages()
	p.pm = newParticleManager(pacimages.PacParticle_png, 0, 0)
	p.curPos = pos{y, x}
	p.prvPos = pos{y, x}
	p.nxtPos = pos{y, x}
	p.initialPos = pos{y, x}
	return p
}

func (p *player) loadImages() {
	copy(p.images[:], loadImages(pacimages.PlayerImages[:]))
}

func (p *player) image() *ebiten.Image {
	return p.images[p.currentImg]
}

func (p *player) draw(screen *ebiten.Image) {
	if p.isExploding() {
		p.pm.draw(screen)
		return
	}
	if p.lost {
		return
	}
	x := float64(p.curPos.x*stageBlocSize + p.stepsLength.x)
	y := float64(p.curPos.y*stageBlocSize + p.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(p.image(), op)
}

func (p *player) move(m [][]elem, dir input, cb func()) {
	// explosion
	if p.isExploding() {
		p.ctExplosion++
		p.pm.move()
		// end explosion
		if p.ctExplosion == 90 {
			p.reset()
			cb()
		}
		return
	}
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

func (p *player) reset() {
	p.curPos, p.prvPos, p.nxtPos = p.initialPos, p.initialPos, p.initialPos
	p.currentImg = 0
	p.ctExplosion = 0
	p.stepsLength = pos{0, 0}
	p.steps = 0
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

func (p *player) screenPos() (y, x float64) {
	x = float64(p.curPos.x*stageBlocSize + p.stepsLength.x)
	y = float64(p.curPos.y*stageBlocSize + p.stepsLength.y)
	return
}

func (p *player) explode() {
	if p.isExploding() {
		return
	}
	x := float64(p.curPos.x*stageBlocSize + p.stepsLength.x)
	y := float64(p.curPos.y*stageBlocSize + p.stepsLength.y)
	p.curPos.x = 0
	p.curPos.y = 0
	p.pm.reset(x, y)
	p.ctExplosion = 1
}

func (p *player) isExploding() bool {
	if p.ctExplosion > 0 {
		return true
	}
	return false
}

func (p *player) reinit() {
	p.reset()
	p.score = 0
	p.lost = false
}

func (p *player) gameover() {
	p.lost = true
	p.curPos.x = 0
	p.curPos.y = 0
}
