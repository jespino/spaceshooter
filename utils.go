package main

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func getAudioFromFile(filePath string) (*vorbis.Stream, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	s, err := vorbis.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		return nil, err
	}
	return s, err
}

func getImageFromFilePath(filePath string) (*ebiten.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(image), nil
}
