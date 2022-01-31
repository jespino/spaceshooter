package sprites

import (
	_ "embed"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/rect"
)

//go:embed pow_shield.png
var powShieldImage []byte

//go:embed pow_gun.png
var powGunImage []byte

const (
	PowTypeShield = iota
	PowTypeGun    = iota
)

type Pow struct {
	PowType      int
	image        *ebiten.Image
	speedy       int
	rect         rect.Rect
	isAlive      bool
	screenHeight int
}

func NewPow(x, y, screenHeight int) *Pow {
	powType := rand.Intn(2)
	var spriteImage *ebiten.Image
	if powType == PowTypeGun {
		spriteImage = imageFromBytes(powGunImage)
	} else {
		spriteImage = imageFromBytes(powShieldImage)
	}
	spriteBounds := spriteImage.Bounds()
	rect := rect.NewRect(
		spriteBounds.Min.X,
		spriteBounds.Min.Y,
		spriteBounds.Max.X,
		spriteBounds.Max.Y,
	)
	rect.SetCenterX(x)
	rect.SetCenterY(y)
	return &Pow{
		image:        spriteImage,
		isAlive:      true,
		speedy:       2,
		PowType:      powType,
		rect:         rect,
		screenHeight: screenHeight,
	}
}

func (p *Pow) Update() {
	p.rect.MoveY(p.speedy)
	if p.rect.Top() > p.screenHeight {
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

func (p *Pow) Rect() *rect.Rect {
	return &p.rect
}

func (p *Pow) Kill() {
	p.isAlive = false
}
