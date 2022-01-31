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
	"github.com/jespino/spaceshooter/media"
	"github.com/jespino/spaceshooter/sprite"
	"github.com/jespino/spaceshooter/sprites"
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
)

const (
	ScreenMainMenu = "main-menu"
	ScreenGetReady = "get-ready"
	ScreenGame     = "game"
	ScreenGameOver = "game-over"
)

type Game struct {
	mainMenuMusicPlayer *audio.Player
	getReadyPlayer      *audio.Player
	gameMusicPlayer     *audio.Player
	explosionPlayers    []*audio.Player
	playerDiePlayer     *audio.Player
	screen              string
	mainFont            font.Face
	bigFont             font.Face
	smallFont           font.Face
	score               int
	lives               int
	shield              int
	mainMenuBackground  *ebiten.Image
	gameBackground      *ebiten.Image
	shipImage           *ebiten.Image
	player              *Player
}

func NewGame() (*Game, error) {
	mainMenuImage := media.GetImageFromFilePath(assets, "assets/main.png")
	gameBackground := media.GetImageFromFilePath(assets, "assets/back.png")
	shipImage := media.GetImageFromFilePath(assets, "assets/playerShip1_orange.png")

	mainMenuMusicPlayer := media.GetLoopAudioPlayerFromFile(assets, "sounds/menu.ogg")
	mainMenuMusicPlayer.Play()
	explosion1Player := media.GetAudioPlayerFromFile(assets, "sounds/expl1.ogg")
	explosion2Player := media.GetAudioPlayerFromFile(assets, "sounds/expl2.ogg")
	playerDiePlayer := media.GetAudioPlayerFromFile(assets, "sounds/rumble1.ogg")
	gameMusicPlayer := media.GetLoopAudioPlayerFromFile(assets, "sounds/tgfcoder-FrozenJam-SeamlessLoop.ogg")
	getReadyPlayer := media.GetAudioPlayerFromFile(assets, "sounds/getready.ogg")

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
		mobs.Add(sprites.NewMob(screenWidth, screenHeight))
	}

	return &Game{
		screen:              ScreenMainMenu,
		gameMusicPlayer:     gameMusicPlayer,
		getReadyPlayer:      getReadyPlayer,
		mainMenuMusicPlayer: mainMenuMusicPlayer,
		explosionPlayers:    []*audio.Player{explosion1Player, explosion2Player},
		playerDiePlayer:     playerDiePlayer,
		mainFont:            mainFont,
		bigFont:             bigFont,
		smallFont:           smallFont,
		mainMenuBackground:  mainMenuImage,
		gameBackground:      gameBackground,
		shipImage:           shipImage,
		score:               0,
		lives:               3,
		shield:              100,
		player:              player,
	}, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	switch g.screen {
	case ScreenMainMenu:
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
						g.shield = 100
						g.lives = 3
					}
				}
			}()
		}
		return nil
	case ScreenGame:
		allSprites.Update()
		g.handleBulletsMobsCollitions()
		g.handlePlayerPowerupsCollitions()
		g.handlePlayerMobsCollitions()
	}
	return nil
}

func (g *Game) handleBulletsMobsCollitions() {
	collitions := sprite.SpritesGroupsCollides(bullets, mobs)
	for _, collition := range collitions {
		if !collition.Member1.IsAlive() || !collition.Member2.IsAlive() {
			continue
		}
		expl := g.explosionPlayers[rand.Intn(2)]
		expl.Rewind()
		expl.Play()
		collition.Member1.Kill()
		collition.Member2.Kill()
		mobs.Add(sprites.NewMob(screenWidth, screenHeight))
		explosion := sprites.NewExplosion(collition.Member2.Rect().Width() < 30, collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
		allSprites.Add(explosion)
		if rand.Intn(10) == 9 {
			pow := sprites.NewPow(collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY(), screenHeight)
			powerups.Add(pow)
		}
		g.score += 10
	}
}

func (g *Game) handlePlayerPowerupsCollitions() {
	playerGroup := sprite.SpritesGroup{}
	playerGroup.Add(g.player)
	collitions := sprite.SpritesGroupsCollides(playerGroup, powerups)
	for _, collition := range collitions {
		if collition.Member2.(*sprites.Pow).PowType == sprites.PowTypeShield {
			g.shield += 10 + rand.Intn(20)
			if g.shield > 100 {
				g.shield = 100
			}
			collition.Member2.Kill()
		}
		if collition.Member2.(*sprites.Pow).PowType == sprites.PowTypeGun {
			g.player.powerup()
			collition.Member2.Kill()
		}
	}
}

func (g *Game) handlePlayerMobsCollitions() {
	playerGroup := sprite.SpritesGroup{}
	playerGroup.Add(g.player)
	collitions := sprite.SpritesGroupsCollides(playerGroup, mobs)
	for _, collition := range collitions {
		radius := (float64(collition.Member2.Rect().Width()) * 0.90) / 2
		g.shield -= int(radius * 2)
		explosion := sprites.NewExplosion(collition.Member2.Rect().Width() < 30, collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
		allSprites.Add(explosion)
		if g.shield <= 0 {
			explosion := sprites.NewPlayerExplosion(g.player.Rect().CenterX(), g.player.Rect().CenterY())
			allSprites.Add(explosion)
			g.playerDiePlayer.Rewind()
			g.playerDiePlayer.Play()
			g.lives -= 1
			g.shield = 100
			if g.lives == 0 {
				g.player.Kill()
				go func() {
					time.Sleep(2 * time.Second)
					g.screen = ScreenGameOver
				}()
			}
		}
		collition.Member2.Kill()
		mobs.Add(sprites.NewMob(screenWidth, screenHeight))
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

func (g *Game) drawShield(screen *ebiten.Image) {
	shieldBackground := ebiten.NewImage(102, 1)
	shieldBackground.Fill(&color.White)
	optionsBackground := &ebiten.DrawImageOptions{}
	optionsBackground.GeoM.Translate(float64(10), float64(10))
	screen.DrawImage(shieldBackground, optionsBackground)
	optionsBackground.GeoM.Translate(float64(0), float64(9))
	screen.DrawImage(shieldBackground, optionsBackground)
	shieldBackground = ebiten.NewImage(1, 8)
	shieldBackground.Fill(&color.White)
	optionsBackground.GeoM.Translate(float64(0), float64(-8))
	screen.DrawImage(shieldBackground, optionsBackground)
	optionsBackground.GeoM.Translate(float64(101), float64(0))
	screen.DrawImage(shieldBackground, optionsBackground)

	shield := g.shield
	if g.shield <= 0 {
		shield = 0
	}

	shieldBar := ebiten.NewImage(shield, 8)
	shieldBar.Fill(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	optionsBar := &ebiten.DrawImageOptions{}
	optionsBar.GeoM.Translate(float64(11), float64(11))
	screen.DrawImage(shieldBar, optionsBar)
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
		g.drawShield(screen)
	case ScreenGameOver:
		options := &ebiten.DrawImageOptions{}
		screen.DrawImage(g.gameBackground, options)
		gameOverText := "GAME OVER"
		gameOverRect := text.BoundString(g.bigFont, gameOverText)
		text.Draw(screen, gameOverText, g.bigFont, (screenWidth/2)-(gameOverRect.Max.X/2), (screenHeight/2)-(gameOverRect.Max.Y/2), color.White)
	default:
		panic(fmt.Sprintf("Invalid screen: %s\n", g.screen))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
