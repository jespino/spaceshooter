package main

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed assets sounds
var assets embed.FS

var allSprites SpritesGroup
var bullets SpritesGroup
var mobs SpritesGroup
var audioContext *audio.Context

func init() {
	bullets = SpritesGroup{}
	mobs = SpritesGroup{}
	allSprites = SpritesGroup{}
	allSprites.Add(&bullets)
	allSprites.Add(&mobs)
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
