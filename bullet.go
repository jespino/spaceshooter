package main

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	image   *ebiten.Image
	speedy  int
	rect    Rect
	isAlive bool
}

func NewBullet(x, y int) (*Bullet, error) {
	spriteImage, err := getImageFromFilePath("./assets/laserRed16.png")
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
	return &Bullet{
		image:   spriteImage,
		isAlive: true,
		speedy:  -10,
		rect:    rect,
	}, nil
}

func (b *Bullet) Update() {
	b.rect.MoveY(b.speedy)
	if b.rect.Top() < 0 {
		b.isAlive = false
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.rect.Left()), float64(b.rect.Top()))
	screen.DrawImage(b.image, options)
}

func (b *Bullet) IsAlive() bool {
	return b.isAlive
}

func (b *Bullet) Rect() Rect {
	return b.rect
}
