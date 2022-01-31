package screens

import (
	_ "embed"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jespino/spaceshooter/media"
	"github.com/jespino/spaceshooter/sprite"
	"github.com/jespino/spaceshooter/sprites"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

//go:embed mainGame/ship.png
var shipImageData []byte

type MainGame struct {
	gameBackground *ebiten.Image
	nextScreen     func()
	allSprites     *sprite.SpritesGroup
	bullets        *sprite.SpritesGroup
	mobs           *sprite.SpritesGroup
	powerups       *sprite.SpritesGroup
	player         *sprites.Player
	shipImage      *ebiten.Image
	smallFont      font.Face
	lives          int
	score          int
	shield         int
	screenWidth    int
	screenHeight   int
}

func NewMainGame(nextScreen func(), gameBackground *ebiten.Image, screenWidth, screenHeight int) *MainGame {
	fontObj, err := opentype.Parse(goregular.TTF)
	if err != nil {
		panic(err.Error())
	}
	smallFont, err := opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		panic(err.Error())
	}

	bullets := &sprite.SpritesGroup{}
	mobs := &sprite.SpritesGroup{}
	powerups := &sprite.SpritesGroup{}
	allSprites := &sprite.SpritesGroup{}
	allSprites.Add(bullets)
	allSprites.Add(mobs)
	allSprites.Add(powerups)

	shipImage := media.ImageFromBytes(shipImageData)
	player, err := sprites.NewPlayer(shipImage, bullets, screenWidth/2, screenHeight-100, screenWidth, screenHeight)
	if err != nil {
		panic(err.Error)
	}

	return &MainGame{
		gameBackground: gameBackground,
		nextScreen:     nextScreen,
		allSprites:     allSprites,
		bullets:        bullets,
		mobs:           mobs,
		powerups:       powerups,
		smallFont:      smallFont,
		shipImage:      shipImage,
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
		player:         player,
		lives:          3,
		score:          0,
		shield:         100,
	}
}

func (mg *MainGame) Start() error {
	for x := 0; x < 8; x++ {
		mg.mobs.Add(sprites.NewMob(mg.screenWidth, mg.screenHeight))
	}

	return nil
}

func (mg *MainGame) Stop() error {
	return nil
}

func (mg *MainGame) Update() error {
	mg.player.Update()
	mg.allSprites.Update()
	mg.handleBulletsMobsCollitions()
	mg.handlePlayerPowerupsCollitions()
	mg.handlePlayerMobsCollitions()
	return nil
}

func (mg *MainGame) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(mg.gameBackground, options)
	mg.allSprites.Draw(screen)
	mg.player.Draw(screen)
	mg.drawScore(screen)
	mg.drawLives(screen)
	mg.drawShield(screen)
}

func (mg *MainGame) drawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", mg.score)
	scoreRect := text.BoundString(mg.smallFont, scoreText)
	text.Draw(screen, scoreText, mg.smallFont, (mg.screenWidth/2)-(scoreRect.Max.X/2), 20, color.White)
}

func (mg *MainGame) drawShield(screen *ebiten.Image) {
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

	shield := mg.shield
	if mg.shield <= 0 {
		shield = 1
	}

	shieldBar := ebiten.NewImage(shield, 8)
	shieldBar.Fill(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	optionsBar := &ebiten.DrawImageOptions{}
	optionsBar.GeoM.Translate(float64(11), float64(11))
	screen.DrawImage(shieldBar, optionsBar)
}

func (mg *MainGame) drawLives(screen *ebiten.Image) {
	for x := 0; x < mg.lives; x++ {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(0.2, 0.2)
		options.GeoM.Translate(float64(mg.screenWidth-30-(30*x)), float64(10))
		screen.DrawImage(mg.shipImage, options)
	}
}
func (mg *MainGame) handleBulletsMobsCollitions() {
	collitions := sprite.SpritesGroupsCollides(mg.bullets, mg.mobs)
	for _, collition := range collitions {
		if !collition.Member1.IsAlive() || !collition.Member2.IsAlive() {
			continue
		}
		collition.Member1.Kill()
		collition.Member2.Kill()
		mg.mobs.Add(sprites.NewMob(mg.screenWidth, mg.screenHeight))
		explosion := sprites.NewExplosion(collition.Member2.Rect().Width() < 30, collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
		mg.allSprites.Add(explosion)
		if rand.Intn(10) == 9 {
			pow := sprites.NewPow(collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY(), mg.screenHeight)
			mg.powerups.Add(pow)
		}
		mg.score += 10
	}
}

func (mg *MainGame) handlePlayerPowerupsCollitions() {
	collitions := sprite.SpriteAndGroupCollides(mg.player, mg.powerups)
	for _, collition := range collitions {
		if collition.Member2.(*sprites.Pow).PowType == sprites.PowTypeShield {
			mg.shield += 10 + rand.Intn(20)
			if mg.shield > 100 {
				mg.shield = 100
			}
			collition.Member2.Kill()
		}
		if collition.Member2.(*sprites.Pow).PowType == sprites.PowTypeGun {
			mg.player.Powerup()
			collition.Member2.Kill()
		}
	}
}

func (mg *MainGame) handlePlayerMobsCollitions() {
	collitions := sprite.SpriteAndGroupCollides(mg.player, mg.mobs)
	for _, collition := range collitions {
		radius := (float64(collition.Member2.Rect().Width()) * 0.90) / 2
		mg.shield -= int(radius * 2)
		explosion := sprites.NewExplosion(collition.Member2.Rect().Width() < 30, collition.Member2.Rect().CenterX(), collition.Member2.Rect().CenterY())
		mg.allSprites.Add(explosion)
		if mg.shield <= 0 {
			explosion := sprites.NewPlayerExplosion(mg.player.Rect().CenterX(), mg.player.Rect().CenterY())
			mg.allSprites.Add(explosion)
			mg.lives -= 1
			mg.shield = 100
			if mg.lives == 0 {
				mg.player.Kill()
				go func() {
					time.Sleep(2 * time.Second)
					mg.nextScreen()
				}()
			}
		}
		collition.Member2.Kill()
		mg.mobs.Add(sprites.NewMob(mg.screenWidth, mg.screenHeight))
	}
}
