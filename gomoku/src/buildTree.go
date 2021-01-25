package main

import "sync"

func playerCanBeCaptured(player int8, CapturableBoard *[19][19]int8) bool {
	x, y := int8(0), int8(0)
	for x < 19 {
		y = 0
		for y < 19 {
			if (*CapturableBoard)[y][x] == player {
				return true
			}
			y++
		}
		x++
	}
	return false
}

func isDoubleFreeAndSafe(start, end *[2]int8, player int8, Board, CapturableBoard *[19][19]int8) bool {

	count := 0

	var dx, dy int8
	dx = end[0] - start[0]
	dy = end[1] - start[1]
	dx, dy = getDir(dx, dy)

	if isInRange(start[0]-dx, start[1]-dy) {
		if Board[start[1]-dy][start[0]-dx] == 0 {
			count++
		}
	} else {
		return false
	}

	if isInRange(end[0]+dx, end[1]+dy) {
		if Board[end[1]+dy][end[0]+dx] == 0 {
			count++
		}
	} else {
		return false
	}

	if count == 2 {
		var i int8
		for i = 0; Board[start[1]+dy*i][start[0]+dx*i] == player; i++ {
			if CapturableBoard[start[1]+dy*i][start[0]+dx*i] != 0 {
				return false
			}
		}
		if isInRange(end[0]+dx*2, end[1]+dy*2) && Board[end[1]+dy*2][end[0]+dx*2] == player {
			return false
		} else if isInRange(start[0]-dx*2, start[1]-dy*2) && Board[start[1]-dy*2][start[0]-dx*2] == player {
			return false
		}
		return true
	}
	return false
}

func checkLoosingPosition(rootNode *Node) bool {

	alignements := make([]alignementHeuristic, len(rootNode.StonesPlayed)*2)
	count := 0

	getAlignementProps(&rootNode.Board, rootNode.Pos[0], rootNode.Pos[1], rootNode.CurrentPlayer, &alignements, &count)
	for _, align := range alignements {
		if align.Player == 0 {
			break
		}

		if align.Length == 4 && align.Player == rootNode.CurrentPlayer {
			doublefree := isDoubleFreeAndSafe(&align.Start, &align.End, align.Player, &rootNode.Board, &rootNode.CapturableBoard)
			if doublefree {
				return true
			}
		}
	}
	return false
}

func computeLooseInBuild(rootNode *Node, move [2]int8, k int) bool {

	node := Node{}
	node.Pos = move
	node.Alpha = -1000000
	node.Beta = 1000000
	node.CurrentPlayer = rootNode.CurrentPlayer
	node.Board = rootNode.Board
	node.Board[move[1]][move[0]] = node.CurrentPlayer
	node.ThreeBoard = doubleThree(&node.Board)
	node.Cut = false
	node.CapturePlayerOne = rootNode.CapturePlayerOne
	node.CapturePlayerTwo = rootNode.CapturePlayerTwo
	if node.CurrentPlayer == 1 {
		node.CapturePlayerOne += isCaptured(&node.Board, node.Pos[0], node.Pos[1], node.CurrentPlayer)
	} else {
		node.CapturePlayerTwo += isCaptured(&node.Board, node.Pos[0], node.Pos[1], node.CurrentPlayer)
	}
	node.CapturableBoard = fcb(&node.Board, &node.ThreeBoard)
	lenStones := len(rootNode.StonesPlayed)
	node.StonesPlayed = make([][2]int8, lenStones+1)
	for i := 0; i < lenStones; i++ {
		node.StonesPlayed[i] = rootNode.StonesPlayed[i]
	}
	node.StonesPlayed[lenStones][0] = move[0]
	node.StonesPlayed[lenStones][1] = move[1]

	if computeWinInBuild(&node, k) {
		return true
	} else {
		return checkLoosingPosition(&node)
	}
}

func computeWinInBuild(node *Node, k int) bool {
	if node.CurrentPlayer == 1 {
		if node.CapturePlayerOne >= 5 {
			return true
		}
	} else {
		if node.CapturePlayerTwo >= 5 {
			return true
		}
	}

	for _, pos := range node.StonesPlayed {
		if node.Board[pos[1]][pos[0]] == node.CurrentPlayer {

			st, ed, aligned := isAligned(pos[0], pos[1], node.CurrentPlayer, &node.Board)

			if aligned {
				for i, elem := range st {
					end := ed[i]

					b := isSafeAlignement(elem, end, node.CurrentPlayer, &node.CapturableBoard)

					if b {
						if node.CurrentPlayer == 1 && node.CapturePlayerTwo == 4 {

							if !playerCanBeCaptured(node.CurrentPlayer, &node.CapturableBoard) {
								return true
							}

						} else if node.CurrentPlayer == 2 && node.CapturePlayerOne == 4 {

							if !playerCanBeCaptured(node.CurrentPlayer, &node.CapturableBoard) {
								return true
							}

						} else {
							return true

						}
					}
				}
			}
		}
	}
	return false
}

func buildChildP(rootNode *Node, move [2]int8, k, i int, c chan bool, d chan bool, w int) {
	var mutex = &sync.Mutex{}
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
		mutex.Lock()
		winPos = newNode.Pos
		mutex.Unlock()
		d <- true
		c <- true
		return
	}

	loose := computeLooseInBuild(rootNode, move, i)

	if loose {
		mutex.Lock()
		loosePos = newNode.Pos
		mutex.Unlock()
	}

	rootNode.Children = append(rootNode.Children, &newNode)
	d <- true
	buildTree(rootNode.Children[k], i+1)
	c <- true
	return
}

func buildChild(rootNode *Node, move *[2]int8, k, i int) {
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
	buildTree(rootNode.Children[k], i+1)
	return
}

func buildTreeP(rootNode *Node, i int) {
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
	getBestMoves(rootNode.CurrentPlayer, &listMoves, &rootNode.Board, &rootNode.ThreeBoard, &rootNode.CapturableBoard)

	w := width

	if width > len(listMoves) {
		w = len(listMoves)
	}
	for k, move := range listMoves[0:w] {
		go buildChildP(rootNode, move, k, i, a, b, w)
		<-b
	}
	for j := 0; j < w; j++ {
		<-a
	}
	return
}

func buildTree(rootNode *Node, i int) {
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
		getBestMoves(rootNode.CurrentPlayer, &listMoves, &rootNode.Board, &rootNode.ThreeBoard, &rootNode.CapturableBoard)

		w := width

		if width > len(listMoves) {
			w = len(listMoves)
		}
		for k, move := range listMoves[0:w] {
			buildChild(rootNode, &move, k, i)
		}
	}
	return
}
