package pacman

import (
	"container/list"

	"github.com/hajimehoshi/ebiten"
)

type explosionManager struct {
	explosions *list.List
}

func newExplosionManager() *explosionManager {
	em := &explosionManager{}
	em.explosions = list.New()
	return em
}

func (em *explosionManager) addExplosion(img []byte, x, y float64) {
	p := newParticleManager(img, x, y)
	em.explosions.PushBack(p)
}

func (em *explosionManager) reinit() {
	em.explosions.Init()
}

func (em *explosionManager) move() {
	var toRemove []*list.Element
	for e := em.explosions.Front(); e != nil; e = e.Next() {
		pm := e.Value.(*particleManager)
		if pm.counter == 90 {
			toRemove = append(toRemove, e)
			continue
		}
		pm.move()
	}
	for i := 0; i < len(toRemove); i++ {
		em.explosions.Remove(toRemove[i])
	}
}

func (em *explosionManager) draw(screen *ebiten.Image) {
	for e := em.explosions.Front(); e != nil; e = e.Next() {
		pm := e.Value.(*particleManager)
		pm.draw(screen)
	}
}
