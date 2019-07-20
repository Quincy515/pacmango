package pacman

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type particle struct {
	angle, speed int
	x, y         float64
	alpha        float64
}

type particleManager struct {
	particles [10]*particle
	image     *ebiten.Image
	counter   int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newParticleManager(img []byte, x, y float64) *particleManager {
	pm := &particleManager{}
	pm.image = loadImage(img)
	pm.createParticles(x, y)
	return pm
}

func (pm *particleManager) createParticles(x, y float64) {
	for i := 0; i < len(pm.particles); i++ {
		pm.particles[i] = &particle{
			angle: rand.Intn(361),
			speed: 1,
			alpha: 1,
			x:     x,
			y:     y,
		}
	}
}

func (pm *particleManager) reset(x, y float64) {
	pm.createParticles(x, y)
}

func (pm *particleManager) move() {
	pm.counter++
	for i := 0; i < len(pm.particles); i++ {
		p := pm.particles[i]
		radian := float64(p.angle) * float64(180/math.Pi)
		p.x += float64(p.speed) * math.Cos(radian)
		p.y += float64(p.speed) * math.Sin(radian)
		p.alpha -= 0.009
	}
}

func (pm *particleManager) draw(screen *ebiten.Image) {
	for i := 0; i < len(pm.particles); i++ {
		p := pm.particles[i]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.x), float64(p.y))
		op.ColorM.Scale(1, 1, 1, p.alpha)
		screen.DrawImage(pm.image, op)
	}
}
