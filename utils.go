package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func getAudioFromFile(filePath string) *vorbis.Stream {
	f, err := assets.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	s, err := vorbis.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		panic(err.Error())
	}
	return s
}

func getImageFromFilePath(filePath string) *ebiten.Image {
	f, err := assets.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		panic(err.Error())
	}

	return ebiten.NewImageFromImage(image)
}
