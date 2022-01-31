package screens

import (
	"embed"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jespino/spaceshooter/media"
	"golang.org/x/image/font"
)

//go:embed getReady
var getReadyFiles embed.FS

type GetReady struct {
	getReadyPlayer  *audio.Player
	gameMusicPlayer *audio.Player
	gameBackground  *ebiten.Image
	nextScreen      func()
	font            font.Face
	screenWidth     int
	screenHeight    int
}

func NewGetReady(nextScreen func(), gameBackground *ebiten.Image, font font.Face, screenWidth, screenHeight int) *GetReady {
	return &GetReady{
		gameBackground:  gameBackground,
		getReadyPlayer:  media.GetAudioPlayerFromFile(getReadyFiles, "getReady/getready.ogg"),
		gameMusicPlayer: media.GetAudioPlayerFromFile(getReadyFiles, "getReady/mainMusic.ogg"),
		nextScreen:      nextScreen,
		font:            font,
		screenWidth:     screenWidth,
		screenHeight:    screenHeight,
	}
}

func (mm *GetReady) Start() error {
	mm.getReadyPlayer.Rewind()
	mm.getReadyPlayer.Play()
	go func() {
		mm.gameMusicPlayer.SetVolume(0)
		mm.gameMusicPlayer.Play()
		for x := 0; x < 100; x++ {
			time.Sleep(30 * time.Millisecond)
			mm.gameMusicPlayer.SetVolume(float64(x) / 100)
			if x == 50 {
				mm.nextScreen()
			}
		}
	}()
	return nil
}

func (mm *GetReady) Stop() error {
	mm.getReadyPlayer.Pause()
	return nil
}

func (mm *GetReady) Update() error {
	return nil
}

func (mm *GetReady) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(mm.gameBackground, options)
	getReadyText := "GET READY!"
	getReadyRect := text.BoundString(mm.font, getReadyText)
	text.Draw(screen, getReadyText, mm.font, (mm.screenWidth/2)-(getReadyRect.Max.X/2), (mm.screenHeight/2)-(getReadyRect.Max.Y/2), color.White)
}
