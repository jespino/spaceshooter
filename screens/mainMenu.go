package screens

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jespino/spaceshooter/media"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

//go:embed mainMenu
var mainMenuFiles embed.FS

type MainMenu struct {
	mainMenuMusicPlayer *audio.Player
	mainMenuBackground  *ebiten.Image
	nextScreen          func()
	mainFont            font.Face
	screenWidth         int
	screenHeight        int
}

func NewMainMenu(nextScreen func(), screenWidth, screenHeight int) *MainMenu {
	fontObj, err := opentype.Parse(goregular.TTF)
	if err != nil {
		panic(err.Error())
	}
	mainFont, err := opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		panic(err.Error())
	}
	return &MainMenu{
		mainMenuBackground:  media.GetImageFromFilePath(mainMenuFiles, "mainMenu/main.png"),
		mainMenuMusicPlayer: media.GetLoopAudioPlayerFromFile(mainMenuFiles, "mainMenu/menu.ogg"),
		nextScreen:          nextScreen,
		mainFont:            mainFont,
		screenWidth:         screenWidth,
		screenHeight:        screenHeight,
	}
}

func (mm *MainMenu) Start() error {
	mm.mainMenuMusicPlayer.Play()
	return nil
}

func (mm *MainMenu) Stop() error {
	mm.mainMenuMusicPlayer.Pause()
	return nil
}

func (mm *MainMenu) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		mm.nextScreen()
	}
	return nil
}

func (mm *MainMenu) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(mm.mainMenuBackground, options)
	enterOptionText := "Press [ENTER] To Begin"
	enterOptionRect := text.BoundString(mm.mainFont, enterOptionText)
	text.Draw(screen, enterOptionText, mm.mainFont, (mm.screenWidth/2)-(enterOptionRect.Max.X/2), mm.screenHeight/2, color.White)
	quitOptionText := "or [Q] To Quit"
	quitOptionRect := text.BoundString(mm.mainFont, quitOptionText)
	text.Draw(screen, quitOptionText, mm.mainFont, (mm.screenWidth/2)-(quitOptionRect.Max.X/2), (mm.screenHeight/2)+40, color.White)
}
