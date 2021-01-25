package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep/speaker"

	"github.com/faiface/pixel"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel/pixelgl"
)

func graphicHook(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeftControl) ||
		win.Pressed(pixelgl.KeyRightControl) {
		if win.JustPressed(pixelgl.KeyF) {
			if mycfg.Fullscreen == false {
				win.SetMonitor(pixelgl.PrimaryMonitor())
				win.Clear(colornames.Black)
				mycfg.Fullscreen = true
				redraw = true
			} else if mycfg.Fullscreen == true {
				win.SetMonitor(nil)
				win.Clear(colornames.Black)
				mycfg.Fullscreen = false
				redraw = true
			}
		}
		if win.JustPressed(pixelgl.KeyV) {
			if mycfg.Vsync == false {
				win.SetVSync(true)
				win.Clear(colornames.Black)
				mycfg.Vsync = true
				redraw = true
			} else if mycfg.Vsync == true {
				win.SetVSync(false)
				win.Clear(colornames.Black)
				mycfg.Vsync = false
				redraw = true
			}
		}

		if win.JustPressed(pixelgl.KeyA) {
			if mycfg.Smooth == false {
				win.SetSmooth(true)
				win.Clear(colornames.Black)
				mycfg.Smooth = true
				redraw = true
			} else if mycfg.Smooth == true {
				win.SetSmooth(false)
				win.Clear(colornames.Black)
				mycfg.Smooth = false
				redraw = true
			}
		}
	}
}

