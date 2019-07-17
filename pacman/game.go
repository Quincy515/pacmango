package pacman

import "github.com/hajimehoshi/ebiten"

// Game holds all the pacman game data
type Game struct {
	scene *scene
}

// NewGame is a Game constructor
func NewGame() *Game {
	g := &Game{}
	return g
}

// ScreenWidth returns the game screen width
func (g *Game) ScreenWidth() int {
	return 320
}

// ScreenHight returns the game screen height
func (g *Game) ScreenHeight() int {
	return 240
}

// Update updates the screen
func (g *Game) Update(screen *ebiten.Image) error {
	if g.scene == nil {
		g.scene = newScene()
	}
	return g.scene.update(screen)
}
