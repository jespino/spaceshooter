package main

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jespino/spaceshooter/sprite"
)

//go:embed assets sounds
var assets embed.FS

var allSprites sprite.SpritesGroup
var bullets sprite.SpritesGroup
var mobs sprite.SpritesGroup
var powerups sprite.SpritesGroup

func init() {
	bullets = sprite.SpritesGroup{}
	mobs = sprite.SpritesGroup{}
	powerups = sprite.SpritesGroup{}
	allSprites = sprite.SpritesGroup{}
	allSprites.Add(&bullets)
	allSprites.Add(&mobs)
	allSprites.Add(&powerups)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Shooter")
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
