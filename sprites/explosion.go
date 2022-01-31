package sprites

import (
	"embed"
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jespino/spaceshooter/media"
	"github.com/jespino/spaceshooter/rect"
)

//go:embed explosion
var explosionFiles embed.FS
var explosionAnimation []*ebiten.Image
var explosionPlayers []*audio.Player

func init() {
	for x := 0; x < 9; x++ {
		img := media.GetImageFromFilePath(explosionFiles, fmt.Sprintf("explosion/explosion%d.png", x))
		explosionAnimation = append(explosionAnimation, img)
		explosion1Player := media.GetAudioPlayerFromFile(explosionFiles, "explosion/expl1.ogg")
		explosion2Player := media.GetAudioPlayerFromFile(explosionFiles, "explosion/expl2.ogg")
		explosionPlayers = []*audio.Player{explosion1Player, explosion2Player}
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
	expl := explosionPlayers[rand.Intn(2)]
	expl.Rewind()
	expl.Play()
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
