package main

import (
	"math"
	"time"
)

func scoreAlignement2(align alignementHeuristic, board *[19][19]int8, player int8) int {
	score := 0
	var dx, dy int8
	dx = align.End[0] - align.Start[0]
	dy = align.End[1] - align.Start[1]

	dx, dy = getDir(dx, dy)

	if b, blocked := alignementIsInRange(&align, board, dx, dy); b {

		if align.Length == 2 {
			score += 0
		} else if align.Length == 3 {
			score += 30
			score += blocked * 5
		} else if align.Length == 4 {
			score += 100
			score += blocked * 10
		} else if align.Length >= 5 {
			score += 1500
			score += blocked * 10
		}

	}

	if player != align.Player {
		score *= -1
	}

	return score
}

func heuristicAlignement2(board *[19][19]int8, stonesPlayed *[][2]int8, x, y, player int8) int {
	score := 0
	count := 0
	alignements := make([]alignementHeuristic, len(*stonesPlayed)*2)
	for _, pos := range *stonesPlayed {

		if board[pos[1]][pos[0]] == 0 {
			continue
		}

		currentP := board[pos[1]][pos[0]]

		getAlignementProps(board, pos[0], pos[1], currentP, &alignements, &count)
	}

	for _, align := range alignements {
		if align.Player == 0 {
			break
		}
		score += scoreAlignement2(align, board, player)
	}

	return score
}

func heuristicCapture2(rootNode *Node) int {
	score := 0
	captureP1, captureP2 := countCapturable(&rootNode.CapturableBoard)
	if rootNode.CurrentPlayer == 1 {
		score += (10 * int(rootNode.CapturePlayerOne))
		score += (-10 * int(rootNode.CapturePlayerTwo))
		if rootNode.CapturePlayerOne == 4 && captureP1 > 1 {
			score *= 5
		} else if rootNode.CapturePlayerTwo == 4 && captureP2 > 1 {
			score /= 5
		}
		score += captureP1*100 - captureP2*100

	} else {
		score += (-10 * int(rootNode.CapturePlayerOne))
		score += (10 * int(rootNode.CapturePlayerTwo))
		if rootNode.CapturePlayerOne == 4 && captureP1 > 1 {
			score /= 5
		} else if rootNode.CapturePlayerTwo == 4 && captureP2 > 1 {
			score *= 5
		}
		score += captureP2*100 - captureP1*100
	}
	return score
}

func heuristic2(rootNode *Node, depth int) int {
	score := 0

	if state.Capture {
		score += heuristicCapture2(rootNode)
	}

	score += heuristicAlignement2(&rootNode.Board, &rootNode.StonesPlayed, rootNode.Pos[0], rootNode.Pos[1], rootNode.CurrentPlayer)

	if rootNode.Cut {
		score = 1000000
		score *= depth + 1
		return score
	}

	return score
}

