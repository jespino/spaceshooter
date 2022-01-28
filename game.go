package main

import (
	"fmt"
	"image/color"
	"math/rand"
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
	screenWidth  = 600
	screenHeight = 800

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
	smallFont           font.Face
	score               int
	lives               int
	mainMenuBackground  *ebiten.Image
	gameBackground      *ebiten.Image
	shipImage           *ebiten.Image
}

func NewGame() (*Game, error) {
	mainMenuImage := getImageFromFilePath("assets/main.png")
	gameBackground := getImageFromFilePath("assets/back.png")
	shipImage := getImageFromFilePath("assets/playerShip1_orange.png")
	mainMenuMusic := getAudioFromFile("sounds/menu.ogg")
	mainMenuMusicPlayer, err := audioContext.NewPlayer(audio.NewInfiniteLoop(mainMenuMusic, mainMenuMusic.Length()))
	if err != nil {
		return nil, err
	}
	mainMenuMusicPlayer.Play()

	gameMusic := getAudioFromFile("sounds/tgfcoder-FrozenJam-SeamlessLoop.ogg")
	gameMusicPlayer, err := audioContext.NewPlayer(audio.NewInfiniteLoop(gameMusic, gameMusic.Length()))
	if err != nil {
		return nil, err
	}

	getReadyAudio := getAudioFromFile("sounds/getready.ogg")
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

	smallFont, err := opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    16,
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
	for x := 0; x < 8; x++ {
		mobs.Add(NewMob())
	}

	return &Game{
		screen:              ScreenMainMenu,
		gameMusicPlayer:     gameMusicPlayer,
		getReadyPlayer:      getReadyPlayer,
		mainMenuMusicPlayer: mainMenuMusicPlayer,
		mainFont:            mainFont,
		bigFont:             bigFont,
		smallFont:           smallFont,
		mainMenuBackground:  mainMenuImage,
		gameBackground:      gameBackground,
		shipImage:           shipImage,
		score:               0,
		lives:               3,
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
		g.handleBulletsMobsCollitions()
	}
	return nil
}

func (g *Game) handleBulletsMobsCollitions() {
	collitions := SpritesGroupsCollides(bullets, mobs)
	for _, collition := range collitions {
		collition.Member1.Kill()
		collition.Member2.Kill()
		mobs.Add(NewMob())
		explosion := NewExplosion(collition.Member2.Rect().Width() < 30, collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
		allSprites.Add(explosion)
		if rand.Intn(10) == 9 {
			pow := NewPow(collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
			powerups.Add(pow)
		}
		g.score += 10
	}
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

func (g *Game) drawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", g.score)
	scoreRect := text.BoundString(g.smallFont, scoreText)
	text.Draw(screen, scoreText, g.smallFont, (screenWidth/2)-(scoreRect.Max.X/2), 20, color.White)
}

func (g *Game) drawLives(screen *ebiten.Image) {
	for x := 0; x < g.lives; x++ {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(0.2, 0.2)
		options.GeoM.Translate(float64(screenWidth-30-(30*x)), float64(10))
		screen.DrawImage(g.shipImage, options)
	}
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
		g.drawScore(screen)
		g.drawLives(screen)
	default:
		panic(fmt.Sprintf("Invalid screen: %s\n", g.screen))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