func hookMenu(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.MouseButton1) {
		mouseX := win.MousePosition().X
		mouseY := win.MousePosition().Y
		if mouseX > X*0.058 && mouseX < X*0.229 && mouseY < Y-Y*0.32 && mouseY > Y-Y*0.461 {
			state.IA = true
			state.TwoP = false
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.058 && mouseX < X*0.229 && mouseY < Y-Y*0.524 && mouseY > Y-Y*0.664 {
			state.IA = false
			state.TwoP = true
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.397 && mouseX < X*0.568 && mouseY < Y-Y*0.32 && mouseY > Y-Y*0.461 {
			if state.Capture {
				state.Capture = false
			} else {
				state.Capture = true
			}
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.397 && mouseX < X*0.568 && mouseY < Y-Y*0.524 && mouseY > Y-Y*0.664 {
			if state.DoubleThree {
				state.DoubleThree = false
			} else {
				state.DoubleThree = true
			}
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.733 && mouseX < X*0.904 && mouseY < Y-Y*0.32 && mouseY > Y-Y*0.461 {
			state.Standard = true
			state.Long = false
			state.Pro = false
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.733 && mouseX < X*0.904 && mouseY < Y-Y*0.524 && mouseY > Y-Y*0.664 {
			state.Standard = false
			state.Long = false
			state.Pro = true
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.733 && mouseX < X*0.904 && mouseY < Y-Y*0.722 && mouseY > Y-Y*0.863 {
			state.Standard = false
			state.Long = true
			state.Pro = false
			redraw = true
			if mycfg.Sound == true {
				go playSound(sounds.Stone)
			}
		} else if mouseX > X*0.393 && mouseX < X*0.607 && mouseY < Y-Y*0.837 && mouseY > Y-Y*0.956 {
			lastTurnState = []gameState{state}
			timer = time.Now()
			redraw = true
			state.Menu = false
		} else if mouseX > X*0.059 && mouseX < X*0.273 && mouseY < Y-Y*0.837 && mouseY > Y-Y*0.956 {
			hookLoad()
			redraw = true
		} else if mouseX > X*0.928 && mouseX < X*0.971 && mouseY < Y-Y*0.05 && mouseY > Y-Y*0.126 {
			hookSound()
		}
	}
}

func hookLoad() {
	if _, err := os.Stat("./save/save.gok"); os.IsNotExist(err) {
		fmt.Printf("No save found.\n")
		return
	}

	file, err := os.Open("./save/save.gok")
	if err != nil {
		fmt.Print(err)
		return
	}
	dec := json.NewDecoder(file)
	var tmpSave saveGame
	err = dec.Decode(&tmpSave)

	if err != nil {
		fmt.Printf("Unable to load save: %v\n", err)
		return
	}
	timer = tmpSave.Timer
	lastTurnState = tmpSave.LastTurnState
	state = tmpSave.State
	fmt.Printf("Save loaded successfully.\n")
	return
}

func hookReset() {
	state.Board = [19][19]int8{}
	redraw = true
	timer = time.Now()
	state = gameState{false, state.IA, state.TwoP, state.Capture, state.DoubleThree, state.Standard, state.Pro, state.Long, [19][19]int8{},
		[19][19]int8{}, [19][19]int8{}, 1, 1, [2]int8{-1, -1}, 0, 0, false, false, false,
		false, []winInProgress{}, [][2]int8{}}
	IATimer = 0
}

func undoHook() {
	if len(lastTurnState) > 1 {
		state = lastTurnState[len(lastTurnState)-1]
		lastTurnState = lastTurnState[:len(lastTurnState)-1]
		IATimer = 0
		redraw = true
	}
}

func hookSound() {
	if mycfg.Sound == true {
		speaker.Lock()
		sounds.Ctrl.Paused = true
		speaker.Unlock()
		mycfg.Sound = false
	} else {
		speaker.Lock()
		sounds.Ctrl.Paused = false
		speaker.Unlock()
		mycfg.Sound = true
	}
	redraw = true
}

func hookHint(win *pixelgl.Window, sprites Sprite) {
	rootNode := &Node{Pos: [2]int8{-1, -1}, Alpha: -1000000, Beta: 1000000, Board: state.Board, ThreeBoard: state.ThreeBoard, CapturableBoard: state.CapturableBoard, StonesPlayed: state.StonesPlayed, CurrentPlayer: getEnnemy(state.CurrentPlayer)}
	buildTree(rootNode, 0)
	x, y := algo(rootNode)
	hinter = hintInfo{true, x, y, time.Now()}
}

func hookSave() {

	if _, err := os.Stat("./save"); os.IsNotExist(err) {
		os.Mkdir("."+string(filepath.Separator)+"save", 0777)
	}

	file, err := os.Create("./save/save.gok")
	if err != nil {
		fmt.Print(err)
		return
	}
	enc := json.NewEncoder(file)
	save := saveGame{state, timer, lastTurnState}
	err = enc.Encode(save)
	if err != nil {
		fmt.Print(err)
	}
	file.Close()
}

func hookOptions(win *pixelgl.Window, sprites Sprite) {
	if win.JustPressed(pixelgl.MouseButton1) {
		mouseX := win.MousePosition().X
		mouseY := win.MousePosition().Y
		if mouseX > X*0.0137586 && mouseX < X*0.102785 &&
			mouseY > Y*0.302901 && mouseY < Y*0.37402685 {
			undoHook()
		} else if mouseX > X*0.130097 && mouseX < X*0.219215 &&
			mouseY > Y*0.3019649 && mouseY < Y*0.37402685 {
			hookReset()
		} else if mouseX > X*0.0749097 && mouseX < X*0.164138 &&
			mouseY > Y*0.1416959 && mouseY < Y*0.212943 &&
			(!state.WinOne && !state.WinTwo) && state.Turn > 1 {
			hookHint(win, sprites)
			redraw = true
		} else if mouseX > X*0.787604 && mouseX < X*0.9691 &&
			mouseY > Y*0.814936 && mouseY < Y*0.93291 {
			// MENU
			state = gameState{true, true, false, true, true, true, false, false, [19][19]int8{}, [19][19]int8{}, [19][19]int8{}, 1, 1, [2]int8{-1, -1}, 0, 0, false, false, false, false, []winInProgress{}, [][2]int8{}}
			redraw = true
			// reset counter, list undo
		} else if mouseX > X*0.787757 && mouseX < X*0.96908 &&
			mouseY > Y*0.64186 && mouseY < Y*0.75865 &&
			(!state.WinOne && !state.WinTwo) {
			hookSave()
			saveNotif = notif{true, "Game saved", time.Now()}
			redraw = true
		} else if mouseX > X*0.787757 && mouseX < X*0.96908 &&
			mouseY > Y*0.46552 && mouseY < Y*0.58411 {
			hookSound()
		}
	}
}

func calculatePosition(mouseX, mouseY float64) (int8, int8) {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	var xsq float64
	var ysq float64
	var px int8
	var py int8
	x1 = 0.265278 * X
	x2 = 0.7340278 * X
	y1 = 0.079532 * Y
	y2 = 0.91462 * Y
	xsq = (x2 - x1) / 18
	ysq = (y2 - y1) / 18
	px = int8(math.Round((mouseX - x1) / xsq))
	py = int8(math.Round((mouseY - y1) / ysq))
	if px > 18 {
		px = 18
	}
	if py > 18 {
		py = 18
	}
	return px, py
}

func gameHook(win *pixelgl.Window, sprites Sprite) {
	mouseX := win.MousePosition().X
	mouseY := win.MousePosition().Y
	var transparencyStone *pixel.Sprite
	if mouseX > X*0.25427 && mouseX < X*0.747730 && mouseY > Y*0.05991 &&
		mouseY < Y*0.9357 && (!state.WinOne && !state.WinTwo) {
		px, py := calculatePosition(mouseX, mouseY)
		if (state.Long || state.Pro) && state.Turn == 3 {
			nb := 0
			if state.Pro {
				nb += 1
			}
			if (px >= int8(6+nb) && px <= int8(12-nb)) && (py >= int8(6+nb) && py <= int8(12-nb)) {
				return
			}
		}
		if win.JustPressed(pixelgl.MouseButton1) {
			if state.Board[py][px] == 0 && (state.ThreeBoard[py][px] != state.CurrentPlayer && state.ThreeBoard[py][px] != 3) && validPosition(px, py) {
				hinter.Draw = false
				PutStone(px, py)
			}
			redraw = true
		}
		if state.Board[py][px] != 0 && transpaDrawn {
			redraw = true
		}
		if (state.Board[py][px] == 0 && (state.ThreeBoard[py][px] == state.CurrentPlayer || state.ThreeBoard[py][px] == 3)) || !validPosition(px, py) {
			if lastpx != px || lastpy != py {
				redraw = true
			}
			lastpx = px
			lastpy = py
		}
		if state.Board[py][px] == 0 && (state.ThreeBoard[py][px] != state.CurrentPlayer && state.ThreeBoard[py][px] != 3) && validPosition(px, py) {
			if lastpx != px || lastpy != py {
				redraw = true
			}
			lastpx = px
			lastpy = py
			if state.CurrentPlayer == 1 {
				transparencyStone = sprites.TransparencyBlackStone
			} else {
				transparencyStone = sprites.TransparencyWhiteStone
			}
			drawStone(win, px, py, transparencyStone)
			transpaDrawn = true
		} else {
			transpaDrawn = false
		}
	} else if lastpx != -1 && lastpy != -1 && (!state.WinOne && !state.WinTwo) {
		redraw = true
		lastpx = -1
		lastpy = -1
	}
}

func ctrlHook(win *pixelgl.Window, sprites Sprite) {
	if win.Pressed(pixelgl.KeyLeftControl) ||
		win.Pressed(pixelgl.KeyRightControl) {
		if win.JustPressed(pixelgl.KeyR) {
			hookReset()
		}
		if win.JustPressed(pixelgl.KeyC) {
			if showCapturable {
				showCapturable = false
			} else {
				showCapturable = true
			}
			redraw = true
		}
		if win.JustPressed(pixelgl.KeyS) {
			hookSave()
			saveNotif = notif{true, "Game saved", time.Now()}
			redraw = true
		}
		if win.JustPressed(pixelgl.KeyZ) {
			undoHook()
		}

		if win.JustPressed(pixelgl.KeyK) {
			if state.IA {
				state.IA = false
			} else {
				state.IA = true
			}
		}

		if win.JustPressed(pixelgl.KeyH) {
			hookHint(win, sprites)
			redraw = true
		}
	}
}

func treeHook(win *pixelgl.Window) {

	if !state.WinOne && !state.WinTwo {
		if win.JustPressed(pixelgl.KeyKPAdd) {
			deepth += 1
			fmt.Printf("Tree Size updated: %d %d\n", deepth, width)
		} else if win.JustPressed(pixelgl.KeyKPSubtract) {
			if deepth > 2 {
				deepth -= 1
				fmt.Printf("Tree Size updated: %d %d\n", deepth, width)
			}
		} else if win.JustPressed(pixelgl.KeyKPDivide) {
			if width > 2 {
				width -= 1
				fmt.Printf("Tree Size updated: %d %d\n", deepth, width)
			}
		} else if win.JustPressed(pixelgl.KeyKPMultiply) {
			width += 1
			fmt.Printf("Tree Size updated: %d %d\n", deepth, width)
		}
	}
}

func hook(win *pixelgl.Window, sprites Sprite) {
	X = win.Bounds().Size().X
	Y = win.Bounds().Size().Y
	if lastX != X || lastY != Y {
		redraw = true
	}
	lastX = X
	lastY = Y
	if win.JustPressed(pixelgl.KeyEscape) {
		os.Exit(0)
	}
	if state.Menu {
		hookMenu(win)
	} else {
		hookOptions(win, sprites)
		gameHook(win, sprites)
		ctrlHook(win, sprites)
		treeHook(win)
	}
	graphicHook(win)
}
