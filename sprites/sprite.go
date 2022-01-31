package sprites

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

type Collition struct {
	Member1 Sprite
	Member2 Sprite
}

func SpritesCollides(s1 Sprite, s2 Sprite) *Collition {
	if s1.IsAlive() && s2.IsAlive() {
		rect1 := s1.Rect()
		rect2 := s2.Rect()
		if rect1.Overlaps(rect2.Rectangle) {
			return &Collition{s1, s2}
		}
	}
	return nil
}
