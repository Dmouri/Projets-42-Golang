package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

func PutStone(px, py int8) {
	if state.Turn%2 == 0 {
		lastTurnState = append(lastTurnState, state)
	}
	state.Board[py][px] = state.CurrentPlayer
	state.LastPlayed = [2]int8{px, py}
	state.StonesPlayed = append(state.StonesPlayed, state.LastPlayed)
	state.Turn++
	if mycfg.Sound {
		go playSound(sounds.Stone)
	}
	state.CurrentPlayer = getEnnemy(state.CurrentPlayer)
}

func gameRules(turn int) int {
	if turn > state.Turn {
		turn = state.Turn
	}
	if turn != state.Turn {
		applyRules()
	}
	if turn != state.Turn && (state.WinConditionOne || state.WinConditionTwo) &&
		(!state.WinOne && !state.WinTwo) {
		mgtWinCondition()
	}
	return state.Turn
}

func showFpsAndTimer(win *pixelgl.Window, cfg pixelgl.WindowConfig, frames int, Sprites Sprite) int {
	select {
	case <-second:
		if !state.Menu {
			drawTimer(win, Sprites)
		}
		if !mycfg.Vsync {
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
		} else {
			win.SetTitle(fmt.Sprintf("%s | FPS: %g", cfg.Title, mycfg.RefreshRate))
		}
		return 0
	default:
	}
	return frames
}
