package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/rect"
)

type SpritesGroup struct {
	sprites []Sprite
	num     int
}

func (s *SpritesGroup) Update() {
	newSprites := []Sprite{}
	count := 0
	for i := 0; i < s.num; i++ {
		if s.sprites[i].IsAlive() {
			s.sprites[i].Update()
			newSprites = append(newSprites, s.sprites[i])
			count++
		}
	}
	s.sprites = newSprites
	s.num = count
}

func (s *SpritesGroup) Draw(screen *ebiten.Image) {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Draw(screen)
	}
}

func (s *SpritesGroup) Add(sprite Sprite) {
	s.sprites = append(s.sprites, sprite)
	s.num++
}

func (s *SpritesGroup) IsAlive() bool {
	return true
}

func (s *SpritesGroup) Rect() *rect.Rect {
	// TODO: Create a valid rect from the underneath sprites
	rect := rect.NewRect(0, 0, 0, 0)
	return &rect
}

func (s *SpritesGroup) Kill() {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Kill()
	}
}
