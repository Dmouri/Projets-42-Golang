package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func purifyExpression(left, right []string) ([]string, []string) {
	for i, _ := range left {
		reg := regexp.MustCompile("(^[+|-]{0,1})[0]*([1-9]{1,}.*)$")
		left[i] = reg.ReplaceAllString(left[i], "$1$2")
		reg = regexp.MustCompile("(^[+|-]{0,1}[0-9]{1,}[.]{0,1}[0-9]*[*]{1}[X]{1}[\\^])[0]*([1-9]{0,}[0]*$)")
		left[i] = reg.ReplaceAllString(left[i], "$1$2")
		if left[i][len(left[i])-1] == '^' {
			left[i] = left[i] + "0"
		}
		left[i] = left[i] + "\000"
		left[i] = strings.Replace(left[i], "X^1\000", "X", -1)
		left[i] = strings.Replace(left[i], "*X^0\000", "", -1)
		left[i] = strings.Replace(left[i], "\000", "", -1)
	}
	for i, _ := range right {
		reg := regexp.MustCompile("(^[+|-]{0,1})[0]*([1-9]{1,}.*)$")
		right[i] = reg.ReplaceAllString(right[i], "$1$2")
		reg = regexp.MustCompile("(^[+|-]{0,1}[0-9]{1,}[.]{0,1}[0-9]*[*]{1}[X]{1}[\\^])[0]*([1-9]{0,}[0]*$)")
		right[i] = reg.ReplaceAllString(right[i], "$1$2")
		if right[i][len(right[i])-1] == '^' {
			right[i] = right[i] + "0"
		}
		right[i] = right[i] + "\000"
		right[i] = strings.Replace(right[i], "X^1\000", "X", -1)
		right[i] = strings.Replace(right[i], "*X^0\000", "", -1)
		right[i] = strings.Replace(right[i], "\000", "", -1)
	}
	return left, right
}

func balanceExpression(left, right []string) []string {
	for i, _ := range right {
		if right[i][0] == '-' {
			tmp := right[i][1:]
			left = append(left, "+"+tmp)
		} else if right[i][0] == '+' {
			tmp := right[i][1:]
			left = append(left, "-"+tmp)
		} else if right[i][0] == '0' && len(right[i]) == 1 {

		} else {
			left = append(left, "-"+right[i])
		}
	}
	return left
}

func reduceExpression(left []string) (map[int]float64, int) {
	tokenMap := make(map[int]float64)
	for _, elem := range left {
		result := regexMatching(elem, "^(.*)[*]{1}[X]{1}[\\^]{1}([0-9]{1,})$|^(.*)[*]{1}([X]{1})$")
		if len(result) == 0 {
			if regex(elem, "^[+|-]{0,1}[x|X]{1}[\\^]{0,1}[0-9]{0,}$") {
				result2 := regexMatching(elem, "^([+|-]{0,1})[xX]{1}[\\^]{0,1}([0-9]{0,})$")
				sign := 1
				if result2[1] == "-" {
					sign = -1
				}
				if result2[2] == "" {
					tokenMap[1] = tokenMap[1] + float64(sign)
				} else {
					exposant, _ := strconv.Atoi(result2[2])
					tokenMap[exposant] = tokenMap[exposant] + float64(sign)
				}
			} else {
				coeff, _ := strconv.ParseFloat(elem, 64)
				tokenMap[0] = tokenMap[0] + coeff
			}
		} else {
			if result[4] == "X" {
				coeff, _ := strconv.ParseFloat(result[3], 64)
				tokenMap[1] = tokenMap[1] + coeff
			} else {
				exposant, _ := strconv.Atoi(result[2])
				coeff, _ := strconv.ParseFloat(result[1], 64)
				tokenMap[exposant] = tokenMap[exposant] + coeff
			}
		}
	}
	keysArray := make([]int, 0, len(tokenMap))
	for elem := range tokenMap {
		keysArray = append(keysArray, elem)
	}
	sort.Ints(keysArray)
	keysArray = reverseArrayInt(keysArray)
	exposant := printReduced(keysArray, tokenMap)
	return tokenMap, exposant
}

func power(x float64, y int) float64 {
	xbase := x
	for i := 1; i < y; i++ {
		x = x * xbase
	}
	return x
}

func sqrt(x float64) float64 {
	var z float64 = 1
	for i := 1; i <= 10; i++ {
		z = (z - (power(z, 2)-x)/(2*z))
	}
	return z
}

func fixZero(nb float64) float64 {
	if nb == -0 {
		nb = 0
	}
	return nb
}

