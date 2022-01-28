package main

import "github.com/hajimehoshi/ebiten/v2"

type Missile struct {
	image   *ebiten.Image
	speedy  int
	rect    Rect
	isAlive bool
}

func NewMissile(x, y int) (*Missile, error) {
	spriteImage := getImageFromFilePath("assets/missile.png")
	spriteBounds := spriteImage.Bounds()
	rect := NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetCenterX(x)
	rect.SetCenterY(y)
	return &Missile{
		image:   spriteImage,
		isAlive: true,
		speedy:  -10,
		rect:    rect,
	}, nil
}

func (b *Missile) Update() {
	b.rect.MoveY(b.speedy)
	if b.rect.Top() < 0 {
		b.isAlive = false
	}
}

func (b *Missile) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.rect.Left()), float64(b.rect.Top()))
	screen.DrawImage(b.image, options)
}

func (b *Missile) IsAlive() bool {
	return b.isAlive
}

func (b *Missile) Rect() *Rect {
	return &b.rect
}

func (b *Missile) Kill() {
	b.isAlive = false
}
