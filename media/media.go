package media

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const sampleRate = 48000

var audioContext *audio.Context

func init() {
	audioContext = audio.NewContext(sampleRate)
}

func GetAudioFromFile(assets embed.FS, filePath string) *vorbis.Stream {
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

func GetAudioPlayerFromFile(assets embed.FS, filePath string) *audio.Player {
	audioFile := GetAudioFromFile(assets, filePath)
	audioPlayer, err := audioContext.NewPlayer(audioFile)
	if err != nil {
		panic(err.Error)
	}
	return audioPlayer
}

func GetLoopAudioPlayerFromFile(assets embed.FS, filePath string) *audio.Player {
	audioFile := GetAudioFromFile(assets, filePath)
	audioPlayer, err := audioContext.NewPlayer(audio.NewInfiniteLoop(audioFile, audioFile.Length()))
	if err != nil {
		panic(err.Error)
	}
	return audioPlayer
}

func GetImageFromFilePath(assets embed.FS, filePath string) *ebiten.Image {
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

func ImageFromBytes(data []byte) *ebiten.Image {
	image, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err.Error())
	}

	return ebiten.NewImageFromImage(image)
}
