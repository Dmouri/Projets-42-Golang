package main

func assignValidMoves(validMoves, ThreeBoard *[19][19]int8, listMoves *[][2]int8, x, y int8) {
	if isInRange(x-1, y+1) && validMoves[y+1][x-1] == 0 && ThreeBoard[y+1][x-1] == 0 {
		validMoves[y+1][x-1] = 3
		*listMoves = append(*listMoves, [2]int8{x - 1, y + 1})
	}
	if isInRange(x, y+1) && validMoves[y+1][x] == 0 && ThreeBoard[y+1][x] == 0 {
		validMoves[y+1][x] = 3
		*listMoves = append(*listMoves, [2]int8{x, y + 1})
	}
	if isInRange(x+1, y+1) && validMoves[y+1][x+1] == 0 && ThreeBoard[y+1][x+1] == 0 {
		validMoves[y+1][x+1] = 3
		*listMoves = append(*listMoves, [2]int8{x + 1, y + 1})
	}
	if isInRange(x+1, y) && validMoves[y][x+1] == 0 && ThreeBoard[y][x+1] == 0 {
		validMoves[y][x+1] = 3
		*listMoves = append(*listMoves, [2]int8{x + 1, y})
	}
	if isInRange(x+1, y-1) && validMoves[y-1][x+1] == 0 && ThreeBoard[y-1][x+1] == 0 {
		validMoves[y-1][x+1] = 3
		*listMoves = append(*listMoves, [2]int8{x + 1, y - 1})
	}
	if isInRange(x, y-1) && validMoves[y-1][x] == 0 && ThreeBoard[y-1][x] == 0 {
		validMoves[y-1][x] = 3
		*listMoves = append(*listMoves, [2]int8{x, y - 1})
	}
	if isInRange(x-1, y-1) && validMoves[y-1][x-1] == 0 && ThreeBoard[y-1][x-1] == 0 {
		validMoves[y-1][x-1] = 3
		*listMoves = append(*listMoves, [2]int8{x - 1, y - 1})
	}
	if isInRange(x-1, y) && validMoves[y][x-1] == 0 && ThreeBoard[y][x-1] == 0 {
		validMoves[y][x-1] = 3
		*listMoves = append(*listMoves, [2]int8{x - 1, y})
	}
}

func getValidMoves(validMoves, ThreeBoard *[19][19]int8, listMoves, StonesPlayed *[][2]int8) {

	for _, elem := range *StonesPlayed {
		if validMoves[elem[1]][elem[0]] == 1 || validMoves[elem[1]][elem[0]] == 2 {
			assignValidMoves(validMoves, ThreeBoard, listMoves, elem[0], elem[1])
		}
	}
}

func swapIntTab(tab *[][2]int8, src, dest int) {
	var tmp = [2]int8{(*tab)[src][0], (*tab)[src][1]}
	(*tab)[src] = (*tab)[dest]
	(*tab)[dest] = tmp
}

func swapInt(tab *[]int, src, dest int) {
	tmp := (*tab)[src]
	(*tab)[src] = (*tab)[dest]
	(*tab)[dest] = tmp
}

func getBestMoves(player int8, listMoves *[][2]int8, Board, ThreeBoard, CapturableBoard *[19][19]int8) {
	scoreTab := make([]int, len(*listMoves))

	for i, elem := range *listMoves {
		scoreTab[i] = getScore(&elem, player, Board, ThreeBoard, CapturableBoard)
	}

	var first int
	length := len(scoreTab)
	for i := length - 1; i > 0; i-- {
		first = 0
		for j := 1; j <= i; j++ {
			if scoreTab[j] < scoreTab[first] {
				first = j
			}
		}
		swapInt(&scoreTab, first, i)
		swapIntTab(listMoves, first, i)
	}
}