func minimax2(rootNode *Node, depth int, alpha, beta int, maximizing bool) int {
	if depth == 0 || rootNode.Cut {
		return heuristic2(rootNode, depth)
	}
	if maximizing {
		maxEval := -1000000
		for _, child := range rootNode.Children {
			eval := minimax(child, depth-1, alpha, beta, false)
			maxEval = int(math.Max(float64(maxEval), float64(eval)))
			alpha = int(math.Max(float64(alpha), float64(eval)))
			child.Alpha = alpha
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := 1000000
		for _, child := range rootNode.Children {
			eval := minimax(child, depth-1, alpha, beta, true)
			minEval = int(math.Min(float64(minEval), float64(eval)))
			beta = int(math.Min(float64(beta), float64(eval)))
			child.Beta = beta
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func algo2(rootNode *Node) (int8, int8) {
	value := minimax2(rootNode, deepth, -1000000, 1000000, true)
	for _, elem := range rootNode.Children {
		if elem.Alpha == value {
			x := elem.Pos[0]
			y := elem.Pos[1]
			return x, y
		}
	}
	return 0, 0
}

func buildChildP2(rootNode *Node, move [2]int8, k, i int, c chan bool, d chan bool, w /*, z*/ int) {

	newNode := Node{}
	newNode.Pos = move
	newNode.Alpha = -1000000
	newNode.Beta = 1000000
	newNode.CurrentPlayer = getEnnemy(rootNode.CurrentPlayer)
	newNode.Board = rootNode.Board
	newNode.Board[move[1]][move[0]] = newNode.CurrentPlayer
	newNode.ThreeBoard = doubleThree(&newNode.Board)
	newNode.Cut = false
	newNode.CapturePlayerOne = rootNode.CapturePlayerOne
	newNode.CapturePlayerTwo = rootNode.CapturePlayerTwo

	if newNode.CurrentPlayer == 1 {
		newNode.CapturePlayerOne += isCaptured(&newNode.Board, newNode.Pos[0], newNode.Pos[1], newNode.CurrentPlayer)
	} else {
		newNode.CapturePlayerTwo += isCaptured(&newNode.Board, newNode.Pos[0], newNode.Pos[1], newNode.CurrentPlayer)
	}
	newNode.CapturableBoard = fcb(&newNode.Board, &newNode.ThreeBoard)
	lenStones := len(rootNode.StonesPlayed)
	newNode.StonesPlayed = make([][2]int8, lenStones+1)
	for i := 0; i < lenStones; i++ {
		newNode.StonesPlayed[i] = rootNode.StonesPlayed[i]
	}
	newNode.StonesPlayed[lenStones][0] = move[0]
	newNode.StonesPlayed[lenStones][1] = move[1]

	win := computeWinInBuild(&newNode, i)

	if win {
		newNode.Cut = true
		rootNode.Children = append(rootNode.Children, &newNode)

		winPos = newNode.Pos

		d <- true

		c <- true
		return
	}

	loose := computeLooseInBuild(rootNode, move, i)

	if loose {
		loosePos = newNode.Pos
	}

	rootNode.Children = append(rootNode.Children, &newNode)

	d <- true

	buildTree2(rootNode.Children[k], i+1)
	c <- true
	return
}

func buildChild2(rootNode *Node, move *[2]int8, k, i /*, z*/ int) {
	newNode := Node{}
	newNode.Pos = *move
	newNode.Alpha = -1000000
	newNode.Beta = 1000000
	newNode.CurrentPlayer = getEnnemy(rootNode.CurrentPlayer)
	newNode.Board = rootNode.Board
	newNode.Board[move[1]][move[0]] = newNode.CurrentPlayer
	newNode.ThreeBoard = doubleThree(&newNode.Board)
	newNode.Cut = false

	newNode.CapturePlayerOne = rootNode.CapturePlayerOne
	newNode.CapturePlayerTwo = rootNode.CapturePlayerTwo
	if newNode.CurrentPlayer == 1 {
		newNode.CapturePlayerOne += isCaptured(&newNode.Board, newNode.Pos[0], newNode.Pos[1], newNode.CurrentPlayer)
	} else {
		newNode.CapturePlayerTwo += isCaptured(&newNode.Board, newNode.Pos[0], newNode.Pos[1], newNode.CurrentPlayer)
	}
	newNode.CapturableBoard = fcb(&newNode.Board, &newNode.ThreeBoard)
	lenStones := len(rootNode.StonesPlayed)
	newNode.StonesPlayed = make([][2]int8, lenStones+1)
	for i := 0; i < lenStones; i++ {
		newNode.StonesPlayed[i] = rootNode.StonesPlayed[i]
	}
	newNode.StonesPlayed[lenStones][0] = (*move)[0]
	newNode.StonesPlayed[lenStones][1] = (*move)[1]

	win := computeWinInBuild(&newNode, i)
	if win {
		newNode.Cut = true
		rootNode.Children = append(rootNode.Children, &newNode)
		return
	}
	rootNode.Children = append(rootNode.Children, &newNode)
	buildTree2(rootNode.Children[k], i+1)
	return
}

func buildTreeP2(rootNode *Node, i int) {
	a := make(chan bool)
	b := make(chan bool)

	listMoves := [][2]int8{}
	tmpBoard := rootNode.Board

	if state.Turn+i == 3 && (state.Long || state.Pro) {
		if state.Pro {
			listMoves = getMovesLongPro()
		} else {
			listMoves = getMovesLong()
		}
	} else {
		getValidMoves(&tmpBoard, &rootNode.ThreeBoard, &listMoves, &rootNode.StonesPlayed)
	}
	getBestMoves2(rootNode.CurrentPlayer, &listMoves, &rootNode.Board, &rootNode.ThreeBoard, &rootNode.CapturableBoard)

	w := width

	if width > len(listMoves) {
		w = len(listMoves)
	}
	for k, move := range listMoves[0:w] {
		go buildChildP2(rootNode, move, k, i, a, b, w)
		<-b
	}
	for j := 0; j < w; j++ {
		<-a
	}
	return
}

func computeMyScore2(score int, nb int8) int {
	if nb == 1 {
		score += 100
	} else if nb == 2 {
		score += 600
	} else if nb == 3 {
		score += 1500
	} else if nb >= 4 {
		score += 10000
	}

	return score
}

func computeEnnemyScore2(score int, nb int8) int {
	if nb == 2 {
		score += 100
	} else if nb == 3 {
		score += 800
	} else if nb == 4 {
		score += 5000
	} else if nb >= 5 {
		score += 20000
	}

	return score
}

func lengthOfAlignementScore2(px, py, player int8, board *[19][19]int8, score int, ennemy bool) (int8, int) {
	var alignementCount int8
	alignementCount = 0

	if ennemy {
		if nb := checkVerticalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore2(score, nb)
		}
		if nb := checkHorizontalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore2(score, nb)
		}
		if nb := checkDiagRAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore2(score, nb)
		}
		if nb := checkDiagLAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore2(score, nb)
		}
	} else {
		if nb := checkVerticalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore2(score, nb)
		}
		if nb := checkHorizontalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore2(score, nb)
		}
		if nb := checkDiagRAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore2(score, nb)
		}
		if nb := checkDiagLAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore2(score, nb)
		}
	}

	return alignementCount, score
}

