package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var meteorImages []string = []string{
	"assets/meteorBrown_big1.png",
	"assets/meteorBrown_big2.png",
	"assets/meteorBrown_med1.png",
	"assets/meteorBrown_med3.png",
	"assets/meteorBrown_small1.png",
	"assets/meteorBrown_small2.png",
	"assets/meteorBrown_tiny1.png",
}

type Mob struct {
	image          *ebiten.Image
	speedy         int
	speedx         int
	rotation       int
	rotation_speed int
	rect           Rect
	isAlive        bool
	lastUpdate     time.Time
}

func NewMob() *Mob {
	spriteImage := getImageFromFilePath(meteorImages[rand.Intn(7)])
	spriteBounds := spriteImage.Bounds()
	rect := NewRect(
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
		speedy:         rand.Intn(15) + 5,
		speedx:         rand.Intn(7) - 3,
		rotation:       0,
		rotation_speed: rand.Intn(17) - 8,
		lastUpdate:     time.Now(),
		rect:           rect,
	}
}

func (b *Mob) rotate() {
	// TODO
	// time_now = pygame.time.get_ticks()
	// if time_now - self.last_update > 50: # in milliseconds
	// 	self.last_update = time_now
	// 	self.rotation = (self.rotation + self.rotation_speed) % 360
	// 	new_image = pygame.transform.rotate(self.image_orig, self.rotation)
	// 	old_center = self.rect.center
	// 	self.image = new_image
	// 	self.rect = self.image.get_rect()
	// 	self.rect.center = old_center
}

func (b *Mob) Update() {
	b.rotate()
	b.rect.MoveY(b.speedy)
	b.rect.MoveX(b.speedx)

	if (b.rect.Top() > screenHeight+10) || (b.rect.Left() < -25) || (b.rect.Right() > screenWidth+20) {
		b.rect.SetLeft(rand.Intn(screenWidth - b.rect.Width()))
		b.rect.SetTop(rand.Intn(61) - 100)
		b.speedy = rand.Intn(8) + 1
	}
}

func (b *Mob) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.rect.Left()), float64(b.rect.Top()))
	screen.DrawImage(b.image, options)
}

func (b *Mob) IsAlive() bool {
	return b.isAlive
}

func (b *Mob) Rect() *Rect {
	return &b.rect
}

func (b *Mob) Kill() {
	b.isAlive = false
}
