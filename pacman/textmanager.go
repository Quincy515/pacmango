package pacman

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kgosse/pacmanresources/fonts"
	pacimages "github.com/kgosse/pacmanresources/images"
	"golang.org/x/image/font"
)

const (
	keyText     = "KEYS"
	rText       = "r: Restart"
	hText       = "hjkl: Move"
	livesText   = "LIVES"
	scoreText   = "SCORE"
	restartText = "R: Restart"
	moveText    = "←↓↑→: Move"
	pauseText   = "P: pause"
	soundText   = "s: sound %s"
	studyText   = "Learn From: github.com/kgosse"
)

var (
	arialbdFontTitle font.Face
	arialbdFontBody  font.Face
	gold             = color.RGBA{255, 204, 0, 255}
)

type textManager struct {
	titleFF              font.Face
	bodyFF               font.Face
	entranceFF           font.Face
	keyX, livesX, scoreX int
	studyX               int
	titleY               int
	count                int
	entrance             bool
	gameOverImage        *ebiten.Image
	gameOverAlpha        float64
}

func newTextManager(w, h int) *textManager {
	tm := &textManager{}
	tt, err := truetype.Parse(fonts.Arialbd_ttf)
	if err != nil {
		log.Fatal(err)
	}
	tm.gameOverImage = loadImage(pacimages.GameOver_png[:])

	tm.titleFF = truetype.NewFace(tt, &truetype.Options{
		Size: 24,
	})
	tm.bodyFF = truetype.NewFace(tt, &truetype.Options{
		Size: 14,
	})
	tm.entranceFF = truetype.NewFace(tt, &truetype.Options{
		Size: 70,
	})

	tm.scoreX = w - 5*stageBlocSize
	tm.keyX = 20
	tm.livesX = w/2 - 2*stageBlocSize
	tm.studyX = w/2 - 4*stageBlocSize
	tm.titleY = h + 25

	return tm
}

func (tm *textManager) reinit() {
	tm.gameOverAlpha = 0
}

func (tm *textManager) entranceAnim(b bool) {
	if b {
		tm.count = 0
	}
	tm.entrance = b
}

func (tm *textManager) showGameOverImage(screen *ebiten.Image) {
	tm.gameOverAlpha += 0.01
	if tm.gameOverAlpha > 1 {
		tm.gameOverAlpha = 1
	}
	x := float64(3 * stageBlocSize)
	y := float64(4 * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(1, 1, 1, tm.gameOverAlpha)
	screen.DrawImage(tm.gameOverImage, op)
}

func (tm *textManager) showEntranceAnim(screen *ebiten.Image) {
	if !tm.entrance {
		return
	}
	tm.count++
	three := "3"
	two := "2"
	one := "1"
	goText := "GO!"

	if tm.count <= 60 {
		text.Draw(screen, three, tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 120 {
		text.Draw(screen, two, tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 180 {
		text.Draw(screen, one, tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 240 {
		text.Draw(screen, goText, tm.entranceFF, 7*stageBlocSize, 5*stageBlocSize, gold)
	} else {
		tm.entranceAnim(false)
	}
}

func (tm *textManager) draw(screen *ebiten.Image, score, lives int, pac *ebiten.Image, status string) {
	text.Draw(screen, keyText, tm.titleFF, tm.keyX, tm.titleY, gold)
	text.Draw(screen, rText, tm.bodyFF, tm.keyX, tm.titleY+stageBlocSize, gold)
	text.Draw(screen, hText, tm.bodyFF, tm.keyX, tm.titleY+2*stageBlocSize, gold)
	text.Draw(screen, moveText, tm.bodyFF, tm.keyX, tm.titleY+3*stageBlocSize, gold)
	text.Draw(screen, fmt.Sprintf(soundText, status), tm.bodyFF, tm.keyX, tm.titleY+4*stageBlocSize, gold)
	text.Draw(screen, livesText, tm.titleFF, tm.livesX, tm.titleY, gold)
	for i := lives; 0 < i; i-- {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tm.livesX+(lives-i)*stageBlocSize), float64(tm.titleY+stageBlocSize))
		screen.DrawImage(pac, op)
	}
	text.Draw(screen, studyText, tm.titleFF, tm.studyX, tm.titleY+4*stageBlocSize-9, gold)
	text.Draw(screen, scoreText, tm.titleFF, tm.scoreX, tm.titleY, gold)
	text.Draw(screen, strconv.Itoa(score), tm.titleFF, tm.scoreX, tm.titleY+2*stageBlocSize-9, gold)

	if lives == 0 {
		tm.showGameOverImage(screen)
	}

	tm.showEntranceAnim(screen)
}

func (tm *textManager) livesPos(l int) (x, y float64) {
	x = float64(tm.livesX + l*stageBlocSize)
	y = float64(tm.titleY + stageBlocSize)
	return
}
