package main

func checkVerticalHeuristic(px, py, player int8, board *[19][19]int8, alignements *[]alignementHeuristic, count *int) {
	var tmp int8
	tmp = 1
	var dy int8
	var startX int8
	var startY int8
	var endX int8
	var endY int8
	startX = -1
	startY = -1
	endX = -1
	endY = -1
	for dy = -1; isInRange(px, py+dy); dy-- {
		if board[py+dy][px] == player {
			tmp++
		} else {
			startX = px
			startY = py + dy + 1
			break
		}
	}
	if startX == -1 || startY == -1 {
		startX = px
		startY = py
	}
	for dy = 1; isInRange(px, py+dy); dy++ {
		if board[py+dy][px] == player {
			tmp++
		} else {
			endX = px
			endY = py + dy - 1
			break
		}
	}
	if endX == -1 || endY == -1 {
		endX = px
		endY = py
	}
	if tmp != 1 {
		mystr := alignementHeuristic{tmp, player, [2]int8{startX, startY}, [2]int8{endX, endY}}
		b := mystr.contains(alignements)
		if !b {
			(*alignements)[(*count)] = mystr
			*count++
		}
	}
}

func checkHorizontalHeuristic(px, py, player int8, board *[19][19]int8, alignements *[]alignementHeuristic, count *int) {
	var tmp int8
	tmp = 1
	var dx int8
	var startX int8
	var startY int8
	var endX int8
	var endY int8
	startX = -1
	startY = -1
	endX = -1
	endY = -1
	for dx = -1; isInRange(px+dx, py); dx-- {
		if board[py][px+dx] == player {
			tmp++
		} else {
			startX = px + dx + 1
			startY = py
			break
		}
	}
	if startX == -1 || startY == -1 {
		startX = px
		startY = py
	}

	for dx = 1; isInRange(px+dx, py); dx++ {
		if board[py][px+dx] == player {
			tmp++
		} else {
			endX = px + dx - 1
			endY = py
			break
		}
	}
	if endX == -1 || endY == -1 {
		endX = px
		endY = py
	}
	if tmp != 1 {
		mystr := alignementHeuristic{tmp, player, [2]int8{startX, startY}, [2]int8{endX, endY}}
		b := mystr.contains(alignements)
		if !b {
			(*alignements)[(*count)] = mystr
			*count++
		}
	}
}

func checkDiagLHeuristic(px, py, player int8, board *[19][19]int8, alignements *[]alignementHeuristic, count *int) {
	var tmp int8
	tmp = 1
	var dx int8
	var dy int8
	var startX int8
	var startY int8
	var endX int8
	var endY int8
	startX = -1
	startY = -1
	endX = -1
	endY = -1
	for dx, dy = 1, -1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			tmp++
		} else {
			startX = px + dx - 1
			startY = py + dy + 1
			break
		}
		dy--
	}
	if startX == -1 || startY == -1 {
		startX = px
		startY = py
	}

	for dx, dy = -1, 1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			tmp++
		} else {
			endX = px + dx + 1
			endY = py + dy - 1
			break
		}
		dy++
	}
	if endX == -1 || endY == -1 {
		endX = px
		endY = py
	}
	if tmp != 1 {
		mystr := alignementHeuristic{tmp, player, [2]int8{startX, startY}, [2]int8{endX, endY}}
		b := mystr.contains(alignements)
		if !b {
			(*alignements)[(*count)] = mystr
			*count++
		}
	}
}

func checkDiagRHeuristic(px, py, player int8, board *[19][19]int8, alignements *[]alignementHeuristic, count *int) {
	var tmp int8
	tmp = 1
	var dx int8
	var dy int8
	var startX int8
	var startY int8
	var endX int8
	var endY int8
	startX = -1
	startY = -1
	endX = -1
	endY = -1
	for dx, dy = -1, -1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			tmp++
		} else {
			startX = px + dx + 1
			startY = py + dy + 1
			break
		}
		dy--
	}
	if startX == -1 || startY == -1 {
		startX = px
		startY = py
	}

	for dx, dy = 1, 1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			tmp++
		} else {
			endX = px + dx - 1
			endY = py + dy - 1
			break
		}
		dy++
	}
	if endX == -1 || endY == -1 {
		endX = px
		endY = py
	}
	if tmp != 1 {
		mystr := alignementHeuristic{tmp, player, [2]int8{startX, startY}, [2]int8{endX, endY}}
		b := mystr.contains(alignements)
		if !b {
			(*alignements)[(*count)] = mystr
			*count++
		}
	}
}

func checkVerticalAli(px, py, player int8, board *[19][19]int8) int8 {
	var count int8
	count = 1
	var dy int8
	for dy = -1; isInRange(px, py+dy); dy-- {
		if board[py+dy][px] == player {
			count++
		} else {
			break
		}
	}
	for dy = 1; isInRange(px, py+dy); dy++ {
		if board[py+dy][px] == player {
			count++
		} else {
			break
		}
	}
	return count
}

func checkHorizontalAli(px, py, player int8, board *[19][19]int8) int8 {
	var count int8
	count = 1
	var dx int8
	for dx = -1; isInRange(px+dx, py); dx-- {
		if board[py][px+dx] == player {
			count++
		} else {
			break
		}
	}
	for dx = 1; isInRange(px+dx, py); dx++ {
		if board[py][px+dx] == player {
			count++
		} else {
			break
		}
	}
	return count
}

func checkDiagRAli(px, py, player int8, board *[19][19]int8) int8 {
	var count int8
	count = 1
	var dx int8
	var dy int8
	for dx, dy = -1, -1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			count++
		} else {
			break
		}
		dy--
	}
	for dx, dy = 1, 1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			count++
		} else {
			break
		}
		dy++
	}
	return count
}

func checkDiagLAli(px, py, player int8, board *[19][19]int8) int8 {
	var count int8
	count = 1
	var dx int8
	var dy int8
	for dx, dy = 1, -1; isInRange(px+dx, py+dy); dx++ {
		if board[py+dy][px+dx] == player {
			count++
		} else {
			break
		}
		dy--
	}
	for dx, dy = -1, 1; isInRange(px+dx, py+dy); dx-- {
		if board[py+dy][px+dx] == player {
			count++
		} else {
			break
		}
		dy++
	}
	return count
}
