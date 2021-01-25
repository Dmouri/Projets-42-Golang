package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func drawing(win *pixelgl.Window, Sprites Sprite) {
	if state.Menu && redraw {
		drawMenu(win, Sprites)
		redraw = false
	} else if redraw {
		drawBoard(win, Sprites)
		drawInfo(win, Sprites)
		redraw = false
	}
}

func drawMenu(win *pixelgl.Window, sprites Sprite) {
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V((X/1920.0), (Y/1080)))
	sprites.Menu.Draw(win, mat)
	vector := pixel.Vec{X: X * 0.949479, Y: Y - Y*0.087963}
	mat = pixel.IM
	mat = mat.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/512)*0.04114583, (Y/512)*0.073148))
	if mycfg.Sound {
		sprites.LogoSoundOn.Draw(win, mat)
	} else {
		sprites.LogoSoundOff.Draw(win, mat)
	}
	drawSelectStone(win, sprites)
}

func drawAStone(win *pixelgl.Window, vector pixel.Vec, stone *pixel.Sprite) {
	mat := pixel.IM
	mat = mat.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/384)*0.0187, ((Y/384)*0.0325)))
	stone.Draw(win, mat)
}

func drawSelectStone(win *pixelgl.Window, sprites Sprite) {
	if state.IA {
		drawAStone(win, pixel.Vec{X: X * 0.038542, Y: Y - Y*0.39259}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.248438, Y: Y - Y*0.39259}, sprites.BlackStone)
	} else if state.TwoP {
		drawAStone(win, pixel.Vec{X: X * 0.038542, Y: Y - Y*0.59537}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.248438, Y: Y - Y*0.59537}, sprites.BlackStone)
	}
	if state.Capture {
		drawAStone(win, pixel.Vec{X: X * 0.376563, Y: Y - Y*0.39259}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.588542, Y: Y - Y*0.39259}, sprites.BlackStone)
	}
	if state.DoubleThree {
		drawAStone(win, pixel.Vec{X: X * 0.376563, Y: Y - Y*0.59537}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.588542, Y: Y - Y*0.59537}, sprites.BlackStone)
	}
	if state.Standard {
		drawAStone(win, pixel.Vec{X: X * 0.711459, Y: Y - Y*0.39259}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.923958, Y: Y - Y*0.39259}, sprites.BlackStone)
	} else if state.Pro {
		drawAStone(win, pixel.Vec{X: X * 0.711459, Y: Y - Y*0.59537}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.923958, Y: Y - Y*0.59537}, sprites.BlackStone)
	} else if state.Long {
		drawAStone(win, pixel.Vec{X: X * 0.711459, Y: Y - Y*0.79074}, sprites.WhiteStone)
		drawAStone(win, pixel.Vec{X: X * 0.923958, Y: Y - Y*0.79074}, sprites.BlackStone)
	}
}

func drawUnstatic(win *pixelgl.Window, sprites Sprite) {
	var sound *pixel.Sprite
	var vsync *pixel.Sprite
	var aa *pixel.Sprite
	if mycfg.Sound {
		sound = sprites.SoundOn
	} else {
		sound = sprites.SoundOff
	}
	vector := pixel.Vec{X: X * 0.877083, Y: Y * 0.5259259}
	mat := pixel.IM
	mat = pixel.IM.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V(((X/3518)*0.18385), ((Y/1628)*0.12129)))
	sound.Draw(win, mat)
	if mycfg.Vsync {
		vsync = sprites.On
	} else {
		vsync = sprites.Off
	}
	vector = pixel.Vec{X: X * 0.960416, Y: Y * 0.39907407}
	mat = pixel.IM
	mat = pixel.IM.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/1920)*0.04635416, ((Y/1080)*0.041)))
	vsync.Draw(win, mat)
	if mycfg.Smooth {
		aa = sprites.On
	} else {
		aa = sprites.Off
	}
	vector = pixel.Vec{X: X * 0.960416, Y: Y * 0.3444444}
	mat = pixel.IM
	mat = pixel.IM.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/1920)*0.04635416, ((Y/1080)*0.041)))
	aa.Draw(win, mat)
}

func drawStone(win *pixelgl.Window, x, y int8, stone *pixel.Sprite) {
	vector := pixel.Vec{X: (X*0.26510417 + ((0.0260416 * X) * float64(x))) + float64(x)*0.00008*X,
		Y: (Y*0.0842593 + ((0.04537037 * Y) * float64(y))) + float64(y)*0.000926*Y}
	mat := pixel.IM
	mat = pixel.IM.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/384)*0.0187, ((Y/384)*0.0325)))
	stone.Draw(win, mat)
}

