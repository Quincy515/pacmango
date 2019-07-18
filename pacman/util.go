package pacman

func canMove(m [][]elem, p pos) bool {
	return !isWall(m[p.y][p.x])
}

func isWall(e elem) bool {
	if w0 <= e && e <= w24 {
		return true
	}
	return false
}
func addPosDir(d input, p pos) pos {
	r := pos{p.y, p.x}

	switch d {
	case up:
		r.y--
	case right:
		r.x++
	case down:
		r.y++
	case left:
		r.x--
	}

	if r.x < 0 {
		r.x = 0
	}
	if r.y < 0 {
		r.y = 0
	}

	return r
}
func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
