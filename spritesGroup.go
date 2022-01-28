package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	Update()
	Draw(screen *ebiten.Image)
	IsAlive() bool
	Rect() *Rect
	Kill()
}

type SpritesGroup struct {
	sprites []Sprite
	num     int
}

type Collition struct {
	Member1 Sprite
	Member2 Sprite
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

func (s *SpritesGroup) Rect() *Rect {
	// TODO: Create a valid rect from the underneath sprites
	rect := NewRect(0, 0, 0, 0)
	return &rect
}

func SpritesGroupsCollides(g1 SpritesGroup, g2 SpritesGroup) []Collition {
	collitions := []Collition{}
	for i := 0; i < g1.num; i++ {
		if g1.sprites[i].IsAlive() {
			rect1 := g1.sprites[i].Rect()
			for j := 0; j < g2.num; j++ {
				rect2 := g2.sprites[j].Rect()
				if rect1.Overlaps(rect2.Rectangle) {
					collitions = append(collitions, Collition{g1.sprites[i], g2.sprites[j]})
				}
			}
		}
	}
	return collitions
}

func (s *SpritesGroup) Kill() {
}
