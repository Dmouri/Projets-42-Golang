package main

func computeMyScore(score int, nb int8) int {
	if nb == 1 {
		score += 1
	} else if nb == 2 {
		score += 3
	} else if nb == 3 {
		score += 10
	} else if nb == 4 {
		score += 30
	} else if nb >= 5 {
		score += 70
	}

	return score
}

func computeEnnemyScore(score int, nb int8) int {
	if nb == 2 {
		score += 1
	} else if nb == 3 {
		score += 2
	} else if nb == 4 {
		score += 30
	} else if nb >= 5 {
		score += 60
	}

	return score
}

func lengthOfAlignementScore(px, py, player int8, board *[19][19]int8, score int, ennemy bool) (int8, int) {
	var alignementCount int8
	alignementCount = 0

	if ennemy {
		if nb := checkVerticalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore(score, nb)
		}
		if nb := checkHorizontalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore(score, nb)
		}
		if nb := checkDiagRAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore(score, nb)
		}
		if nb := checkDiagLAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeEnnemyScore(score, nb)
		}
	} else {
		if nb := checkVerticalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore(score, nb)
		}
		if nb := checkHorizontalAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore(score, nb)
		}
		if nb := checkDiagRAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore(score, nb)
		}
		if nb := checkDiagLAli(px, py, player, board); nb > 1 {
			alignementCount++
			score = computeMyScore(score, nb)
		}
	}

	return alignementCount, score
}

func getScore(pos *[2]int8, player int8, Board, ThreeBoard, CapturableBoard *[19][19]int8) int {
	score := 0

	newPlayer := getEnnemy(player)

	if state.Capture {
		if newPlayer == 1 {
			// si je me protège
			if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%(newPlayer+3) == 0 {
				score += 25

				// si j'attaque
			} else if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%5 == 0 {
				score += 25
			}
		} else {
			if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%(newPlayer+3) == 0 {
				score += 25
			} else if CapturableBoard[pos[1]][pos[0]] > 3 && CapturableBoard[pos[1]][pos[0]]%4 == 0 {
				score += 25
			}
		}
	}

	if alignementCount, scoreTmp := lengthOfAlignementScore(pos[0], pos[1], newPlayer, Board, score, false); alignementCount > 0 {
		score = scoreTmp
	}

	if alignementCount, scoreTmp := lengthOfAlignementScore(pos[0], pos[1], getEnnemy(newPlayer), Board, score, true); alignementCount > 0 {
		score = scoreTmp
	}

	return score
}

func getPlayerTurn(i int) int {
	if i == 0 {
		return 1
	}
	if i%2 == 0 {
		return 1
	} else {
		return 2
	}
}

func getAlignementProps(board *[19][19]int8, x, y, currentP int8, alignements *[]alignementHeuristic, count *int) {

	checkVerticalHeuristic(x, y, currentP, board, alignements, count)

	checkHorizontalHeuristic(x, y, currentP, board, alignements, count)

	checkDiagLHeuristic(x, y, currentP, board, alignements, count)

	checkDiagRHeuristic(x, y, currentP, board, alignements, count)

}

func alignementIsInRange(align *alignementHeuristic, board *[19][19]int8, dx, dy int8) (bool, int) {

	count := 5
	freeCount := 0
	firstHalf := false
	countFirstHalf := 0
	score := 0
	ennemy := getEnnemy(board[align.Start[1]][align.Start[0]])
	var i int8
	for i = 0; count > 0; count-- {
		if isInRange(align.Start[0]+dx*i, align.Start[1]+dy*i) && board[align.Start[1]+dy*i][align.Start[0]+dx*i] != ennemy {
			freeCount++
		} else {
			break
		}
		i++
	}
	if freeCount == 5 {
		score += 2
		firstHalf = true
		if isInRange(align.Start[0]+dx*6, align.Start[1]+dy*6) && board[align.Start[1]+dy*6][align.Start[0]+dx*6] != ennemy {
			score += 1
		}
	} else {
		countFirstHalf = freeCount
	}

	count = 5
	freeCount = 0
	for i = 0; count > 0; count-- {
		if isInRange(align.End[0]-dx*i, align.End[1]-dy*i) && board[align.End[1]-dy*i][align.End[0]-dx*i] != ennemy {
			freeCount++
		} else {
			break
		}
		i++
	}
	if freeCount == 5 {
		score += 2
		if isInRange(align.End[0]-dx*6, align.End[1]-dy*6) && board[align.End[1]-dy*6][align.End[0]-dx*6] != ennemy {
			score += 1
		}
		return true, score
	} else {
		if !firstHalf {
			if countFirstHalf+freeCount >= 5 {
				score = 1
				return true, score
			}
		} else {
			return true, score
		}
	}
	return false, 0
}

func scoreAlignement(align alignementHeuristic, board *[19][19]int8, player int8) int {
	score := 0
	var dx, dy int8
	dx = align.End[0] - align.Start[0]
	dy = align.End[1] - align.Start[1]

	dx, dy = getDir(dx, dy)

	if b, blocked := alignementIsInRange(&align, board, dx, dy); b {

		if align.Length == 2 {
			score += 5
		} else if align.Length == 3 {
			score += 100
			score += blocked * 15
		} else if align.Length == 4 {
			score += 400
			score += blocked * 50
		} else if align.Length >= 5 {
			score += 1500
			score += blocked * 50
		}

	}

	if player != align.Player {
		score *= -1
	}

	return score
}

func heuristicAlignement(board *[19][19]int8, stonesPlayed *[][2]int8, x, y, player int8) int {
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
		score += scoreAlignement(align, board, player)
	}

	return score
}

func heuristicCapture(rootNode *Node) int {
	score := 0
	captureP1, captureP2 := countCapturable(&rootNode.CapturableBoard)
	if rootNode.CurrentPlayer == 1 {
		score += (10 * int(rootNode.CapturePlayerOne))
		score += (-10 * int(rootNode.CapturePlayerTwo))
		if rootNode.CapturePlayerOne == 4 && captureP1 > 1 {
			score *= 5
		} else if rootNode.CapturePlayerTwo == 4 && captureP2 > 1 {
			score = -5 * ft_abs(score)
		}
		score += captureP1*15 - captureP2*15

	} else {
		score += (-10 * int(rootNode.CapturePlayerOne))
		score += (10 * int(rootNode.CapturePlayerTwo))
		if rootNode.CapturePlayerOne == 4 && captureP1 > 1 {
			score = -5 * ft_abs(score)
		} else if rootNode.CapturePlayerTwo == 4 && captureP2 > 1 {
			score *= 5
		}
		score += captureP2*15 - captureP1*15
	}
	return score
}

func heuristic(rootNode *Node, depth int) int {
	score := 0

	if state.Capture {
		score += heuristicCapture(rootNode)
	}

	score += heuristicAlignement(&rootNode.Board, &rootNode.StonesPlayed, rootNode.Pos[0], rootNode.Pos[1], rootNode.CurrentPlayer)

	if rootNode.Cut {
		score += 1000000
		score *= depth + 1
		return score
	}

	return score
}
