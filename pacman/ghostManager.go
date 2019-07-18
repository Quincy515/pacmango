package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type ghostManager struct {
	ghosts              []*ghost
	images              map[elem][8]*ebiten.Image
	vulnerabilityImages [5]*ebiten.Image
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
	gm.vulnerabilityImages = loadVulnerabilityImages()
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

func loadGhostImages(g [8][]byte) [8]*ebiten.Image {
	var arr [8]*ebiten.Image
	for i := 0; i < 8; i++ {
		img, _, err := image.Decode(bytes.NewReader(g[i]))
		handleError(err)
		arr[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return arr
}

func loadVulnerabilityImages() [5]*ebiten.Image {
	var arr [5]*ebiten.Image
	for i := 0; i < 5; i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.VulnerabilityImages[i]))
		handleError(err)
		arr[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return arr
}
