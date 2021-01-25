package main

import (
	"math"
)

func isCaptured(Board *[19][19]int8, px, py, player int8) int8 {
	counter := int8(0)
	ennemy := getEnnemy(player)
	if isCapturable(Board, px, py+3, player) {
		if Board[py+2][px] == ennemy && Board[py+1][px] == ennemy {
			counter++
			Board[py+2][px] = 0
			Board[py+1][px] = 0
		}
	}
	if isCapturable(Board, px+3, py+3, player) {
		if Board[py+2][px+2] == ennemy && Board[py+1][px+1] == ennemy {
			counter++
			Board[py+2][px+2] = 0
			Board[py+1][px+1] = 0
		}
	}
	if isCapturable(Board, px+3, py, player) {
		if Board[py][px+2] == ennemy && Board[py][px+1] == ennemy {
			counter++
			Board[py][px+2] = 0
			Board[py][px+1] = 0
		}
	}
	if isCapturable(Board, px+3, py-3, player) {
		if Board[py-2][px+2] == ennemy && Board[py-1][px+1] == ennemy {
			counter++
			Board[py-2][px+2] = 0
			Board[py-1][px+1] = 0
		}
	}
	if isCapturable(Board, px, py-3, player) {
		if Board[py-2][px] == ennemy && Board[py-1][px] == ennemy {
			counter++
			Board[py-2][px] = 0
			Board[py-1][px] = 0
		}
	}
	if isCapturable(Board, px-3, py-3, player) {
		if Board[py-2][px-2] == ennemy && Board[py-1][px-1] == ennemy {
			counter++
			Board[py-2][px-2] = 0
			Board[py-1][px-1] = 0
		}
	}
	if isCapturable(Board, px-3, py, player) {
		if Board[py][px-2] == ennemy && Board[py][px-1] == ennemy {
			counter++
			Board[py][px-2] = 0
			Board[py][px-1] = 0
		}
	}
	if isCapturable(Board, px-3, py+3, player) {
		if Board[py+2][px-2] == ennemy && Board[py+1][px-1] == ennemy {
			counter++
			Board[py+2][px-2] = 0
			Board[py+1][px-1] = 0
		}
	}
	return counter
}

func checkVertical(px, py, player int8, board *[19][19]int8) ([2]int8, [2]int8, bool) {
	count := 1
	var dy int8
	start, end := [2]int8{px, py}, [2]int8{px, py}
	for dy = -1; isInRange(px, py+dy); dy-- {
		if board[py+dy][px] == player {
			count++
			start = [2]int8{px, py + dy}
		} else {
			break
		}
	}
	for dy = 1; isInRange(px, py+dy); dy++ {
		if board[py+dy][px] == player {
			count++
			end = [2]int8{px, py + dy}
		} else {
			break
		}
	}
	if count >= 5 {
		return start, end, true
	}
	return start, end, false
}

func checkHorizontal(px, py, player int8, board *[19][19]int8) ([2]int8, [2]int8, bool) {
	count := 1
	var dx int8
	start, end := [2]int8{px, py}, [2]int8{px, py}
	for dx = -1; isInRange(px+dx, py); dx-- {
		if board[py][px+dx] == player {
			count++
			start = [2]int8{px + dx, py}
		} else {
			break
		}
	}
	for dx = 1; isInRange(px+dx, py); dx++ {
		if board[py][px+dx] == player {
			count++
			end = [2]int8{px + dx, py}
		} else {
			break
		}
	}
	if count >= 5 {
		return start, end, true
	}
	return start, end, false
}

func checkDiagR(px, py, player int8, board *[19][19]int8) ([2]int8, [2]int8, bool) {
	count := 1
	var dx int8
	var dy int8
	start, end := [2]int8{px, py}, [2]int8{px, py}
	for dx, dy = -1, -1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			count++
			start = [2]int8{px + dx, py + dy}
		} else {
			break
		}
		dy--
	}
	for dx, dy = 1, 1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			count++
			end = [2]int8{px + dx, py + dy}
		} else {
			break
		}
		dy++
	}
	if count >= 5 {
		return start, end, true
	}
	return start, end, false
}

func checkDiagL(px, py, player int8, board *[19][19]int8) ([2]int8, [2]int8, bool) {
	count := 1
	var dx int8
	var dy int8
	start, end := [2]int8{px, py}, [2]int8{px, py}
	for dx, dy = 1, -1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			count++
			start = [2]int8{px + dx, py + dy}
		} else {
			break
		}
		dy--
	}
	for dx, dy = -1, 1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			count++
			end = [2]int8{px + dx, py + dy}
		} else {
			break
		}
		dy++
	}
	if count >= 5 {
		return start, end, true
	}
	return start, end, false
}

func isAligned(px, py, player int8, board *[19][19]int8) ([][2]int8, [][2]int8, bool) {
	starttab := [][2]int8{}
	endtab := [][2]int8{}
	count := 0
	aligned := false
	if start, end, aligned := checkVertical(px, py, player, board); aligned {
		starttab = append(starttab, start)
		endtab = append(endtab, end)
		count++
	}
	if start, end, aligned := checkHorizontal(px, py, player, board); aligned {
		starttab = append(starttab, start)
		endtab = append(endtab, end)
		count++
	}
	if start, end, aligned := checkDiagR(px, py, player, board); aligned {
		starttab = append(starttab, start)
		endtab = append(endtab, end)
		count++
	}
	if start, end, aligned := checkDiagL(px, py, player, board); aligned {
		starttab = append(starttab, start)
		endtab = append(endtab, end)
		count++
	}
	if count != 0 {
		aligned = true
	}
	return starttab, endtab, aligned
}

