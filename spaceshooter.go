package main

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jespino/spaceshooter/sprites"
)

//go:embed assets sounds
var assets embed.FS

var allSprites sprites.SpritesGroup
var bullets sprites.SpritesGroup
var mobs sprites.SpritesGroup
var powerups sprites.SpritesGroup
var audioContext *audio.Context

func init() {
	bullets = sprites.SpritesGroup{}
	mobs = sprites.SpritesGroup{}
	powerups = sprites.SpritesGroup{}
	allSprites = sprites.SpritesGroup{}
	allSprites.Add(&bullets)
	allSprites.Add(&mobs)
	allSprites.Add(&powerups)
	audioContext = audio.NewContext(sampleRate)
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
