package sprites

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func imageFromBytes(data []byte) *ebiten.Image {
	image, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err.Error())
	}

	return ebiten.NewImageFromImage(image)
}
