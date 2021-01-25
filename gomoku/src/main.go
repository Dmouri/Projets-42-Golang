package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	frames := 0
	cfg, win, Sprites := loadGame()
	turn := state.Turn
	go playAmbiance()
	for !win.Closed() {
		hook(win, Sprites)
		turn = gameRules(turn)
		if !state.WinOne && !state.WinTwo {
			if state.IA && !state.Menu && state.Turn%2 != 0 {
				turn = playIA(turn)
				redraw = true
				if iavsia {
					time.Sleep(500 * time.Millisecond)
				}
			} else if iavsia && state.IA && !state.Menu && state.Turn%2 == 0 {
				turn = playIA2(turn)
				redraw = true
				time.Sleep(500 * time.Millisecond)
			}
		}
		drawing(win, Sprites)
		win.Update()
		frames++
		frames = showFpsAndTimer(win, cfg, frames, Sprites)
	}
}

func mgtEntryParam() {
	r := false
	d := false
	w := false
	for k, vle := range os.Args {
		if vle == "-r" || vle == "--routine" {
			if r {
				log.Fatal(": Error: Duplicate parameters.\n")
				os.Exit(1)
			}
			routine = false
			r = true
		} else if vle == "-d" || vle == "--depth" {
			if d {
				log.Fatal(": Error: Duplicate parameters.\n")
				os.Exit(1)
			}
			d = true
			deepth = parseInt(k)
		} else if vle == "-w" || vle == "--width" {
			if w {
				log.Fatal(": Error: Duplicate parameters.\n")
				os.Exit(1)
			}
			w = true
			width = parseInt(k)
		} else if vle == "-ia" {
			iavsia = true
		}
	}
}

func main() {
	mycfg = config{false, 0, false, false, 1920, 1080, true}
	state = gameState{true, true, false, true, true, true, false, false, [19][19]int8{}, [19][19]int8{}, [19][19]int8{}, 1, 1,
		[2]int8{-1, -1}, 0, 0, false, false, false, false, []winInProgress{}, [][2]int8{}}
	redraw = true
	deepth, width, routine = 5, 8, true
	hinter = hintInfo{false, 0, 0, time.Now()}
	saveNotif = notif{false, "", time.Now()}
	showCapturable = false
	winPos = [2]int8{-1, -1}
	loosePos = [2]int8{-1, -1}

	if len(os.Args) != 0 {
		mgtEntryParam()
	}
	second = time.Tick(time.Second)
	pixelgl.Run(run)
}
