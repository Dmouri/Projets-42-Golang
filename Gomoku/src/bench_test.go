package main

import (
	"testing"
)

// func TestSum(t *testing.T) {
// 	board, _ := generateBoard2()

// 	count, size := lengthOfAlignement(10, 10, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("8 10")
// 	}
// 	count, size = lengthOfAlignement(8, 10, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("8 10")
// 	}

// 	count, size = lengthOfAlignement(9, 10, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("9 10")
// 	}

// 	count, size = lengthOfAlignement(10, 10, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("10 10")
// 	}

// 	count, size = lengthOfAlignement(10, 9, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("10 9")
// 	}

// 	count, size = lengthOfAlignement(10, 8, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("10 8")
// 	}

// 	count, size = lengthOfAlignement(9, 8, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("9 8")
// 	}

// 	count, size = lengthOfAlignement(8, 8, 1, &board)
// 	if count != 1 || !Equal(size, []int8{2}) {
// 		t.Errorf("8 8")
// 	}

// 	board[8][10] = 1
// 	board[0][0] = 2

// 	count, size = lengthOfAlignement(8, 10, 1, &board)
// 	if count != 1 || !Equal(size, []int8{3}) {
// 		t.Errorf("Align 3 - ")
// 	}

// 	board[10][8] = 1
// 	board[1][1] = 2

// 	count, size = lengthOfAlignement(7, 11, 1, &board)
// 	if count != 1 || !Equal(size, []int8{4}) {
// 		t.Errorf("Align 4 - ")
// 	}
// }

func Equal(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func generateBoard2() ([19][19]int8, [][2]int8) {
	board := [19][19]int8{}
	board[9][9] = 1
	board[9][8] = 2
	stonesPlayed := [][2]int8{}
	stonesPlayed = append(stonesPlayed, [2]int8{9, 9})
	stonesPlayed = append(stonesPlayed, [2]int8{8, 9})
	return board, stonesPlayed
}

func generateBoard() ([19][19]int8, [][2]int8) {
	// 10 stones
	board := [19][19]int8{}
	board[9][9] = 1
	board[8][8] = 2
	board[9][8] = 1
	board[8][9] = 2
	board[7][8] = 1
	board[8][7] = 2
	board[7][7] = 1
	board[10][10] = 2
	board[9][10] = 1
	board[7][10] = 2
	// ++
	// board[14][2] = 1
	// board[2][14] = 2
	// board[11][4] = 1
	// board[4][11] = 2
	// board[18][18] = 1
	// board[18][0] = 2
	// board[12][12] = 1
	// board[11][11] = 2
	// board[4][4] = 1
	// board[5][5] = 2
	stonesPlayed := [][2]int8{}
	stonesPlayed = append(stonesPlayed, [2]int8{9, 9})
	stonesPlayed = append(stonesPlayed, [2]int8{8, 8})
	stonesPlayed = append(stonesPlayed, [2]int8{8, 9})
	stonesPlayed = append(stonesPlayed, [2]int8{9, 8})
	stonesPlayed = append(stonesPlayed, [2]int8{7, 8})
	stonesPlayed = append(stonesPlayed, [2]int8{8, 7})
	stonesPlayed = append(stonesPlayed, [2]int8{7, 7})
	stonesPlayed = append(stonesPlayed, [2]int8{10, 10})
	stonesPlayed = append(stonesPlayed, [2]int8{10, 9})
	stonesPlayed = append(stonesPlayed, [2]int8{10, 7})
	return board, stonesPlayed
}

func BenchmarkDoubleThree(t *testing.B) {
	board, _ := generateBoard()
	for i := 0; i < t.N; i++ {
		doubleThree(&board)
	}
}

func BenchmarkFillCapturableBoard(t *testing.B) {
	board, _ := generateBoard()
	ThreeBoard := doubleThree(&board)
	for i := 0; i < t.N; i++ {
		fillCapturableBoard(&board, &ThreeBoard)
	}
}

func BenchmarkFcb(t *testing.B) {
	board, _ := generateBoard()
	ThreeBoard := doubleThree(&board)
	for i := 0; i < t.N; i++ {
		fcb(&board, &ThreeBoard)
	}
}

func BenchmarkGetMoves(t *testing.B) {
	width = 4
	board, stonesPlayed := generateBoard()
	ThreeBoard := doubleThree(&board)
	CapturableBoard := fillCapturableBoard(&board, &ThreeBoard)
	var player int8
	player = 1
	for i := 0; i < t.N; i++ {
		listMoves := [][2]int8{}
		validMoves := board
		getValidMoves(&validMoves, &ThreeBoard, &listMoves, &stonesPlayed)
		getBestMoves(player, &listMoves, &board, &ThreeBoard, &CapturableBoard)
	}
}

func BenchmarkGetScore(t *testing.B) {
	board, _ := generateBoard()
	ThreeBoard := doubleThree(&board)
	capturableBoard := fillCapturableBoard(&board, &ThreeBoard)
	move := [2]int8{10, 11}
	for i := 0; i < t.N; i++ {
		getScore(&move, 1, &board, &ThreeBoard, &capturableBoard)
	}
}

func BenchmarkHeuristic(t *testing.B) {
	board, stonesPlayed := generateBoard()
	ThreeBoard := doubleThree(&board)
	capturableBoard := fillCapturableBoard(&board, &ThreeBoard)
	rootNode := &Node{Pos: [2]int8{10, 11}, Alpha: -1000000, Beta: 1000000,
		Board: board, ThreeBoard: ThreeBoard, CapturableBoard: capturableBoard, StonesPlayed: stonesPlayed, CurrentPlayer: 1}
	for i := 0; i < t.N; i++ {
		heuristic(rootNode, 0)
	}
}

func BenchmarkBuildTree(t *testing.B) {
	deepth, width, routine = 5, 10, true
	board, stonesPlayed := generateBoard()
	ThreeBoard := doubleThree(&board)
	capturableBoard := fillCapturableBoard(&board, &ThreeBoard)
	rootNode := &Node{Pos: [2]int8{-1, -1}, Alpha: -1000000, Beta: 1000000,
		Board: board, ThreeBoard: ThreeBoard, CapturableBoard: capturableBoard, StonesPlayed: stonesPlayed, CurrentPlayer: 1}
	for i := 0; i < t.N; i++ {
		buildTree(rootNode, 0)
	}
}

func BenchmarkAlgo(t *testing.B) {
	deepth, width, routine = 5, 10, true
	board, stonesPlayed := generateBoard()
	ThreeBoard := doubleThree(&board)
	capturableBoard := fillCapturableBoard(&board, &ThreeBoard)
	rootNode := &Node{Pos: [2]int8{-1, -1}, Alpha: -1000000, Beta: 1000000,
		Board: board, ThreeBoard: ThreeBoard, CapturableBoard: capturableBoard, StonesPlayed: stonesPlayed, CurrentPlayer: 1}
	buildTree(rootNode, 0)
	for i := 0; i < t.N; i++ {
		algo(rootNode)
	}
}
