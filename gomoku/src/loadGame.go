package main

import (
	"image"
	png "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func loadGame() (pixelgl.WindowConfig, *pixelgl.Window, Sprite) {
	winChan := make(chan *pixelgl.Window)
	cfgChan := make(chan pixelgl.WindowConfig)
	go initWin(winChan, cfgChan)
	go loadAudio()
	cfg, win := <-cfgChan, <-winChan
	return cfg, win, generateSprites()
}

func initWin(winChan chan *pixelgl.Window, cfgChan chan pixelgl.WindowConfig) {
	cfg := pixelgl.WindowConfig{
		Title:     "Gomoku - jjaouen & drimo",
		Bounds:    pixel.R(0, 0, mycfg.X, mycfg.Y),
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}
	mycfg.RefreshRate = pixelgl.PrimaryMonitor().RefreshRate()
	cfgChan <- cfg
	winChan <- win
}

func assignChan() chanSprite {
	var Schan chanSprite
	Schan.Menu = make(chan *pixel.Sprite)
	Schan.Board = make(chan *pixel.Sprite)
	Schan.Off = make(chan *pixel.Sprite)
	Schan.On = make(chan *pixel.Sprite)
	Schan.Soundoff = make(chan *pixel.Sprite)
	Schan.Soundon = make(chan *pixel.Sprite)
	Schan.Whitestone = make(chan *pixel.Sprite)
	Schan.Blackstone = make(chan *pixel.Sprite)
	Schan.Goldstone = make(chan *pixel.Sprite)
	Schan.RedStone = make(chan *pixel.Sprite)
	Schan.BlueStone = make(chan *pixel.Sprite)
	Schan.Transpblackstone = make(chan *pixel.Sprite)
	Schan.Transpwhitestone = make(chan *pixel.Sprite)
	Schan.Winplayerone = make(chan *pixel.Sprite)
	Schan.Winplayertwo = make(chan *pixel.Sprite)
	Schan.Winia = make(chan *pixel.Sprite)
	Schan.Wood = make(chan *pixel.Sprite)
	Schan.LogoSoundOff = make(chan *pixel.Sprite)
	Schan.LogoSoundOn = make(chan *pixel.Sprite)
	Schan.Atlas = make(chan *text.Atlas)
	return Schan
}

func generateSprites() Sprite {
	var Sprites Sprite
	Schan := assignChan()
	go func() {
		pic, err := loadPicture("assets/menu.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Menu <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/board.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Board <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/off.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Off <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/on.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.On <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/whitestone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Whitestone <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/blackstone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Blackstone <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/transparentblackstone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Transpblackstone <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicturePng("assets/transparentwhitestone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Transpwhitestone <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/soundoff.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Soundoff <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/soundon.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Soundon <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/winplayerone.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Winplayerone <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/winplayertwo.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Winplayertwo <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/winia.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Winia <- pixel.NewSprite(pic, pic.Bounds())
	}()
	go func() {
		pic, err := loadPicture("assets/woodtimer.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Wood <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		pic, err := loadPicturePng("assets/logosoundoff.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.LogoSoundOff <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		pic, err := loadPicturePng("assets/logosoundon.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.LogoSoundOn <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		pic, err := loadPicturePng("assets/goldstone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.Goldstone <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		pic, err := loadPicturePng("assets/redstone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.RedStone <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		pic, err := loadPicturePng("assets/bluestone.png")
		if err != nil {
			log.Fatal(err)
		}
		Schan.BlueStone <- pixel.NewSprite(pic, pic.Bounds())
	}()

	go func() {
		face, err := loadTTF("fonts/hiroshima.ttf", 62)
		if err != nil {
			log.Fatal(err)
		}
		Schan.Atlas <- text.NewAtlas(face, text.ASCII)
	}()
	Sprites = Sprite{<-Schan.Menu, <-Schan.Board, <-Schan.Off, <-Schan.On,
		<-Schan.Soundon, <-Schan.Soundoff, <-Schan.Blackstone, <-Schan.Whitestone, <-Schan.Goldstone, <-Schan.RedStone, <-Schan.BlueStone,
		<-Schan.Transpwhitestone, <-Schan.Transpblackstone, <-Schan.Winplayerone,
		<-Schan.Winplayertwo, <-Schan.Winia, <-Schan.Wood, <-Schan.LogoSoundOff, <-Schan.LogoSoundOn,
		<-Schan.Atlas}
	return Sprites
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadPicturePng(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
