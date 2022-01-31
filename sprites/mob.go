package sprites

import (
	"embed"
	"fmt"
	"image"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/rect"
)

//go:embed mob
var mobFiles embed.FS
var mobImages []*ebiten.Image

func init() {
	for x := 0; x < 7; x++ {
		f, err := mobFiles.Open(fmt.Sprintf("mob/mob%d.png", x))
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err.Error())
		}
		mobImages = append(mobImages, ebiten.NewImageFromImage(img))
	}
}

type Mob struct {
	image          *ebiten.Image
	speedy         int
	speedx         int
	rotation       int
	rotation_speed int
	rect           rect.Rect
	isAlive        bool
	lastUpdate     time.Time
	screenWidth    int
	screenHeight   int
}

func NewMob(screenWidth, screenHeight int) *Mob {
	spriteImage := mobImages[rand.Intn(7)]
	spriteBounds := spriteImage.Bounds()
	rect := rect.NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetLeft(rand.Intn(screenWidth - rect.Width()))
	rect.SetTop(rand.Intn(50) - 150)
	return &Mob{
		image:          spriteImage,
		isAlive:        true,
		speedy:         5 + rand.Intn(15),
		speedx:         3 - rand.Intn(7),
		rotation:       0,
		rotation_speed: 2 - rand.Intn(5),
		lastUpdate:     time.Now(),
		rect:           rect,
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
	}
}

func (b *Mob) rotate() {
	if time.Now().After(b.lastUpdate.Add(30 * time.Millisecond)) {
		b.lastUpdate = time.Now()
		b.rotation = (b.rotation + b.rotation_speed) % 360
	}
}

func (b *Mob) Update() {
	b.rotate()
	b.rect.MoveY(b.speedy)
	b.rect.MoveX(b.speedx)

	if (b.rect.Top() > b.screenHeight+10) || (b.rect.Left() < -25) || (b.rect.Right() > b.screenWidth+20) {
		b.rect.SetLeft(rand.Intn(b.screenWidth - b.rect.Width()))
		b.rect.SetTop(rand.Intn(61) - 100)
		b.speedy = rand.Intn(8) + 1
	}
}

func (b *Mob) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(b.rect.Width())/2, -float64(b.rect.Height())/2)
	options.GeoM.Rotate(float64(b.rotation) / (2.0 * math.Pi))
	options.GeoM.Translate(float64(b.rect.Left())+(float64(b.rect.Width())/2), float64(b.rect.Top())+(float64(b.rect.Height())/2))
	screen.DrawImage(b.image, options)
}

func (b *Mob) IsAlive() bool {
	return b.isAlive
}

func (b *Mob) Rect() *rect.Rect {
	return &b.rect
}

func (b *Mob) Kill() {
	b.isAlive = false
}
