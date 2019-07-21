package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func keyPressed() input {
	if inpututil.KeyPressDuration(ebiten.KeyUp) > 0 || inpututil.KeyPressDuration(ebiten.KeyK) > 0 {
		return up
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyH) > 0 {
		return left
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 || inpututil.KeyPressDuration(ebiten.KeyL) > 0 {
		return right
	}
	if inpututil.KeyPressDuration(ebiten.KeyDown) > 0 || inpututil.KeyPressDuration(ebiten.KeyJ) > 0 {
		return down
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return sKey
	}
	return 0
}