func drawNotif(win *pixelgl.Window, sprites Sprite) {

	if time.Now().Sub(saveNotif.Timer).Seconds() > 3.0 {
		saveNotif.Draw = false
	} else {
		vector := pixel.Vec{X: X * 0.43, Y: Y * 0.95}
		txt := text.New(vector, sprites.Atlas)
		txt.Color = colornames.Black
		txt.WriteString(fmt.Sprintf("%s", saveNotif.Text))
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/3000))
	}

}

func drawGameStones(win *pixelgl.Window, sprites Sprite) {
	var stone *pixel.Sprite
	var y, x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			if state.Board[y][x] != 0 {
				if state.Board[y][x] == 1 {
					stone = sprites.BlackStone
				} else {
					stone = sprites.WhiteStone
				}
				if showCapturable && state.CapturableBoard[y][x] == 1 {
					drawStone(win, x, y, sprites.BlueStone)
				} else if showCapturable && state.CapturableBoard[y][x] == 2 {
					drawStone(win, x, y, sprites.RedStone)
				} else {
					drawStone(win, x, y, stone)
				}
			}
		}
	}
	if hinter.Draw {
		if time.Now().Sub(hinter.Timestamp).Seconds() > 3.0 || state.Board[hinter.Y][hinter.X] != 0 {
			hinter.Draw = false
		} else {
			drawStone(win, hinter.X, hinter.Y, sprites.Goldstone)
		}
	}
}

func drawWin(win *pixelgl.Window, sprites Sprite) {
	asset := sprites.WinIA
	sound := sounds.WinIA
	if !state.IA {
		if state.WinOne {
			asset = sprites.WinPlayerOne
		} else {
			asset = sprites.WinPlayerTwo
		}
		sound = sounds.WinPlayer
	} else {
		if state.WinTwo {
			asset = sprites.WinPlayerTwo
			sound = sounds.Winvsia
		}
	}
	if mycfg.Sound {
		go playSound(sound)
	}
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V((X/3518)*0.46875,
		(Y/1628)*0.3259259))
	asset.Draw(win, mat)
}

func drawInfo(win *pixelgl.Window, sprite Sprite) {
	time := make(chan bool)
	go func() {
		drawTimer(win, sprite)
		time <- true
	}()

	vector := pixel.Vec{X: X * 0.2, Y: Y * 0.71}
	txt := text.New(vector, sprite.Atlas)
	txt.Color = colornames.Black
	txt.WriteString(fmt.Sprintf("%s", strconv.Itoa((state.Turn / 2))))
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/2750))

	vector = pixel.Vec{X: X * 0.2, Y: Y * 0.605}
	txt = text.New(vector, sprite.Atlas)
	txt.Color = colornames.Black
	txt.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(state.CapturePlayerOne))))
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/2750))
	vector = pixel.Vec{X: X * 0.2, Y: Y * 0.5}
	txt = text.New(vector, sprite.Atlas)
	txt.Color = colornames.Black
	txt.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(state.CapturePlayerTwo))))
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/2750))
	<-time
}

func drawTimer(win *pixelgl.Window, sprite Sprite) {
	vector := pixel.Vec{X: X * 0.1729167, Y: Y - Y*0.097222}
	mat := pixel.IM
	mat = mat.Moved(vector)
	mat = mat.ScaledXY(vector, pixel.V((X/248)*0.1291666666, (Y/144)*0.13333333))
	sprite.Wood.Draw(win, mat)
	vector = pixel.Vec{X: X * 0.13, Y: Y * 0.92}
	txt := text.New(vector, sprite.Atlas)
	txt.Color = colornames.Black
	if tt := time.Now().Sub(timer); tt >= time.Second {
		str := regexMatch(tt.String(), "(.*)[.]")
		txt.WriteString(fmt.Sprintf("%s", str[1]+"s"))
	} else {
		txt.WriteString(fmt.Sprintf("0s"))
	}

	txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/2750))
	vector = pixel.Vec{X: X * 0.13, Y: Y * 0.82}
	txt = text.New(vector, sprite.Atlas)
	txt.Color = colornames.Black
	if IATimer == 0 {
		txt.WriteString(fmt.Sprintf("    ms"))
	} else {
		txt.WriteString(fmt.Sprintf("%d ms", int(IATimer.Seconds()*1000)))
	}
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, (X+Y)/2750))
}

func drawBoard(win *pixelgl.Window, sprites Sprite) {
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V((X/1920.0), (Y/1080)))
	sprites.Board.Draw(win, mat)
	drawUnstatic(win, sprites)
	drawGameStones(win, sprites)
	if saveNotif.Draw {
		drawNotif(win, sprites)
	}
	if state.WinOne || state.WinTwo {
		drawWin(win, sprites)
	}
}
