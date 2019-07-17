package pacman

import "github.com/hajimehoshi/ebiten"

// Game holds all the pacman game data
type Game struct {
	scene *scene
}

// NewGame is a Game constructor
func NewGame() *Game {
	g := &Game{}
	g.scene = newScene(nil)
	return g
}

// ScreenWidth returns the game screen width
func (g *Game) ScreenWidth() int {
	return g.scene.screenWidth()
}

// ScreenHight returns the game screen height
func (g *Game) ScreenHeight() int {
	return g.scene.screenHeight()
}

// Update updates the screen
func (g *Game) Update(screen *ebiten.Image) error {
	return g.scene.update(screen)
}
