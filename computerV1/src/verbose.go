package main

import (
	"bufio"
	"fmt"
	"os"
)

var CLR_RED = "\x1b[39;0m"
var CLR_GREEN = "\x1b[39;0m"
var CLR_YELLOW = "\x1b[39;0m"
var CLR_BLUE = "\x1b[39;0m"
var CLR_MAGEN = "\x1b[39;0m"
var CLR_CYAN = "\x1b[39;0m"
var CLR_WHITE = "\x1b[39;0m"
var CLR_PURPLE = "\x1b[39;0m"
var CLR_CLOSE = "\x1b[39;0m"

func setColor() {
	CLR_RED = "\x1b[31;1m"
	CLR_GREEN = "\x1b[32;1m"
	CLR_YELLOW = "\x1b[33;1m"
	CLR_BLUE = "\x1b[34;1m"
	CLR_MAGEN = "\x1b[35;1m"
	CLR_CYAN = "\x1b[36;1m"
	CLR_WHITE = "\x1b[37;1m"
	CLR_PURPLE = "\x1b[0;35m"
	CLR_CLOSE = "\x1b[0m"
}

func printPurify(left, right []string) {
	waitTime()
	fmt.Printf("%sNatural expression: %s", CLR_BLUE, CLR_CLOSE)
	for i, elem := range left {
		if i == 0 {
			fmt.Printf("%s ", elem)
		} else {
			fmt.Printf("%c %s ", elem[0], elem[1:])
		}
	}
	fmt.Printf("= ")
	for i, elem := range right {
		if i == 0 {
			fmt.Printf("%s ", elem)
		} else if i == len(right) {
			fmt.Printf("%c %s\n", elem[0], elem[1:])
		} else {
			fmt.Printf("%c %s ", elem[0], elem[1:])
		}
	}
	fmt.Printf("\n")
}

func printBalance(left []string) {
	waitTime()
	fmt.Printf("%sExpression equal 0: %s", CLR_BLUE, CLR_CLOSE)
	for i, elem := range left {
		if i == 0 {
			fmt.Printf("%s ", elem)
		} else {
			fmt.Printf("%c %s ", elem[0], elem[1:])
		}
	}
	fmt.Printf("= 0\n")
}

func printReduced(keysArray []int, tokenMap map[int]float64) int {
	exposant := 0
	count := 0
	waitTime()
	fmt.Printf("%sReduced form: %s", CLR_BLUE, CLR_CLOSE)
	for _, elem := range keysArray {
		if exposant == 0 && tokenMap[elem] != 0 {
			exposant = elem
		}
		if tokenMap[elem] != 0 || elem == 0 {
			count++
			if elem > 1 {
				if tokenMap[elem] < 0 {
					if elem == exposant {
						if tokenMap[elem] == -1 {
							fmt.Printf("-X^%d ", elem)
						} else {
							fmt.Printf("%g*X^%d ", tokenMap[elem], elem)
						}
					} else {
						if tokenMap[elem] == -1 {
							fmt.Printf("- X^%d ", elem)
						} else {
							fmt.Printf("- %g*X^%d ", tokenMap[elem]*-1, elem)
						}
					}
				} else {
					if elem == exposant {
						if tokenMap[elem] == 1 {
							fmt.Printf("X^%d ", elem)
						} else {
							fmt.Printf("%g*X^%d ", tokenMap[elem], elem)
						}
					} else {
						if tokenMap[elem] == 1 {
							fmt.Printf("+ X^%d ", elem)
						} else {
							fmt.Printf("+ %g*X^%d ", tokenMap[elem], elem)
						}
					}
				}
			} else if elem == 1 {
				if tokenMap[elem] < 0 {
					if elem == exposant {
						if tokenMap[elem] == -1 {
							fmt.Printf("-X ")
						} else {
							fmt.Printf("%g*X ", tokenMap[elem])
						}
					} else {
						if tokenMap[elem] == -1 {
							fmt.Printf("- X ")
						} else {
							fmt.Printf("- %g*X ", tokenMap[elem]*-1)
						}
					}
				} else {
					if elem == exposant {
						if tokenMap[elem] == 1 {
							fmt.Printf("X ")
						} else {
							fmt.Printf("%g*X ", tokenMap[elem])
						}
					} else {
						if tokenMap[elem] == 1 {
							fmt.Printf("+ X ")
						} else {
							fmt.Printf("+ %g*X ", tokenMap[elem])
						}
					}
				}
			} else {
				if tokenMap[elem] < 0 {
					if elem == exposant {
						fmt.Printf("%g ", tokenMap[elem])
					} else {
						fmt.Printf("- %g ", tokenMap[elem]*-1)
					}
				} else {
					if elem == exposant {
						fmt.Printf("%g ", tokenMap[elem])
					} else {
						fmt.Printf("+ %g ", tokenMap[elem])
					}
				}
			}
		}
	}
	if exposant == 0 && tokenMap[0] == 0 && count == 0 {
		fmt.Printf("0 = 0\n")
	} else {
		fmt.Printf("= 0\n")
	}
	return exposant
}

func printDegree(exposant int) bool {
	waitTime()
	fmt.Printf("%sPolynomial degree:%s %d\n", CLR_BLUE, CLR_CLOSE, exposant)
	if exposant > 2 {
		waitTime()
		fmt.Printf("%sDegree is greater than 2, I can solve it but I am not allowed to%s\n", CLR_RED, CLR_CLOSE)
		return false
	}
	return true
}

func displayNextEntry(args []string, i int) {
	if i+1 != len(args) {
		fmt.Printf("\n\n%sPress <Enter> for the next equation%s\n", CLR_YELLOW, CLR_CLOSE)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
