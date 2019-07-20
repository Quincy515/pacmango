package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"time"
)

type ghost struct {
	kind                   elem
	currentImg             int
	prvPos, curPos, nxtPos pos
	speed                  int
	stepsLength            pos
	steps                  int
	dir                    input
	vision                 int
	ctVulnerable           int
	vulnerableMove         bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func newGhost(y, x int, k elem) *ghost {
	return &ghost{
		kind:        k,
		prvPos:      pos{y, x},
		curPos:      pos{y, x},
		nxtPos:      pos{y, x},
		stepsLength: pos{},
		speed:       4,
		vision:      getVision(k),
	}
}
func getVision(e elem) int {
	switch e {
	case pinkyElem:
		return 10
	case inkyElem:
		return 15
	case blinkyElem:
		return 50
	case clydeElem:
		return 60
	default:
		return 0
	}
}
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	if g.isVulnerable() {
		i := g.currentImg + 8
		if i >= len(imgs) {
			i = 8
		}
		return imgs[i]
	}
	return imgs[g.currentImg]
}

func (g *ghost) isVulnerable() bool {
	if g.ctVulnerable > 0 {
		return true
	}
	return false
}

func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.curPos.x * stageBlocSize)
	y := float64(g.curPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}

func (g *ghost) move() {
	switch g.dir {
	case up:
		g.stepsLength.y -= g.speed
	case right:
		g.stepsLength.x += g.speed
	case down:
		g.stepsLength.y += g.speed
	case left:
		g.stepsLength.x -= g.speed
	}

	if g.steps%4 == 0 {
		g.updateImage()
	}
	g.steps++

	if g.vulnerableMove {
		g.ctVulnerable++
		if g.steps == 16 {
			g.endMove()
			if g.ctVulnerable >= 392 {
				g.endVulnerability()
			}
		}
		return
	}

	if g.steps == 8 {
		g.endMove()
	}
}

func (g *ghost) endVulnerability() {
	g.vulnerableMove = false
	g.ctVulnerable = 0
}

func (g *ghost) updateImage() {
	if g.isVulnerable() {
		if g.ctVulnerable <= 310 {
			if g.currentImg == 0 {
				g.currentImg = 1
			} else {
				g.currentImg = 0
			}
		} else {
			if g.currentImg == 2 {
				g.currentImg = 3
			} else {
				g.currentImg = 2
			}
		}
		return
	}
	switch g.dir {
	case up:
		if g.currentImg == 6 {
			g.currentImg = 7
		} else {
			g.currentImg = 6
		}
	case right:
		if g.currentImg == 0 {
			g.currentImg = 1
		} else {
			g.currentImg = 0
		}
	case down:
		if g.currentImg == 2 {
			g.currentImg = 3
		} else {
			g.currentImg = 2
		}
	case left:
		if g.currentImg == 4 {
			g.currentImg = 5
		} else {
			g.currentImg = 4
		}
	}
}

func (g *ghost) endMove() {
	g.prvPos = g.curPos
	g.curPos = g.nxtPos
	g.stepsLength = pos{0, 0}
	g.steps = 0
}

func (g *ghost) isMoving() bool {
	if g.steps > 0 {
		return true
	}
	return false
}

func (g *ghost) findNextMove(m [][]elem, pac pos) {
	if g.isVulnerable() {
		g.vulnerableMove = true
		g.speed = 2
	} else {
		g.speed = 4
	}

	switch g.localisePlayer(m, pac) {
	case up:
		g.dir = up
	case right:
		g.dir = right
	case down:
		g.dir = down
	case left:
		g.dir = left
	default:
		for _, v := range rand.Perm(5) {
			if v == 0 {
				continue
			}
			dir := input(v)
			np := addPosDir(dir, g.curPos)
			if canMove(m, np) && np != g.prvPos {
				g.dir = dir
				g.nxtPos = np
				return
			}
		}

		g.dir = oppDir(g.dir)
	}
	g.nxtPos = addPosDir(g.dir, g.curPos)
}

func (g *ghost) localisePlayer(m [][]elem, pac pos) input {
	if g.isVulnerable() {
		return 0
	}

	maxY := len(m)
	maxX := len(m[0])

	// up
	if g.curPos.x == pac.x && g.curPos.y > pac.y {
		for y, v := g.curPos.y-1, 1; y >= 0 && v <= g.vision && !isWall(m[y][g.curPos.x]); y, v = y-1, v+1 {
			if y == pac.y {
				return up
			}
		}
	}

	// down
	if g.curPos.x == pac.x && g.curPos.y < pac.y {
		for y, v := g.curPos.y+1, 1; y < maxY && v <= g.vision && !isWall(m[y][g.curPos.x]); y, v = y+1, v+1 {
			if y == pac.y {
				return down
			}
		}
	}

	// right
	if g.curPos.y == pac.y && g.curPos.x < pac.x {
		for x, v := g.curPos.x+1, 1; x < maxX && v <= g.vision && !isWall(m[g.curPos.y][x]); x, v = x+1, v+1 {
			if x == pac.x {
				return right
			}
		}
	}

	// left
	if g.curPos.y == pac.y && g.curPos.x > pac.x {
		for x, v := g.curPos.x-1, 1; x >= 0 && v <= g.vision && !isWall(m[g.curPos.y][x]); x, v = x-1, v+1 {
			if x == pac.x {
				return left
			}
		}
	}
	return 0
}

func (g *ghost) makeVulnerable() {
	g.ctVulnerable = 1
}
