package pacman

func isWall(e elem) bool {
	if w0 <= e && e <= w24 {
		return true
	}
	return false
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
