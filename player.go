package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	powerup_time = 1000
	hide_time    = 1000
)

type Player struct {
	image              *ebiten.Image
	speed              int
	shield             int
	shoot_delay        int
	hidden             bool
	power              int
	speedx             int
	power_time         int64
	hide_time          int64
	last_shot          int64
	shot_delay         int64
	lives              int
	rect               Rect
	bulletSoundPlayer  *audio.Player
	missileSoundPlayer *audio.Player
}

func NewPlayer(x, y int) (*Player, error) {
	bulletSound, err := getAudioFromFile("./sounds/pew.ogg")
	if err != nil {
		panic(err.Error())
	}
	bulletSoundPlayer, err := audioContext.NewPlayer(bulletSound)

	missileSound, err := getAudioFromFile("./sounds/rocket.ogg")
	if err != nil {
		panic(err.Error())
	}
	missileSoundPlayer, err := audioContext.NewPlayer(missileSound)

	spriteImage, err := getImageFromFilePath("./assets/playerShip1_orange.png")
	if err != nil {
		return nil, err
	}
	spriteBounds := spriteImage.Bounds()
	rect := NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetCenterX(x)
	rect.SetCenterY(y)
	return &Player{
		image:              spriteImage,
		speedx:             0,
		shield:             100,
		shot_delay:         250,
		lives:              3,
		last_shot:          time.Now().UnixMilli(),
		hidden:             false,
		hide_time:          time.Now().UnixMilli(),
		power:              1,
		power_time:         time.Now().UnixMilli(),
		rect:               rect,
		bulletSoundPlayer:  bulletSoundPlayer,
		missileSoundPlayer: missileSoundPlayer,
	}, nil
}

func (p *Player) Update() {
	if p.power >= 2 && time.Now().UnixMilli()-p.power_time > powerup_time {
		p.power -= 1
		p.power_time = time.Now().UnixMilli()
	}

	if p.hidden && time.Now().UnixMilli()-p.hide_time > hide_time {
		p.hidden = false
		p.rect.SetCenterX(screenWidth / 2)
		p.rect.SetBottom(screenHeight - 30)
	}

	p.speedx = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.speedx = -5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.speedx = 5
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shoot()
	}

	if p.rect.Right() > screenWidth {
		p.rect.SetRight(screenWidth)
	} else if p.rect.Left() < 0 {
		p.rect.SetLeft(0)
	}

	p.rect.MoveX(p.speedx)
}
func (p *Player) shoot() {
	now := time.Now().UnixMilli()
	if now-p.last_shot > p.shot_delay {
		p.last_shot = now
		if p.power == 1 {
			bullet, err := NewBullet(p.rect.CenterX(), p.rect.Top())
			if err != nil {
				panic(err)
			}
			allSprites.Add(bullet)
			bullets.Add(bullet)
			p.bulletSoundPlayer.Rewind()
			p.bulletSoundPlayer.Play()
		}
		if p.power == 2 {
			bullet1, err := NewBullet(p.rect.Left(), p.rect.CenterY())
			if err != nil {
				panic(err)
			}
			bullet2, err := NewBullet(p.rect.Right(), p.rect.CenterY())
			if err != nil {
				panic(err)
			}
			allSprites.Add(bullet1)
			allSprites.Add(bullet2)
			bullets.Add(bullet1)
			bullets.Add(bullet2)
			p.bulletSoundPlayer.Rewind()
			p.bulletSoundPlayer.Play()
		}

		if p.power >= 3 {
			bullet1, err := NewBullet(p.rect.Left(), p.rect.CenterY())
			if err != nil {
				panic(err)
			}
			bullet2, err := NewBullet(p.rect.Right(), p.rect.CenterY())
			if err != nil {
				panic(err)
			}
			missile1, err := NewMissile(p.rect.CenterX(), p.rect.Top())
			if err != nil {
				panic(err)
			}
			allSprites.Add(bullet1)
			allSprites.Add(bullet2)
			allSprites.Add(missile1)
			bullets.Add(bullet1)
			bullets.Add(bullet2)
			bullets.Add(missile1)
			p.bulletSoundPlayer.Rewind()
			p.bulletSoundPlayer.Play()
			p.missileSoundPlayer.Rewind()
			p.missileSoundPlayer.Play()
		}
	}
}

func (p *Player) powerup() {
	p.power += 1
	p.power_time = time.Now().UnixMilli()
}

func (p *Player) hide() {
	p.hidden = true
	p.hide_time = time.Now().UnixMilli()
	p.rect.SetCenterX(screenWidth / 29)
	p.rect.SetCenterY(screenHeight + 200)
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.rect.Left()), float64(p.rect.Top()))
	screen.DrawImage(p.image, options)
}

func (p *Player) IsAlive() bool {
	return true
}

func (p *Player) Rect() Rect {
	return p.rect
}
