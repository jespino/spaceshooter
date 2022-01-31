package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/rect"
)

type Sprite interface {
	Update()
	Draw(screen *ebiten.Image)
	IsAlive() bool
	Rect() *rect.Rect
	Kill()
}
