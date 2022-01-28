package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	Update()
	Draw(screen *ebiten.Image)
	IsAlive() bool
	Rect() Rect
}

type Sprites struct {
	sprites []Sprite
	num     int
}

func (s *Sprites) Update() {
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

func (s *Sprites) Draw(screen *ebiten.Image) {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Draw(screen)
	}
}

func (s *Sprites) Add(sprite Sprite) {
	s.sprites = append(s.sprites, sprite)
	s.num++
}

func (s *Sprites) IsAlive() bool {
	return true
}

func (s *Sprites) Rect() Rect {
	// TODO: Create a valid rect from the underneath sprites
	return NewRect(0, 0, 0, 0)
}
