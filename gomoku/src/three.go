package main

func doubleThreeByPosition(rootThreeBoard, rootBoard *[19][19]int8, pos *[2]int8) [19][19]int8 {

	ThreeBoard := *rootThreeBoard

	var x, y, dx, dy int8

	x = (*pos)[0] - 6
	y = (*pos)[1] - 6
	dx = (*pos)[0] + 6
	dy = (*pos)[1] + 6

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if dx > 19 {
		dx = 19
	}
	if dy > 19 {
		dy = 19
	}

	for ky := y; ky < dy; ky++ {
		for kx := x; kx < dx; kx++ {
			ThreeBoard[ky][kx] = 0
			count := 0
			if rootBoard[ky][kx] == 0 && stoneInTwoSquare(kx, ky, rootBoard) {
				count += checkThreeVert(rootBoard, kx, ky, 1)
				count += checkThreeHoriz(rootBoard, kx, ky, 1)
				count += checkThreeDiagR(rootBoard, kx, ky, 1)
				count += checkThreeDiagL(rootBoard, kx, ky, 1)
				if count >= 2 {
					ThreeBoard[ky][kx] += 1
				}
				count = 0
				count += checkThreeVert(rootBoard, kx, ky, 2)
				count += checkThreeHoriz(rootBoard, kx, ky, 2)
				count += checkThreeDiagR(rootBoard, kx, ky, 2)
				count += checkThreeDiagL(rootBoard, kx, ky, 2)
				if count >= 2 {
					ThreeBoard[ky][kx] += 2
				}
			}
		}
	}
	return ThreeBoard
}

func doubleThree(Board *[19][19]int8) [19][19]int8 {
	ThreeBoard := [19][19]int8{}
	var x, y int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			count := 0
			if Board[y][x] == 0 && stoneInTwoSquare(x, y, Board) {
				count += checkThreeVert(Board, x, y, 1)
				count += checkThreeHoriz(Board, x, y, 1)
				count += checkThreeDiagR(Board, x, y, 1)
				count += checkThreeDiagL(Board, x, y, 1)
				if count >= 2 {
					ThreeBoard[y][x] += 1
				}
				count = 0
				count += checkThreeVert(Board, x, y, 2)
				count += checkThreeHoriz(Board, x, y, 2)
				count += checkThreeDiagR(Board, x, y, 2)
				count += checkThreeDiagL(Board, x, y, 2)
				if count >= 2 {
					ThreeBoard[y][x] += 2
				}
			}
		}
	}
	return ThreeBoard
}

func stoneInTwoSquare(x, y int8, Board *[19][19]int8) bool {
	if isInRange(x-2, y+2) && Board[y+2][x-2] != 0 {
		return true
	}
	if isInRange(x-2, y+2) && Board[y+2][x-2] != 0 {
		return true
	}
	if isInRange(x+2, y+2) && Board[y+2][x+2] != 0 {
		return true
	}
	if isInRange(x+2, y) && Board[y][x+2] != 0 {
		return true
	}
	if isInRange(x+2, y-2) && Board[y-2][x+2] != 0 {
		return true
	}
	if isInRange(x, y-2) && Board[y-2][x] != 0 {
		return true
	}
	if isInRange(x-2, y-2) && Board[y-2][x-2] != 0 {
		return true
	}
	if isInRange(x-2, y) && Board[y][x-2] != 0 {
		return true
	}
	return false
}

func stoneInSquare(x, y int8, Board *[19][19]int8) bool {
	if isInRange(x-1, y+1) && Board[y+1][x-1] != 0 {
		return true
	}
	if isInRange(x, y+1) && Board[y+1][x] != 0 {
		return true
	}
	if isInRange(x+1, y+1) && Board[y+1][x+1] != 0 {
		return true
	}
	if isInRange(x+1, y) && Board[y][x+1] != 0 {
		return true
	}
	if isInRange(x+1, y-1) && Board[y-1][x+1] != 0 {
		return true
	}
	if isInRange(x, y-1) && Board[y-1][x] != 0 {
		return true
	}
	if isInRange(x-1, y-1) && Board[y-1][x-1] != 0 {
		return true
	}
	if isInRange(x-1, y) && Board[y][x-1] != 0 {
		return true
	}
	return false
}

func checkThreeVert(Board *[19][19]int8, x, y, player int8) int {
	if isInRange(x, y-1) && Board[y-1][x] == 0 && isInRange(x, y+3) && Board[y+3][x] == 0 && Board[y+1][x] == player && Board[y+2][x] == player {
		return 1
	} else if isInRange(x, y-1) && Board[y-1][x] == 0 && isInRange(x, y+4) && Board[y+4][x] == 0 && Board[y+1][x] == 0 && Board[y+2][x] == player && Board[y+3][x] == player {
		return 1
	} else if isInRange(x, y+1) && Board[y+1][x] == 0 && isInRange(x, y-3) && Board[y-3][x] == 0 && Board[y-1][x] == player && Board[y-2][x] == player {
		return 1
	} else if isInRange(x, y+1) && Board[y+1][x] == 0 && isInRange(x, y-4) && Board[y-4][x] == 0 && Board[y-1][x] == 0 && Board[y-2][x] == player && Board[y-3][x] == player {
		return 1
	}
	return 0
}