func isSafeAlignement(start, end [2]int8, player int8, CapturableBoard *[19][19]int8) bool {
	dx, dy, length := getLengthAndDir(start, end)
	count := 0
	var i int8
	for i = 0; i < length; i++ {
		if count == 5 {
			return true
		}
		if CapturableBoard[start[1]+dy*i][start[0]+dx*i] == player {
			count = 0
		} else {
			count++
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func isAlignementFree(start, end [2]int8, player int8) bool {
	dx, dy, length := getLengthAndDir(start, end)
	count := 0
	var i int8
	for i = 0; i < length; i++ {
		if count == 5 {
			return true
		}
		if state.CapturableBoard[start[1]+dy*i][start[0]+dx*i] == player {
			count = 0
		} else {
			count++
		}
		if count == 5 {
			return true
		}
	}
	if player == 1 {
		state.WinConditionOne = true
	} else {
		state.WinConditionTwo = true
	}
	appendWinnable(start, end, player)
	return false
}

func fillCapturableBoard(Board, ThreeBoard *[19][19]int8) [19][19]int8 {
	CapturableBoard := [19][19]int8{}
	var x, y int8
	for x = 0; x < 19; x++ {
		for y = 0; y < 19; y++ {
			if stoneInSquare(x, y, Board) && isEatable(Board, ThreeBoard, x, y, Board[y][x]) {
				CapturableBoard[y][x] = Board[y][x]
			}
		}
	}
	return CapturableBoard
}

func fcb(Board, ThreeBoard *[19][19]int8) [19][19]int8 {
	CapturableBoard := [19][19]int8{}
	var x, y int8
	for x = 0; x < 19; x++ {
		for y = 0; y < 19; y++ {
			if stoneInSquare(x, y, Board) {
				b, x1, y1 := ie(Board, ThreeBoard, x, y, Board[y][x])
				if b {
					CapturableBoard[y][x] = Board[y][x]
					CapturableBoard[y1][x1] = 3 + Board[y][x]
				}
			}
		}
	}
	return CapturableBoard
}

func HasUnsafePosition(CapturableBoard [19][19]int8, player int8) bool {
	for x := 0; x < 19; x++ {
		for y := 0; y < 19; y++ {
			if CapturableBoard[y][x] == player {
				return true
			}
		}
	}
	return false
}

func counterCapture(CapturableBoard [19][19]int8, player int8) {
	if player == 2 {
		if state.CapturePlayerOne == 4 {
			unsafe := HasUnsafePosition(state.CapturableBoard, player)
			if unsafe {
				state.WinConditionTwo = true
			} else {
				state.WinTwo = true
			}
		} else {
			state.WinTwo = true
		}
	} else {
		if state.CapturePlayerTwo == 4 {
			unsafe := HasUnsafePosition(state.CapturableBoard, player)
			if unsafe {
				state.WinConditionOne = true
			} else {
				state.WinOne = true
			}
		} else {
			state.WinOne = true
		}
	}
}

func getLengthAndDir(start, end [2]int8) (int8, int8, int8) {
	var length int8
	dx := end[0] - start[0]
	dy := end[1] - start[1]
	if dx > 0 {
		dx = 1
	} else if dx < 0 {
		dx = -1
	}
	if dy > 0 {
		dy = 1
	} else if dy < 0 {
		dy = -1
	}
	sx := end[0] - start[0]
	sy := end[1] - start[1]
	if math.Abs(float64(sx)) >= math.Abs(float64(sy)) {
		length = int8(math.Abs(float64(sx))) + 1
	} else {
		length = int8(math.Abs(float64(sy))) + 1
	}
	return dx, dy, length
}

func applyRules() {
	player := state.Board[state.LastPlayed[1]][state.LastPlayed[0]]
	px := state.LastPlayed[0]
	py := state.LastPlayed[1]
	state.ThreeBoard = [19][19]int8{}
	if state.DoubleThree {
		state.ThreeBoard = doubleThree(&state.Board)
	}
	state.CapturableBoard = [19][19]int8{}
	if state.Capture {
		if nbCaptured := isCaptured(&state.Board, px, py, player); nbCaptured != 0 {
			if mycfg.Sound {
				go playSound(sounds.Capture)
			}
			if player == 1 {
				state.CapturePlayerOne += nbCaptured
			} else {
				state.CapturePlayerTwo += nbCaptured
			}
		}
		state.CapturableBoard = fcb(&state.Board, &state.ThreeBoard)
		if state.CapturePlayerOne >= 5 {
			state.WinOne = true
			return
		}
		if state.CapturePlayerTwo >= 5 {
			state.WinTwo = true
			return
		}
		if len(state.Winnable) > 0 {
			mgtWinnableAlignement()
		}
	}
	starttab, endtab, aligned := isAligned(px, py, player, &state.Board)
	if aligned {
		for i, start := range starttab {
			end := endtab[i]
			if isAlignementFree(start, end, player) {
				counterCapture(state.CapturableBoard, player)
			}
		}
	}
}

func validPosition(px, py int8) bool {
	if state.Standard {
		return true
	}
	if (state.Pro || state.Long) && state.Turn == 1 && (px != 9 || py != 9) {
		return false
	}
	if (state.Pro || state.Long) && state.Turn == 3 {
		var long int8
		long = 0
		if state.Long {
			long += 1
		}
		minX := 7 - long
		minY := 7 - long
		maxX := 11 + long
		maxY := 11 + long
		if px >= minX && px <= maxX && py >= minY && py <= maxY {
			return false
		}
	}
	return true
}
