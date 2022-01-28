package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PowTypeShield = iota
	PowTypeGun    = iota
)

type Pow struct {
	image   *ebiten.Image
	speedy  int
	rect    Rect
	isAlive bool
	powType int
}

func NewPow(x, y int) (*Pow, error) {
	powType := rand.Intn(2)
	imagePath := "./assets/shield_gold.png"
	if powType == PowTypeGun {
		imagePath = "./assets/bolt_gold.png"
	}
	spriteImage, err := getImageFromFilePath(imagePath)
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
	return &Pow{
		image:   spriteImage,
		isAlive: true,
		speedy:  2,
		powType: powType,
		rect:    rect,
	}, nil
}

func (p *Pow) Update() {
	p.rect.MoveY(p.speedy)
	if p.rect.Top() > screenHeight {
		p.isAlive = false
	}
}

func (p *Pow) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.rect.Left()), float64(p.rect.Top()))
	screen.DrawImage(p.image, options)
}

func (p *Pow) IsAlive() bool {
	return p.isAlive
}

func (p *Pow) Rect() Rect {
	return p.rect
}
