package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var explosionImages []string = []string{
	"assets/regularExplosion00.png",
	"assets/regularExplosion01.png",
	"assets/regularExplosion02.png",
	"assets/regularExplosion03.png",
	"assets/regularExplosion04.png",
	"assets/regularExplosion05.png",
	"assets/regularExplosion06.png",
	"assets/regularExplosion07.png",
	"assets/regularExplosion08.png",
}
var explosionAnimation []*ebiten.Image

func init() {
	for x := 0; x < len(explosionImages); x++ {
		img := getImageFromFilePath(explosionImages[x])
		explosionAnimation = append(explosionAnimation, img)
	}
}

type Explosion struct {
	rect       Rect
	isAlive    bool
	lastUpdate time.Time
	small      bool
	frame      int
}

func NewExplosion(small bool, x, y int) *Explosion {
	spriteBounds := explosionAnimation[0].Bounds()
	rect := NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetCenterX(x)
	rect.SetCenterY(y)
	return &Explosion{
		isAlive:    true,
		rect:       rect,
		small:      small,
		frame:      0,
		lastUpdate: time.Now(),
	}
}

func (b *Explosion) Update() {
	if time.Now().After(b.lastUpdate.Add(60 * time.Millisecond)) {
		b.frame++
		b.lastUpdate = time.Now()
		if b.frame > 7 {
			b.isAlive = false
		}
	}
}

func (b *Explosion) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	if b.small {
		options.GeoM.Translate(-float64(b.rect.Width())/2, -float64(b.rect.Height())/2)
		options.GeoM.Scale(0.5, 0.5)
		options.GeoM.Translate(float64(b.rect.Width())/2, float64(b.rect.Height())/2)
	}
	options.GeoM.Translate(float64(b.rect.Left()), float64(b.rect.Top()))
	screen.DrawImage(explosionAnimation[b.frame], options)
}

func (b *Explosion) IsAlive() bool {
	return b.isAlive
}

func (b *Explosion) Rect() *Rect {
	return &b.rect
}

func (b *Explosion) Kill() {
	b.isAlive = false
}
