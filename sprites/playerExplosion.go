package sprites

import (
	"embed"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/media"
	"github.com/jespino/spaceshooter/rect"
)

//go:embed playerExplosion
var playerExplosionFiles embed.FS
var playerExplosionAnimation []*ebiten.Image

func init() {
	for x := 0; x < 9; x++ {
		img := media.GetImageFromFilePath(playerExplosionFiles, fmt.Sprintf("playerExplosion/playerExplosion%d.png", x))
		playerExplosionAnimation = append(playerExplosionAnimation, img)
	}
}

type PlayerExplosion struct {
	rect       rect.Rect
	isAlive    bool
	lastUpdate time.Time
	small      bool
	frame      int
}

func NewPlayerExplosion(x, y int) *PlayerExplosion {
	spriteBounds := playerExplosionAnimation[0].Bounds()
	rect := rect.NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetCenterX(x)
	rect.SetCenterY(y)
	return &PlayerExplosion{
		isAlive:    true,
		rect:       rect,
		frame:      0,
		lastUpdate: time.Now(),
	}
}

func (b *PlayerExplosion) Update() {
	if time.Now().After(b.lastUpdate.Add(60 * time.Millisecond)) {
		b.frame++
		b.lastUpdate = time.Now()
		if b.frame > 7 {
			b.isAlive = false
		}
	}
}

func (b *PlayerExplosion) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	if b.small {
		options.GeoM.Translate(-float64(b.rect.Width())/2, -float64(b.rect.Height())/2)
		options.GeoM.Scale(0.5, 0.5)
		options.GeoM.Translate(float64(b.rect.Width())/2, float64(b.rect.Height())/2)
	}
	options.GeoM.Translate(float64(b.rect.Left()), float64(b.rect.Top()))
	screen.DrawImage(playerExplosionAnimation[b.frame], options)
}

func (b *PlayerExplosion) IsAlive() bool {
	return b.isAlive
}

func (b *PlayerExplosion) Rect() *rect.Rect {
	return &b.rect
}

func (b *PlayerExplosion) Kill() {
	b.isAlive = false
}
