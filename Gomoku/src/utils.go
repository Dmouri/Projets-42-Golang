package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func ft_abs(nb int) int {
	if nb < 0 {
		return nb * -1
	}
	return nb
}

func getMovesLong() [][2]int8 {
	return [][2]int8{
		{13, 13}, {13, 12}, {13, 11}, {13, 10}, {13, 9}, {13, 8}, {13, 7}, {13, 6}, {13, 5},
		{12, 5}, {11, 5}, {10, 5}, {9, 5}, {8, 5}, {7, 5}, {6, 5}, {5, 5},
		{5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}, {5, 11}, {5, 12}, {5, 13},
		{6, 13}, {7, 13}, {8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13},
	}
}

func getMovesLongPro() [][2]int8 {
	return [][2]int8{
		{12, 12}, {12, 11}, {12, 10}, {12, 9}, {12, 8}, {12, 7}, {12, 6},
		{11, 6}, {10, 6}, {9, 6}, {8, 6}, {7, 6}, {6, 6},
		{6, 7}, {6, 8}, {6, 9}, {6, 10}, {6, 11}, {6, 12},
		{7, 12}, {8, 12}, {9, 12}, {10, 12}, {11, 12},
	}
}

func (s alignementHeuristic) contains(ts *[]alignementHeuristic) bool {
	for _, c := range *ts {
		if s.End == c.End && s.Start == c.Start {
			return true
		}
	}
	return false
}

func getDir(x, y int8) (int8, int8) {

	var dx, dy int8
	dx, dy = 0, 0
	if x > 0 {
		dx = 1
	} else if x < 0 {
		dx = -1
	} else {
		dx = 0
	}

	if y > 0 {
		dy = 1
	} else if y < 0 {
		dy = -1
	} else {
		dy = 0
	}

	return dx, dy
}

func contains(s [][2]int8, e [2]int8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func regexMatch(str, regex string) []string {
	match, _ := regexp.Compile(regex)
	return match.FindStringSubmatch(str)
}

func getEnnemy(player int8) int8 {
	if player == 1 {
		return 2
	}
	return 1
}

func parseInt(k int) int {
	l := len(os.Args)
	if l <= k+1 {
		log.Fatal(": Error: Not enough paraneters.\n")
		os.Exit(1)
	}
	vle, err := strconv.Atoi(os.Args[k+1])
	if vle <= 0 || err != nil {
		log.Fatal(": Error: Bad paraneters.\n")
		os.Exit(1)
	}
	return vle
}

func isInRange(px, py int8) bool {
	if px > 18 || px < 0 || py > 18 || py < 0 {
		return false
	}
	return true
}

func isCapturable(Board *[19][19]int8, px, py, player int8) bool {
	if px > 18 || px < 0 || py > 18 || py < 0 {
		return false
	}
	if Board[py][px] != player {
		return false
	}
	return true
}

func getWinnableEntryOfPlayer(player int8) int {
	count := 0
	for _, elem := range state.Winnable {
		if elem.Player == player {
			count++
		}
	}
	return count
}

func countCapturable(CapturableBoard *[19][19]int8) (int, int) {
	captureP1, captureP2 := 0, 0
	y := 0
	for x := 0; x < 19; x++ {
		y = 0
		for y = 0; y < 19; y++ {
			if CapturableBoard[y][x] == 1 {
				captureP2++
			} else if CapturableBoard[y][x] == 2 {
				captureP1++
			}
		}
	}
	return captureP1, captureP2
}