func getScore2(pos *[2]int8, player int8, Board, ThreeBoard, CapturableBoard *[19][19]int8) int {
	score := 0

	newPlayer := getEnnemy(player)

	if state.Capture {
		if newPlayer == 1 {
			// si je me protège
			if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%(newPlayer+3) == 0 {
				score += 1000

				// si j'attaque
			} else if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%5 == 0 {
				score += 1500
			}
		} else {
			if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%(newPlayer+3) == 0 {
				score += 1000
			} else if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%4 == 0 {
				score += 1500
			}
		}
	}

	if alignementCount, scoreTmp := lengthOfAlignementScore2(pos[0], pos[1], newPlayer, Board, score, false); alignementCount > 0 {
		score = scoreTmp
	}

	if alignementCount, scoreTmp := lengthOfAlignementScore2(pos[0], pos[1], getEnnemy(newPlayer), Board, score, true); alignementCount > 0 {
		score = scoreTmp
	}

	return score
}

func getBestMoves2(player int8, listMoves *[][2]int8, Board, ThreeBoard, CapturableBoard *[19][19]int8) {
	scoreTab := make([]int, len(*listMoves))

	for i, elem := range *listMoves {
		scoreTab[i] = getScore2(&elem, player, Board, ThreeBoard, CapturableBoard)
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

func buildTree2(rootNode *Node, i int) {
	if i < deepth {
		listMoves := [][2]int8{}
		tmpBoard := rootNode.Board

		if state.Turn+i == 3 && (state.Long || state.Pro) {
			if state.Pro {
				listMoves = getMovesLongPro()
			} else {
				listMoves = getMovesLong()
			}
		} else {
			getValidMoves(&tmpBoard, &rootNode.ThreeBoard, &listMoves, &rootNode.StonesPlayed)
		}
		getBestMoves2(rootNode.CurrentPlayer, &listMoves, &rootNode.Board, &rootNode.ThreeBoard, &rootNode.CapturableBoard)

		w := width

		if width > len(listMoves) {
			w = len(listMoves)
		}
		for k, move := range listMoves[0:w] {
			buildChild2(rootNode, &move, k, i /* , k+1+((rootNode.NodeWidth-1)*width) */)
		}
	}
	return
}

func buildWorker2(rootNode *Node, i int) {
	if routine {
		buildTreeP2(rootNode, i)
	} else {
		buildTree2(rootNode, i)
	}
	return
}

func playIA2(turn int) int {
	winPos = [2]int8{-1, -1}
	loosePos = [2]int8{-1, -1}
	if turn == 1 {
		PutStone(9, 9)
		return gameRules(turn)
	}
	t1 := time.Now()
	rootNode := &Node{Pos: [2]int8{-1, -1}, Alpha: -1000000, Beta: 1000000, Board: state.Board, ThreeBoard: state.ThreeBoard, CapturableBoard: state.CapturableBoard, StonesPlayed: state.StonesPlayed, CurrentPlayer: getEnnemy(state.CurrentPlayer), Cut: false}
	buildWorker2(rootNode, 0)
	if winPos[0] != -1 {
		PutStone(winPos[0], winPos[1])
		t2 := time.Now()
		IATimer = time.Now().Sub(t1) + time.Now().Sub(t2)
		return gameRules(turn)
	}

	if loosePos[0] != -1 {
		PutStone(loosePos[0], loosePos[1])
		t2 := time.Now()
		IATimer = time.Now().Sub(t1) + time.Now().Sub(t2)
		return gameRules(turn)
	}
	x, y := algo2(rootNode)
	t2 := time.Now()
	IATimer = time.Now().Sub(t1) + time.Now().Sub(t2)
	PutStone(x, y)
	return gameRules(turn)
}
