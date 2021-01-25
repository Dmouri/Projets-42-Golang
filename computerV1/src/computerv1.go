package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var sleep bool = false
var color bool = false
var wait int = 0

func manageEntry() []string {
	var args []string
	nb := 0
	for i, str := range os.Args {
		if i != 0 && (str[0] == '-' && regex(str, "^[-]{1}[a-vy-zA-VY-Z]{1,}.*$")) || (str[0] == '-' && len(str) == 1) {
			fmt.Printf("%sError:%s %sparameter %s is invalid.%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, str, CLR_CLOSE)
			os.Exit(1)
		} else if i != 0 && str[0] == '-' && len(str) >= 2 && str[1] == '-' {
			if regex(str, "^[-]{2}[s]{1}[=]{1}[0-9]{1,}$") {
				if sleep == true {
					fmt.Printf("%sError:%s %sParamter --s is duplicated.%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
					os.Exit(1)
				}
				match := regexMatching(str, "^[-]{2}[s]{1}[=]{1}([0-9]{1,})$")
				if len(match[1]) >= 6 {
					fmt.Printf("%sError:%s %sI can't wait so long!%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
					os.Exit(1)
				}
				sleep = true
				wait, _ = strconv.Atoi(match[1])
			} else if i != 0 && str[0] == '-' && regex(str, "^[-]{2}[c]{1}$") {
				if color == true {
					fmt.Printf("%sError:%s %sParamter --c is duplicated.%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
					os.Exit(1)
				}
				color = true
				setColor()
			} else {
				fmt.Printf("%sError:%s %sparameter %s is invalid.%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, str, CLR_CLOSE)
				os.Exit(1)
			}
		} else if i != 0 {
			nb = nb + 1
			args = append(args, strings.Replace(str, " ", "", -1))
		}
	}
	if nb < 1 {
		fmt.Printf("%sError:%s %sToo few parameters.\n", CLR_RED, CLR_CLOSE, CLR_WHITE)
		fmt.Printf("Usage: ./computorv1 [--s=<uint>]  [--c] <equation1> <equation2> <...>\n")
		fmt.Printf("\t--s=<uint> : Specify time to wait in milliseconds between two operations.\n")
		fmt.Printf("\t--c : Activate color\n")
		fmt.Printf("\tExample: ./computorv1 --s=500 \"4 * X^2 - 8 * X^1 + 50 * X^0 = 4 * X^2 + 50\" %s\n", CLR_CLOSE)
		os.Exit(1)
	}
	return args
}

func waitTime() {
	if sleep == true {
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}

func main() {
	arguments := manageEntry()

	for i, equation := range arguments {
		trimEquation := strings.Replace(equation, " ", "", -1)

		if regex(trimEquation, "^[*0-9xX^+-.]{1,}=[*0-9xX^+-.]{1,}$") {
			trimEquation = strings.ToUpper(trimEquation)
			arrayEquation := strings.Split(trimEquation, "=")
			if len(arrayEquation) == 2 {
				tokenizedLeftSide, tokenizedRightSide := tokenizeMembers(arrayEquation[0], arrayEquation[1])
				if checkTokenValidity(tokenizedLeftSide, tokenizedRightSide) {
					tokenizedLeftSide, tokenizedRightSide = purifyExpression(tokenizedLeftSide, tokenizedRightSide)
					printPurify(tokenizedLeftSide, tokenizedRightSide)
					tokenizedLeftSide = balanceExpression(tokenizedLeftSide, tokenizedRightSide)
					printBalance(tokenizedLeftSide)
					tokenMap, exposant := reduceExpression(tokenizedLeftSide)
					if printDegree(exposant) {
						solveExpression(tokenMap, exposant)
						displayNextEntry(arguments, i)
					} else {
						displayNextEntry(arguments, i)
					}
				} else {
					fmt.Printf("%sERROR: Equation number %d has bad format%s\n", CLR_RED, i, CLR_CLOSE)
					displayNextEntry(arguments, i)
				}
			} else {
				fmt.Printf("%sERROR: Equation number %d has bad format (multiple equals sign)%s\n", CLR_RED, i, CLR_CLOSE)
				displayNextEntry(arguments, i)
			}

		} else {
			fmt.Printf("%sWrong format, forbidden char%s\n", CLR_RED, CLR_CLOSE)
			displayNextEntry(arguments, i)
		}
	}
}
