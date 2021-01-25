package main

import (
	"math"
	"time"
)

func minimax(rootNode *Node, depth int, alpha, beta int, maximizing bool) int {
	if depth == 0 || rootNode.Cut {
		return heuristic(rootNode, depth)
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

func algo(rootNode *Node) (int8, int8) {
	value := minimax(rootNode, deepth, -1000000, 1000000, true)
	for _, elem := range rootNode.Children {
		if elem.Alpha == value {
			x := elem.Pos[0]
			y := elem.Pos[1]
			return x, y
		}
	}
	return 0, 0
}

func buildWorker(rootNode *Node, i int) {
	if routine {
		buildTreeP(rootNode, i)
	} else {
		buildTree(rootNode, i)
	}
	return
}

func playIA(turn int) int {
	NodeAddresses = make(map[int][]*Node)
	winPos = [2]int8{-1, -1}
	loosePos = [2]int8{-1, -1}
	if turn == 1 {
		PutStone(9, 9)
		return gameRules(turn)
	}
	t1 := time.Now()
	rootNode := &Node{Pos: [2]int8{-1, -1}, Alpha: -1000000, Beta: 1000000, Board: state.Board, ThreeBoard: state.ThreeBoard, CapturableBoard: state.CapturableBoard, StonesPlayed: state.StonesPlayed, CurrentPlayer: getEnnemy(state.CurrentPlayer), Cut: false}
	buildWorker(rootNode, 0)
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

	x, y := algo(rootNode)
	t2 := time.Now()
	IATimer = time.Now().Sub(t1) + time.Now().Sub(t2)
	PutStone(x, y)
	return gameRules(turn)
}
