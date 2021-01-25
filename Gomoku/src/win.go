package main

func mgtWinnableAlignement() {
	newTab := []winInProgress{}
	for k, alignement := range state.Winnable {
		dx, dy, length := getLengthAndDir(alignement.Start, alignement.End)
		splitCase := false
		if length >= 11 {
			splitCase = true
		}
		count := 0
		alreadySet := false
		edited := false
		sizeReached := false
		st := alignement.Start
		splitSt := [2]int8{-1, -1}
		splitEnd := [2]int8{-1, -1}
		end := [2]int8{-1, -1}
		var i int8
		for i = 0; i < length; i++ {
			if state.Board[alignement.Start[1]+dy*i][alignement.Start[0]+dx*i] != alignement.Player {
				if splitCase && count >= 5 {
					splitSt = [2]int8{alignement.Start[0] + dx*(i+1), alignement.Start[1] + dy*(i+1)}
					alreadySet = true
				}
				edited = true
				count = 0
				if !sizeReached {
					st = [2]int8{alignement.Start[0] + dx*i, alignement.Start[1] + dy*i}
				}
			} else {
				count++
			}
			if count >= 5 {
				sizeReached = true
				if !alreadySet {
					end = [2]int8{alignement.Start[0] + dx*i, alignement.Start[1] + dy*i}
				} else {
					splitEnd = [2]int8{alignement.Start[0] + dx*i, alignement.Start[1] + dy*i}
				}
			}
		}
		if splitCase && alreadySet && splitEnd[0] != -1 {
			newTab = append(newTab, winInProgress{alignement.Player, splitSt, splitEnd})
		}
		if edited {
			if sizeReached {
				state.Winnable[k].Start = st
				state.Winnable[k].End = end
				newTab = append(newTab, state.Winnable[k])
			}
		} else {
			newTab = append(newTab, alignement)
		}
	}
	state.Winnable = newTab
	if state.WinConditionOne {
		if getWinnableEntryOfPlayer(1) == 0 {
			state.WinConditionOne = false
		}
	}
	if state.WinConditionTwo {
		if getWinnableEntryOfPlayer(2) == 0 {
			state.WinConditionTwo = false
		}
	}
}

func appendWinnable(start, end [2]int8, player int8) {
	seen := 0
	toRemove := -1
	for k, elem := range state.Winnable {
		if elem.Player == player && elem.Start[0] == start[0] && elem.Start[1] == start[1] &&
			elem.End[0] == end[0] && elem.End[1] == end[1] {
			seen++
			break
		}
		dx, dy, _ := getLengthAndDir(elem.Start, elem.End)
		if elem.Player == player {
			if (start[0] == elem.Start[0]-dx && end[0] == elem.End[0] && end[1] == elem.End[1]) ||
				(start[1] == elem.Start[1]-dy && end[0] == elem.End[0] && end[1] == elem.End[1]) {
				toRemove = k
			} else if (end[0] == elem.End[0]+dx && start[0] == elem.Start[0] && start[1] == elem.Start[1]) ||
				(end[1] == elem.End[1]+dy && start[0] == elem.Start[0] && start[1] == elem.Start[1]) {
				toRemove = k
			}
		}
	}
	if seen != 0 {
		return
	} else {
		state.Winnable = append(state.Winnable, winInProgress{player, start, end})
	}
	if toRemove >= 0 {
		state.Winnable = append(state.Winnable[:toRemove], state.Winnable[toRemove+1:]...)
	}
}

func mgtWinCondition() {
	if state.WinConditionOne {
		if len(state.Winnable) > 0 {
			for _, elem := range state.Winnable {
				if elem.Player == 1 {
					if isAlignementFree(elem.Start, elem.End, elem.Player) {
						if state.CapturePlayerTwo != 4 {
							state.WinOne = true
						} else {
							if !HasUnsafePosition(state.CapturableBoard, 1) {
								state.WinOne = true
							}
						}
					}
				}
			}
		} else {
			if !HasUnsafePosition(state.CapturableBoard, 1) {
				state.WinOne = true
			}
		}
	}
	if state.WinConditionTwo {
		if len(state.Winnable) > 0 {
			for _, elem := range state.Winnable {
				if elem.Player == 2 {
					if isAlignementFree(elem.Start, elem.End, elem.Player) {
						if state.CapturePlayerOne != 4 {
							state.WinTwo = true
						} else {
							if !HasUnsafePosition(state.CapturableBoard, 2) {
								state.WinTwo = true
							}
						}
					}
				}
			}
		} else {
			if !HasUnsafePosition(state.CapturableBoard, 2) {
				state.WinTwo = true
			}
		}

	}
}
