package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var allSprites Sprites
var bullets Sprites
var mobs Sprites
var audioContext *audio.Context

func init() {
	allSprites = Sprites{}
	bullets = Sprites{}
	mobs = Sprites{}
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
