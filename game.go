package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 800
	screenHeight = 600

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
	sampleRate  = 48000
)

const (
	ScreenMainMenu = "main-menu"
	ScreenGetReady = "get-ready"
	ScreenGame     = "game"
)

type Game struct {
	mainMenuMusicPlayer *audio.Player
	getReadyPlayer      *audio.Player
	gameMusicPlayer     *audio.Player
	screen              string
	mainFont            font.Face
	bigFont             font.Face
	mainMenuBackground  *ebiten.Image
	gameBackground      *ebiten.Image
}

func NewGame() (*Game, error) {
	mainMenuImage, err := getImageFromFilePath("./assets/main.png")
	if err != nil {
		panic(err.Error())
	}

	gameBackground, err := getImageFromFilePath("./assets/back.png")
	if err != nil {
		panic(err.Error())
	}

	mainMenuMusic, err := getAudioFromFile("./sounds/menu.ogg")
	if err != nil {
		panic(err.Error())
	}

	mainMenuMusicPlayer, err := audioContext.NewPlayer(audio.NewInfiniteLoop(mainMenuMusic, mainMenuMusic.Length()))
	if err != nil {
		return nil, err
	}
	mainMenuMusicPlayer.Play()

	gameMusic, err := getAudioFromFile("./sounds/tgfcoder-FrozenJam-SeamlessLoop.ogg")
	if err != nil {
		panic(err.Error())
	}
	gameMusicPlayer, err := audioContext.NewPlayer(audio.NewInfiniteLoop(gameMusic, gameMusic.Length()))
	if err != nil {
		return nil, err
	}

	getReadyAudio, err := getAudioFromFile("./sounds/getready.ogg")
	if err != nil {
		panic(err.Error())
	}
	getReadyPlayer, err := audioContext.NewPlayer(getReadyAudio)
	if err != nil {
		return nil, err
	}

	fontObj, err := opentype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	mainFont, err := opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
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

	player, err := NewPlayer(screenWidth/2, screenHeight-100)
	if err != nil {
		return nil, err
	}
	allSprites.Add(player)
	return &Game{
		screen:              ScreenMainMenu,
		gameMusicPlayer:     gameMusicPlayer,
		getReadyPlayer:      getReadyPlayer,
		mainMenuMusicPlayer: mainMenuMusicPlayer,
		mainFont:            mainFont,
		bigFont:             bigFont,
		mainMenuBackground:  mainMenuImage,
		gameBackground:      gameBackground,
	}, nil
}

func (g *Game) Update() error {
	if g.screen == ScreenMainMenu {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.screen = ScreenGetReady
			g.mainMenuMusicPlayer.Pause()
			g.getReadyPlayer.Play()
			go func() {
				g.gameMusicPlayer.SetVolume(0)
				g.gameMusicPlayer.Play()
				for x := 0; x < 100; x++ {
					time.Sleep(30 * time.Millisecond)
					g.gameMusicPlayer.SetVolume(float64(x) / 100)
					if x == 50 {
						g.screen = ScreenGame
					}
				}
			}()
		}
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			os.Exit(0)
		}
		return nil
	} else {
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			os.Exit(0)
		}
		allSprites.Update()
	}
	return nil
}

func (g *Game) drawMainMenuScreen(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.mainMenuBackground, options)
	enterOptionText := "Press [ENTER] To Begin"
	enterOptionRect := text.BoundString(g.mainFont, enterOptionText)
	text.Draw(screen, enterOptionText, g.mainFont, (screenWidth/2)-(enterOptionRect.Max.X/2), screenHeight/2, color.White)
	quitOptionText := "or [Q] To Quit"
	quitOptionRect := text.BoundString(g.mainFont, quitOptionText)
	text.Draw(screen, quitOptionText, g.mainFont, (screenWidth/2)-(quitOptionRect.Max.X/2), (screenHeight/2)+40, color.White)
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.screen {
	case ScreenMainMenu:
		g.drawMainMenuScreen(screen)
	case ScreenGetReady:
		options := &ebiten.DrawImageOptions{}
		screen.DrawImage(g.gameBackground, options)
		getReadyText := "GET READY!"
		getReadyRect := text.BoundString(g.bigFont, getReadyText)
		text.Draw(screen, getReadyText, g.bigFont, (screenWidth/2)-(getReadyRect.Max.X/2), (screenHeight/2)-(getReadyRect.Max.Y/2), color.White)
	case ScreenGame:
		options := &ebiten.DrawImageOptions{}
		screen.DrawImage(g.gameBackground, options)
		allSprites.Draw(screen)
	default:
		panic(fmt.Sprintf("Invalid screen: %s\n", g.screen))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
