package sprites

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/rect"
)

//go:embed explosion
var explosionFiles embed.FS

var explosionAnimation []*ebiten.Image

func init() {
	for x := 0; x < 9; x++ {
		f, err := explosionFiles.Open(fmt.Sprintf("explosion/explosion%d.png", x))
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err.Error())
		}
		explosionAnimation = append(explosionAnimation, ebiten.NewImageFromImage(img))
	}
}

type Explosion struct {
	rect       rect.Rect
	isAlive    bool
	lastUpdate time.Time
	small      bool
	frame      int
}

func NewExplosion(small bool, x, y int) *Explosion {
	spriteBounds := explosionAnimation[0].Bounds()
	rect := rect.NewRect(
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

func (b *Explosion) Rect() *rect.Rect {
	return &b.rect
}

func (b *Explosion) Kill() {
	b.isAlive = false
}
