package pacman

import (
	"image/color"
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kgosse/pacmanresources/fonts"
	"golang.org/x/image/font"
)

const (
	keyText     = "KEYS"
	rText       = "R: Restart"
	hText       = "hjkl: Move"
	livesText   = "LIVES"
	scoreText   = "SCORE"
	restartText = "R: Restart"
	moveText    = "←↓↑→: Move"
	pauseText   = "P: pause"
)

var (
	arialbdFontTitle font.Face
	arialbdFontBody  font.Face
	gold             = color.RGBA{255, 204, 0, 255}
)

type textManager struct {
	titleFF              font.Face
	bodyFF               font.Face
	keyX, livesX, scoreX int
	titleY               int
}

func newTextManager(w, h int) *textManager {
	tm := &textManager{}
	tt, err := truetype.Parse(fonts.Arialbd_ttf)
	if err != nil {
		log.Fatal(err)
	}
	tm.titleFF = truetype.NewFace(tt, &truetype.Options{
		Size: 24,
	})
	tm.bodyFF = truetype.NewFace(tt, &truetype.Options{
		Size: 14,
	})

	tm.scoreX = w - 5*stageBlocSize
	tm.keyX = 20
	tm.livesX = w/2 - 2*stageBlocSize
	tm.titleY = h + 25

	return tm
}

func (tm *textManager) draw(screen *ebiten.Image, score, lives int, pac *ebiten.Image) {
	text.Draw(screen, keyText, tm.titleFF, tm.keyX, tm.titleY, gold)
	text.Draw(screen, rText, tm.bodyFF, tm.keyX, tm.titleY+stageBlocSize, gold)
	text.Draw(screen, hText, tm.bodyFF, tm.keyX, tm.titleY+2*stageBlocSize, gold)
	text.Draw(screen, moveText, tm.bodyFF, tm.keyX, tm.titleY+3*stageBlocSize, gold)

	text.Draw(screen, livesText, tm.titleFF, tm.livesX, tm.titleY, gold)
	for i := lives; 0 < i; i-- {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tm.livesX+(lives-i)*stageBlocSize), float64(tm.titleY+stageBlocSize))
		screen.DrawImage(pac, op)
	}

	text.Draw(screen, scoreText, tm.titleFF, tm.scoreX, tm.titleY, gold)
	text.Draw(screen, strconv.Itoa(score), tm.titleFF, tm.scoreX, tm.titleY+2*stageBlocSize-9, gold)
}

func (tm *textManager) livesPos(l int) (x, y float64) {
	x = float64(tm.livesX + l*stageBlocSize)
	y = float64(tm.titleY + stageBlocSize)
	return
}