func checkThreeHoriz(Board *[19][19]int8, x, y, player int8) int {
	if isInRange(x-1, y) && Board[y][x-1] == 0 && isInRange(x+3, y) && Board[y][x+3] == 0 && Board[y][x+1] == player && Board[y][x+2] == player {
		return 1
	} else if isInRange(x-1, y) && Board[y][x-1] == 0 && isInRange(x+4, y) && Board[y][x+4] == 0 && Board[y][x+1] == 0 && Board[y][x+2] == player && Board[y][x+3] == player {
		return 1
	} else if isInRange(x+1, y) && Board[y][x+1] == 0 && isInRange(x-3, y) && Board[y][x-3] == 0 && Board[y][x-1] == player && Board[y][x-2] == player {
		return 1
	} else if isInRange(x+1, y) && Board[y][x+1] == 0 && isInRange(x-4, y) && Board[y][x-4] == 0 && Board[y][x-1] == 0 && Board[y][x-2] == player && Board[y][x-3] == player {
		return 1
	}
	return 0
}

func checkThreeDiagR(Board *[19][19]int8, x, y, player int8) int {
	if isInRange(x-1, y-1) && Board[y-1][x-1] == 0 && isInRange(x+3, y+3) && Board[y+3][x+3] == 0 && Board[y+1][x+1] == player && Board[y+2][x+2] == player {
		return 1
	} else if isInRange(x-1, y-1) && Board[y-1][x-1] == 0 && isInRange(x+4, y+4) && Board[y+4][x+4] == 0 && Board[y+1][x+1] == 0 && Board[y+2][x+2] == player && Board[y+3][x+3] == player {
		return 1
	} else if isInRange(x+1, y+1) && Board[y+1][x+1] == 0 && isInRange(x-3, y-3) && Board[y-3][x-3] == 0 && Board[y-1][x-1] == player && Board[y-2][x-2] == player {
		return 1
	} else if isInRange(x+1, y+1) && Board[y+1][x+1] == 0 && isInRange(x-4, y-4) && Board[y-4][x-4] == 0 && Board[y-1][x-1] == 0 && Board[y-2][x-2] == player && Board[y-3][x-3] == player {
		return 1
	}
	return 0
}

func checkThreeDiagL(Board *[19][19]int8, x, y, player int8) int {
	if isInRange(x+1, y-1) && Board[y-1][x+1] == 0 && isInRange(x-3, y+3) && Board[y+3][x-3] == 0 && Board[y+1][x-1] == player && Board[y+2][x-2] == player {
		return 1
	} else if isInRange(x+1, y-1) && Board[y-1][x+1] == 0 && isInRange(x-4, y+4) && Board[y+4][x-4] == 0 && Board[y+1][x-1] == 0 && Board[y+2][x-2] == player && Board[y+3][x-3] == player {
		return 1
	} else if isInRange(x-1, y+1) && Board[y+1][x-1] == 0 && isInRange(x+3, y-3) && Board[y-3][x+3] == 0 && Board[y-1][x+1] == player && Board[y-2][x+2] == player {
		return 1
	} else if isInRange(x-1, y+1) && Board[y+1][x-1] == 0 && isInRange(x+4, y-4) && Board[y-4][x+4] == 0 && Board[y-1][x+1] == 0 && Board[y-2][x+2] == player && Board[y-3][x+3] == player {
		return 1
	}
	return 0
}

func stoneInTwosSquare(x, y int8, Board [19][19]int8) bool {
	if (isInRange(x-1, y+1) && Board[y+1][x-1] != 0) || (isInRange(x-2, y+2) && Board[y+2][x-2] != 0) {
		return true
	}
	if (isInRange(x, y+1) && Board[y+1][x] != 0) || (isInRange(x-2, y+2) && Board[y+2][x-2] != 0) {
		return true
	}
	if (isInRange(x+1, y+1) && Board[y+1][x+1] != 0) || (isInRange(x+2, y+2) && Board[y+2][x+2] != 0) {
		return true
	}
	if (isInRange(x+1, y) && Board[y][x+1] != 0) || (isInRange(x+2, y) && Board[y][x+2] != 0) {
		return true
	}
	if (isInRange(x+1, y-1) && Board[y-1][x+1] != 0) || (isInRange(x+2, y-2) && Board[y-2][x+2] != 0) {
		return true
	}
	if (isInRange(x, y-1) && Board[y-1][x] != 0) || (isInRange(x, y-2) && Board[y-2][x] != 0) {
		return true
	}
	if (isInRange(x-1, y-1) && Board[y-1][x-1] != 0) || (isInRange(x-2, y-2) && Board[y-2][x-2] != 0) {
		return true
	}
	if (isInRange(x-1, y) && Board[y][x-1] != 0) || (isInRange(x-2, y) && Board[y][x-2] != 0) {
		return true
	}
	return false
}
