package pacman

import (
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
	"math"
)

type ghostManager struct {
	ghosts              []*ghost
	images              map[elem][8]*ebiten.Image
	vulnerabilityImages [5]*ebiten.Image
	eaten               int
}

func newGhostManager() *ghostManager {
	gm := &ghostManager{}
	gm.images = make(map[elem][8]*ebiten.Image)
	gm.loadImages()
	return gm
}

func (gm *ghostManager) loadImages() {
	gm.images[blinkyElem] = loadGhostImages(pacimages.BlinkyImages)
	gm.images[clydeElem] = loadGhostImages(pacimages.ClydeImages)
	gm.images[inkyElem] = loadGhostImages(pacimages.InkyImages)
	gm.images[pinkyElem] = loadGhostImages(pacimages.PinkyImages)
	copy(gm.vulnerabilityImages[:], loadImages(pacimages.VulnerabilityImages[:]))
}

func (gm *ghostManager) addGhost(y, x int, e elem) {
	gm.ghosts = append(gm.ghosts, newGhost(y, x, e))
}

func (gm *ghostManager) draw(screen *ebiten.Image) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		imgs, _ := gm.images[g.kind]
		images := make([]*ebiten.Image, 13)
		copy(images, imgs[:])
		copy(images[8:], gm.vulnerabilityImages[:])
		g.draw(screen, images)
	}
}

func (gm *ghostManager) move(m [][]elem, pac pos) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		if !g.isMoving() {
			g.findNextMove(m, pac)
		}
		g.move()
	}
}

func (gm *ghostManager) detectCollision(pY, pX float64, cb func(bool, float64, float64)) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		gY, gX := g.screenPos()
		if math.Abs(pY-gY) < 32 && math.Abs(pX-gX) < 32 {
			if !g.isVulnerable() {
				cb(false, 0, 0)
				return
			}
			gm.eaten++
			g.makeEaten()
			g.reset()
			cb(true, gY, gX)
		}
	}
}

func (gm *ghostManager) reset(em *explosionManager) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		y, x := g.screenPos()
		em.addExplosion(pacimages.PacParticle_png, x, y)
		g.reset()
	}
}

func loadGhostImages(g [8][]byte) [8]*ebiten.Image {
	var arr [8]*ebiten.Image
	copy(arr[:], loadImages(g[:]))
	return arr
}

func (gm *ghostManager) makeVulnerable() {
	gm.eaten = 0
	for i := 0; i < len(gm.ghosts); i++ {
		gm.ghosts[i].makeVulnerable()
	}
}
