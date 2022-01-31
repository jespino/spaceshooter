package screens

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type GameOver struct {
	gameBackground *ebiten.Image
	font           font.Face
	screenWidth    int
	screenHeight   int
}

func NewGameOver(gameBackground *ebiten.Image, font font.Face, screenWidth, screenHeight int) *GameOver {
	return &GameOver{
		gameBackground: gameBackground,
		font:           font,
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
	}
}

func (mm *GameOver) Start() error {
	return nil
}

func (mm *GameOver) Stop() error {
	return nil
}

func (mm *GameOver) Update() error {
	return nil
}

func (mm *GameOver) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(mm.gameBackground, options)
	gameOverText := "GAME OVER"
	gameOverRect := text.BoundString(mm.font, gameOverText)
	text.Draw(screen, gameOverText, mm.font, (mm.screenWidth/2)-(gameOverRect.Max.X/2), (mm.screenHeight/2)-(gameOverRect.Max.Y/2), color.White)
}