func solveExpression(tokenMap map[int]float64, exposant int) {
	waitTime()
	if exposant == 0 {
		if tokenMap[exposant] == 0 {
			fmt.Printf("%sAll real numbers are solution of the equation%s\n", CLR_GREEN, CLR_CLOSE)
		} else {
			fmt.Printf("%sThe equation is false, there is no solution%s\n", CLR_GREEN, CLR_CLOSE)
		}
	} else if exposant == 1 {
		c := tokenMap[0] * -1
		if c == -0 {
			c = 0
		}
		fmt.Printf("%sThe polynomial degree is 1, so there is only one result%s\n", CLR_MAGEN, CLR_CLOSE)
		waitTime()
		fmt.Printf("%sX = %g/%g = %g%s\n", CLR_GREEN, c, tokenMap[1], c/tokenMap[1], CLR_CLOSE)
	} else {
		a := fixZero(tokenMap[2])
		b := fixZero(tokenMap[1])
		c := fixZero(tokenMap[0])
		waitTime()
		fmt.Printf("%sThe equation is under the form of a polynomial expression of 2nd degree. It's form is a*X^2 + b*X + c = 0.\nWhere%s %sa = %g%s, %sb = %g%s%s and %s%sc = %g%s\n", CLR_MAGEN, CLR_CLOSE, CLR_CYAN, a, CLR_CLOSE, CLR_PURPLE, b, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_YELLOW, c, CLR_CLOSE)
		waitTime()
		fmt.Printf("%sLet's calculate the discriminant where Δ = b^2-4*a*c.%s\n", CLR_MAGEN, CLR_CLOSE)
		waitTime()
		fmt.Printf("%sΔ%s%s = %s%s%g%s%s^2 - 4 * %s%s%g%s%s * %s%s%g%s\n", CLR_RED, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_PURPLE, b, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_CYAN, a, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_YELLOW, c, CLR_CLOSE)
		waitTime()
		delta := fixZero(b*b - 4*a*c)
		fmt.Printf("%sHere the discriminant%s %sΔ = %g%s\n", CLR_MAGEN, CLR_CLOSE, CLR_RED, delta, CLR_CLOSE)
		waitTime()
		if delta > 0 {
			fmt.Printf("%sΔ > 0.%s %sSo there is 2 solutions. Let's call the results%s %sx1%s %sand%s %sx2%s\n", CLR_RED, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_GREEN, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_GREEN, CLR_CLOSE)
			waitTime()
			fmt.Printf("%sx1%s %s= (-b - √Δ)/(2 * a)%s %s and %s %sx2%s %s= (-b + √Δ)/(2 * a)%s\n", CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
			waitTime()
			fmt.Printf("%sx1%s %s= (-(%s%s%g%s%s) - √%s%s%g%s%s)/(2 * %s%s%g%s%s) %s%s and %s %sx2%s %s= (-(%s%s%g%s%s) + √%s%s%g%s%s)/(2 * %s%s%g%s%s))%s\n", CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_PURPLE, b, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_RED, delta, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_CYAN, a, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_PURPLE, b, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_RED, delta, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_CYAN, a, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
			waitTime()
			x1 := fixZero((b*-1 - sqrt(delta)) / (2 * a))
			x2 := fixZero((b*-1 + sqrt(delta)) / (2 * a))
			fmt.Printf("%sx1 = %g%s  %sand%s  %sx2 = %g%s\n", CLR_GREEN, x1, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE, CLR_GREEN, x2, CLR_CLOSE)
		} else if delta < 0 {
			fmt.Printf("%sΔ < 0.%s %sSo there is no solution.%s\n", CLR_RED, CLR_CLOSE, CLR_GREEN, CLR_CLOSE)
		} else {
			fmt.Printf("%sΔ = 0.%s %sSo there is 1 solution%s\n", CLR_RED, CLR_CLOSE, CLR_MAGEN, CLR_CLOSE)
			waitTime()
			fmt.Printf("%sx%s%s = -b/(2a)%s\n", CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
			waitTime()
			fmt.Printf("%sx%s%s = -(%s%s%g%s%s)/(2 * %s%s%g%s%s)%s\n", CLR_GREEN, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_PURPLE, b, CLR_CLOSE, CLR_WHITE, CLR_CLOSE, CLR_CYAN, a, CLR_CLOSE, CLR_WHITE, CLR_CLOSE)
			waitTime()
			x := fixZero(b * -1 / (2 * a))
			fmt.Printf("%sx = %g\n%s", CLR_GREEN, x, CLR_CLOSE)
		}
	}
}
