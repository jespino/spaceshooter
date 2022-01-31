package main

import (
	_ "embed"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/media"
	"github.com/jespino/spaceshooter/screens"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/back.png
var backgroundImage []byte

type Screen interface {
	Start() error
	Stop() error
	Update() error
	Draw(screen *ebiten.Image)
}

const (
	screenWidth  = 600
	screenHeight = 800
)

type Game struct {
	currentScreen Screen
}

func NewGame() (*Game, error) {
	gameBackground := media.ImageFromBytes(backgroundImage)

	fontObj, err := opentype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	bigFont, err := opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    72,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}
	game := &Game{
		currentScreen: nil,
	}

	gameOverScreen := screens.NewGameOver(gameBackground, bigFont, screenWidth, screenHeight)
	mainGameScreen := screens.NewMainGame(func() { game.SetCurrentScreen(gameOverScreen) }, gameBackground, screenWidth, screenHeight)
	getReadyScreen := screens.NewGetReady(func() { game.SetCurrentScreen(mainGameScreen) }, gameBackground, bigFont, screenWidth, screenHeight)
	mainMenuScreen := screens.NewMainMenu(func() { game.SetCurrentScreen(getReadyScreen) }, screenWidth, screenHeight)
	game.currentScreen = mainMenuScreen
	game.currentScreen.Start()
	return game, nil
}

func (g *Game) SetCurrentScreen(newScreen Screen) {
	g.currentScreen.Stop()
	g.currentScreen = newScreen
	g.currentScreen.Start()
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	return g.currentScreen.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
