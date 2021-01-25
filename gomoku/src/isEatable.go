package main

func isEatable(Board, ThreeBoard *[19][19]int8, px, py, player int8) bool {
	if player == 0 {
		return false
	}
	ennemy := getEnnemy(player)

	if isInRange(px, py+1) {
		if Board[py+1][px] == player {
			if isInRange(px, py+2) && isInRange(px, py-1) {
				if (Board[py+2][px] == ennemy && Board[py-1][px] == 0 && ThreeBoard[py-1][px] != ennemy && ThreeBoard[py-1][px] != 3) ||
					(Board[py-1][px] == ennemy && Board[py+2][px] == 0 && ThreeBoard[py+2][px] != ennemy && ThreeBoard[py+2][px] != 3) {
					return true
				}
			}
		} else if Board[py+1][px] != 0 {
			if isInRange(px, py-1) && isInRange(px, py-2) {
				if Board[py-1][px] == player && Board[py-2][px] == 0 && ThreeBoard[py-2][px] != ennemy && ThreeBoard[py-2][px] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px+1, py+1) {
		if Board[py+1][px+1] == player {
			if isInRange(px+2, py+2) && isInRange(px-1, py-1) {
				if (Board[py+2][px+2] == ennemy && Board[py-1][px-1] == 0 && ThreeBoard[py-1][px-1] != ennemy && ThreeBoard[py-1][px-1] != 3) ||
					(Board[py-1][px-1] == ennemy && Board[py+2][px+2] == 0 && ThreeBoard[py+2][px+2] != ennemy && ThreeBoard[py+2][px+2] != 3) {
					return true
				}
			}
		} else if Board[py+1][px+1] != 0 {
			if isInRange(px-1, py-1) && isInRange(px-2, py-2) {
				if Board[py-1][px-1] == player && Board[py-2][px-2] == 0 && ThreeBoard[py-2][px-2] != ennemy && ThreeBoard[py-2][px-2] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px+1, py) {
		if Board[py][px+1] == player {
			if isInRange(px+2, py) && isInRange(px-1, py) {
				if (Board[py][px+2] == ennemy && Board[py][px-1] == 0 && ThreeBoard[py][px-1] != ennemy && ThreeBoard[py][px-1] != 3) ||
					(Board[py][px-1] == ennemy && Board[py][px+2] == 0 && ThreeBoard[py][px+2] != ennemy && ThreeBoard[py][px+2] != 3) {
					return true
				}
			}
		} else if Board[py][px+1] != 0 {
			if isInRange(px-1, py) && isInRange(px-2, py) {
				if Board[py][px-1] == player && Board[py][px-2] == 0 && ThreeBoard[py][px-2] != ennemy && ThreeBoard[py][px-2] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px+1, py-1) {
		if Board[py-1][px+1] == player {
			if isInRange(px+2, py-2) && isInRange(px-1, py+1) {
				if (Board[py-2][px+2] == ennemy && Board[py+1][px-1] == 0 && ThreeBoard[py+1][px-1] != ennemy && ThreeBoard[py+1][px-1] != 3) ||
					(Board[py+1][px-1] == ennemy && Board[py-2][px+2] == 0 && ThreeBoard[py-2][px+2] != ennemy && ThreeBoard[py-2][px+2] != 3) {
					return true
				}
			}
		} else if Board[py-1][px+1] != 0 {
			if isInRange(px-1, py+1) && isInRange(px-2, py+2) {
				if Board[py+1][px-1] == player && Board[py+2][px-2] == 0 && ThreeBoard[py+2][px-2] != ennemy && ThreeBoard[py+2][px-2] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px, py-1) {
		if Board[py-1][px] == player {
			if isInRange(px, py-2) && isInRange(px, py+1) {
				if (Board[py-2][px] == ennemy && Board[py+1][px] == 0 && ThreeBoard[py+1][px] != ennemy && ThreeBoard[py+1][px] != 3) ||
					(Board[py+1][px] == ennemy && Board[py-2][px] == 0 && ThreeBoard[py-2][px] != ennemy && ThreeBoard[py-2][px] != 3) {
					return true
				}
			}
		} else if Board[py-1][px] != 0 {
			if isInRange(px, py+1) && isInRange(px, py+2) {
				if Board[py+1][px] == player && Board[py+2][px] == 0 && ThreeBoard[py+2][px] != ennemy && ThreeBoard[py+2][px] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px-1, py-1) {
		if Board[py-1][px-1] == player {
			if isInRange(px-2, py-2) && isInRange(px+1, py+1) {
				if (Board[py-2][px-2] == ennemy && Board[py+1][px+1] == 0 && ThreeBoard[py+1][px+1] != ennemy && ThreeBoard[py+1][px+1] != 3) ||
					(Board[py+1][px+1] == ennemy && Board[py-2][px-2] == 0 && ThreeBoard[py-2][px-2] != ennemy && ThreeBoard[py-2][px-2] != 3) {
					return true
				}
			}
		} else if Board[py-1][px-1] != 0 {
			if isInRange(px+1, py+1) && isInRange(px+2, py+2) {
				if Board[py+1][px+1] == player && Board[py+2][px+2] == 0 && ThreeBoard[py+2][px+2] != ennemy && ThreeBoard[py+2][px+2] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px-1, py) {
		if Board[py][px-1] == player {
			if isInRange(px-2, py) && isInRange(px+1, py) {
				if (Board[py][px-2] == ennemy && Board[py][px+1] == 0 && ThreeBoard[py][px+1] != ennemy && ThreeBoard[py][px+1] != 3) ||
					(Board[py][px+1] == ennemy && Board[py][px-2] == 0 && ThreeBoard[py][px-2] != ennemy && ThreeBoard[py][px-2] != 3) {
					return true
				}
			}
		} else if Board[py][px-1] != 0 {
			if isInRange(px+1, py) && isInRange(px+2, py) {
				if Board[py][px+1] == player && Board[py][px+2] == 0 && ThreeBoard[py][px+2] != ennemy && ThreeBoard[py][px+2] != 3 {
					return true
				}
			}
		}
	}

	if isInRange(px-1, py+1) {
		if Board[py+1][px-1] == player {
			if isInRange(px-2, py+2) && isInRange(px+1, py-1) {
				if (Board[py+2][px-2] == ennemy && Board[py-1][px+1] == 0 && ThreeBoard[py-1][px+1] != ennemy && ThreeBoard[py-1][px+1] != 3) ||
					(Board[py-1][px+1] == ennemy && Board[py+2][px-2] == 0 && ThreeBoard[py+2][px-2] != ennemy && ThreeBoard[py+2][px-2] != 3) {
					return true
				}
			}
		} else if Board[py+1][px-1] != 0 {
			if isInRange(px+1, py+1) && isInRange(px+2, py+2) {
				if Board[py+1][px+1] == player && Board[py+2][px+2] == 0 && ThreeBoard[py+2][px+2] != ennemy && ThreeBoard[py+2][px+2] != 3 {
					return true
				}
			}
		}
	}

	return false
}

func ie(Board, ThreeBoard *[19][19]int8, px, py, player int8) (bool, int8, int8) {
	if player == 0 {
		return false, -1, -1
	}
	ennemy := getEnnemy(player)

	if isInRange(px, py+1) {
		if Board[py+1][px] == player {
			if isInRange(px, py+2) && isInRange(px, py-1) {
				if Board[py+2][px] == ennemy && Board[py-1][px] == 0 && ThreeBoard[py-1][px] != ennemy && ThreeBoard[py-1][px] != 3 {
					return true, px, py - 1
				} else if Board[py-1][px] == ennemy && Board[py+2][px] == 0 && ThreeBoard[py+2][px] != ennemy && ThreeBoard[py+2][px] != 3 {
					return true, px, py + 2
				}
			}
		} else if Board[py+1][px] == ennemy {
			if isInRange(px, py-1) && isInRange(px, py-2) {
				if Board[py-1][px] == player && Board[py-2][px] == 0 && ThreeBoard[py-2][px] != ennemy && ThreeBoard[py-2][px] != 3 {
					return true, px, py - 2
				}
			}
		} else {
			if isInRange(px, py-1) && isInRange(px, py-2) {
				if Board[py-1][px] == player && Board[py-2][px] == ennemy && ThreeBoard[py+1][px] != ennemy && ThreeBoard[py+1][px] != 3 {
					return true, px, py + 1
				}
			}
		}
	}

	if isInRange(px+1, py+1) {
		if Board[py+1][px+1] == player {
			if isInRange(px+2, py+2) && isInRange(px-1, py-1) {
				if Board[py+2][px+2] == ennemy && Board[py-1][px-1] == 0 && ThreeBoard[py-1][px-1] != ennemy && ThreeBoard[py-1][px-1] != 3 {
					return true, px - 1, py - 1
				} else if Board[py-1][px-1] == ennemy && Board[py+2][px+2] == 0 && ThreeBoard[py+2][px+2] != ennemy && ThreeBoard[py+2][px+2] != 3 {
					return true, px + 2, py + 2
				}
			}
		} else if Board[py+1][px+1] == ennemy {
			if isInRange(px-1, py-1) && isInRange(px-2, py-2) {
				if Board[py-1][px-1] == player && Board[py-2][px-2] == 0 && ThreeBoard[py-2][px-2] != ennemy && ThreeBoard[py-2][px-2] != 3 {
					return true, px - 2, py - 2
				}
			}
		} else {
			if isInRange(px-1, py-1) && isInRange(px-2, py-2) {
				if Board[py-1][px-1] == player && Board[py-2][px-2] == ennemy && ThreeBoard[py+1][px+1] != ennemy && ThreeBoard[py+1][px+1] != 3 {
					return true, px + 1, py + 1
				}
			}
		}
	}

	if isInRange(px+1, py) {
		if Board[py][px+1] == player {
			if isInRange(px+2, py) && isInRange(px-1, py) {
				if Board[py][px+2] == ennemy && Board[py][px-1] == 0 && ThreeBoard[py][px-1] != ennemy && ThreeBoard[py][px-1] != 3 {
					return true, px - 1, py
				} else if Board[py][px-1] == ennemy && Board[py][px+2] == 0 && ThreeBoard[py][px+2] != ennemy && ThreeBoard[py][px+2] != 3 {
					return true, px + 2, py
				}
			}
		} else if Board[py][px+1] == ennemy {
			if isInRange(px-1, py) && isInRange(px-2, py) {
				if Board[py][px-1] == player && Board[py][px-2] == 0 && ThreeBoard[py][px-2] != ennemy && ThreeBoard[py][px-2] != 3 {
					return true, px - 2, py
				}
			}
		} else {
			if isInRange(px-1, py) && isInRange(px-2, py) {
				if Board[py][px-1] == player && Board[py][px-2] == ennemy && ThreeBoard[py][px+1] != ennemy && ThreeBoard[py][px+1] != 3 {
					return true, px + 1, py
				}
			}
		}
	}

	if isInRange(px+1, py-1) {
		if Board[py-1][px+1] == player {
			if isInRange(px+2, py-2) && isInRange(px-1, py+1) {
				if Board[py-2][px+2] == ennemy && Board[py+1][px-1] == 0 && ThreeBoard[py+1][px-1] != ennemy && ThreeBoard[py+1][px-1] != 3 {
					return true, px - 1, py + 1
				} else if Board[py+1][px-1] == ennemy && Board[py-2][px+2] == 0 && ThreeBoard[py-2][px+2] != ennemy && ThreeBoard[py-2][px+2] != 3 {
					return true, px + 2, py - 2
				}
			}
		} else if Board[py-1][px+1] == ennemy {
			if isInRange(px-1, py+1) && isInRange(px-2, py+2) {
				if Board[py+1][px-1] == player && Board[py+2][px-2] == 0 && ThreeBoard[py+2][px-2] != ennemy && ThreeBoard[py+2][px-2] != 3 {
					return true, px - 2, py + 2
				}
			}
		} else {
			if isInRange(px-1, py+1) && isInRange(px-2, py+2) {
				if Board[py+1][px-1] == player && Board[py+2][px-2] == ennemy && ThreeBoard[py-1][px+1] != ennemy && ThreeBoard[py-1][px+1] != 3 {
					return true, px + 1, py - 1
				}
			}
		}
	}

	return false, -1, -1
}
