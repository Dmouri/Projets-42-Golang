package main

import (
	"regexp"
)

func regex(str, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

func regexMatching(str, regex string) []string {
	match, _ := regexp.Compile(regex)
	return match.FindStringSubmatch(str)
}

func tokenizeMembers(left, right string) ([]string, []string) {
	var tokenizedLeft []string
	var tokenizedRight []string
	str := left

	for {
		result := regexMatching(str, "(.*)([+|-].*)")
		if len(result) != 0 {
			tokenizedLeft = append(tokenizedLeft, result[2])
			str = result[1]
		} else {
			if str != "" {
				tokenizedLeft = append(tokenizedLeft, str)
			}
			break
		}
	}
	tokenizedLeft = reverseArray(tokenizedLeft)

	str = right
	for {
		result := regexMatching(str, "(.*)([+|-].*)")
		if len(result) != 0 {
			tokenizedRight = append(tokenizedRight, result[2])
			str = result[1]
		} else {
			if str != "" {
				tokenizedRight = append(tokenizedRight, str)
			}
			break
		}
	}
	tokenizedRight = reverseArray(tokenizedRight)
	return tokenizedLeft, tokenizedRight
}

func reverseArray(input []string) []string {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}

func reverseArrayInt(input []int) []int {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}

func checkTokenValidity(left, right []string) bool {
	myReg := "^[+|-]{0,1}[0-9]{1,}[.][0-9]{1,}[*][x|X]$|^[+|-]{0,1}[0-9]{1,}[*][x|X]$|^[+|-]{0,1}[0-9]{1,}[.][0-9]{1,}$|^[+|-]{0,1}[0-9]{1,}$|^[+|-]{0,1}[0-9]{1,}[*][x|X][\\^][0-9]{1,}$|^[+|-]{0,1}[0-9]{1,}[.][0-9]{1,}[*][x|X][\\^][0-9]{1,}$|^[+|-]{0,1}[x|X]$|^[+|-]{0,1}[x|X][\\^][0-9]{1,}$"
	for _, token := range left {
		if regex(token, myReg) != true {
			return false
		}
	}

	for _, token := range right {
		if regex(token, myReg) != true {
			return false
		}
	}
	return true
}
