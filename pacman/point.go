package pacman

import (
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)

type point struct {
	image    *ebiten.Image
	show     bool
	count    int
	x, y     float64
	maxCount int
}

func (p *point) draw(screen *ebiten.Image) {
	p.count++
	p.y--
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, op)

	if p.count >= p.maxCount {
		p.show = false
		p.count = 0
	}
}

type pointManager struct {
	points [5]*point
}

func newPointManage() *pointManager {
	pm := &pointManager{}
	pm.loadImages()
	return pm
}

func (pm *pointManager) loadImages() {
	mc := 22
	for i := 0; i < len(pm.points); i++ {
		mc -= 2
		pm.points[i] = &point{
			image:    loadImage(pacimages.PointImages[i]),
			maxCount: mc,
		}
	}
}

func (pm *pointManager) draw(screen *ebiten.Image) {
	for i := 0; i < len(pm.points); i++ {
		p := pm.points[i]
		if !p.show {
			continue
		}
		p.draw(screen)
	}
}

func (pm *pointManager) show(i int, x, y float64) {
	if i >= len(pm.points) {
		i = len(pm.points) - 1
	}
	p := pm.points[i]
	p.show = true
	p.x = x
	p.y = y
}
