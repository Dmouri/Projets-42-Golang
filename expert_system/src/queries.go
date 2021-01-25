package main

import (
	"context"
	"fmt"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
)

var currentQuery string

var red = "\x1b[31;1m"
var green = "\x1b[32;1m"
var yellow = "\x1b[33;1m"
var white = "\x1b[37;1m"
var stop = "\x1b[0m"

func editRule(ctx context.Context, rule string, graph [][]node) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "editRule")
	defer span.Finish()
	ret := rule
	for i, char := range rule {
		if char >= 'A' && char <= 'Z' {
			if graph[nodeMap[string(char)][0]][0].valid == true {
				ret = replaceAtIndex(ret, '1', i)
			} else {
				ret = replaceAtIndex(ret, '0', i)
			}
		}
	}
	return ret
}

func isRuleTrue(ctx context.Context, rule string, graph [][]node) (string, bool) {
	span, _ := opentracing.StartSpanFromContext(ctx, "isRuleTrue")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	edited := rule
	if regex(edited, "^.*[A-Z]{1}.*$") {
		if verbose {
			fmt.Printf("%sIs %s%s%s %strue%s or %sfalse%s ?\n", tabstr, white, edited, stop, green, stop, red, stop)
		}
		edited = editRule(ctx, rule, graph)
	}
	open, close := getInnerParenthese(ctx, edited)
	for open != -1 && close != -1 {
		tmp, _ := isRuleTrue(ctx, edited[open+1:close], graph)
		edited = edited[:open] + tmp + edited[close+1:]
		open, close = getInnerParenthese(ctx, edited)
	}
	if regex(edited, "^.*[!]{1}.*$") {
		edited = strings.Replace(edited, "!1", "0", -1)
		edited = strings.Replace(edited, "!0", "1", -1)
	}
	if regex(edited, "^.*[\\|]{1}.*$") {
		matches := regexMatch(edited, "^(.*)[\\|]{1}(.*)$")
		left := matches[1]
		right := matches[2]
		tmpLeft, _ := isRuleTrue(ctx, left, graph)
		tmpRight, _ := isRuleTrue(ctx, right, graph)
		if tmpLeft == "1" || tmpRight == "1" {
			return "1", true
		}
		return "0", false
	}
	if regex(edited, ".*[\\^]{1}.*$") {
		matches := regexMatch(edited, "^(.*)[\\^]{1}(.*)$")
		left := matches[1]
		right := matches[2]
		tmpLeft, _ := isRuleTrue(ctx, left, graph)
		tmpRight, _ := isRuleTrue(ctx, right, graph)
		if (tmpLeft == "1" && tmpRight == "0") || (tmpRight == "1" && tmpLeft == "0") {
			return "1", true
		}
		return "0", false
	}
	if regex(edited, "^.*[+].*$") {
		countZero := 0
		for _, elem := range edited {
			if elem == '0' {
				countZero++
			}
		}
		if countZero == 0 {
			if verbose {
				fmt.Printf("%sThe expression is %strue%s\n", tabstr, green, stop)
			}
			return "1", true
		}
		if verbose {
			fmt.Printf("%sThe expression is %sfalse%s\n", tabstr, red, stop)
		}
		return "", false
	}
	if edited == "1" {
		if verbose {
			fmt.Printf("%sThe expression is %strue%s\n", tabstr, green, stop)
		}
		return edited, true
	} else if edited == "0" {
		if verbose {
			fmt.Printf("%sThe expression is %sfalse%s\n", tabstr, red, stop)
		}
		return edited, false
	}
	return "", false
}

var tabstr = ""

func solveQuerie(ctx context.Context, graph [][]node, fact string) (bool, bool) {
	span, _ := opentracing.StartSpanFromContext(ctx, "solveQuerie")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	var inf bool
	count := 0
	if verbose {
		fmt.Printf("%sQuerying the following fact: %s%s%s \n", tabstr, white, fact, stop)
	}
	if nodeMap[fact][1] == 0 {
		if verbose {
			fmt.Printf("%s\t%s not defined so it is %sfalse%s by default\n", tabstr, fact, red, stop)
		}
		return false, false
	}
	line := graph[nodeMap[fact][0]]
	for i, rule := range line {
		if i == 0 {
			if rule.valid == true {
				if verbose {
					if rule.checked == false {
						fmt.Printf("%s\t%s%s%s is %strue%s because it is an inital fact\n", tabstr, white, fact, stop, green, stop)
					} else {
						fmt.Printf("%s\t%s%s%s is %strue%s because this fact has been checked already\n", tabstr, white, fact, stop, green, stop)
					}
				}
				return true, false
			}
			continue
		} else {
			if verbose {
				fmt.Printf("%s\tQuerying the following expression: %s%s%s\n", tabstr, white, rule.name, stop)
			}
			reg := buildRegex(ctx, rule.name, "[!]{0,1}[(]{0,1}[+\\^\\|]{0,1}[!]{0,1}([A-Z]{1})[)]{0,1}[+\\^\\|]{0,1}[)]{0,1}")
			matches := regexMatch(rule.name, reg)
			for j, letter := range matches {
				if j == 0 {
					continue
				} else {
					if letter == currentQuery {
						inf = true
						break
					}
					tabstr += "\t"
					myBool, inf := solveQuerie(ctx, graph, letter)
					tabstr = strings.Replace(tabstr, "\t", "", 1)
					if myBool {
						graph[nodeMap[letter][0]][0].valid = true
						graph[nodeMap[letter][0]][0].checked = true
					} else if inf == true {
						count++
						break
					} else {
						graph[nodeMap[letter][0]][0].checked = true
					}
				}
			}
			if i != len(line)-1 && (count != 0 || inf == true) {
				continue
			} else if i == len(line)-1 && (count != 0 || inf == true) {
				if _, check := isRuleTrue(ctx, rule.name, graph); check {
					return true, false
				}
				return false, true
			}
			if count != 0 || inf == true {
				return false, true
			}
			if _, check := isRuleTrue(ctx, rule.name, graph); check {
				return true, false
			}
		}
	}
	return false, false
}

func processQueries(ctx context.Context, graph [][]node, queries []string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "processQueries")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	for i, elem := range queries {
		currentQuery = elem
		if b, inf := solveQuerie(ctx, graph, elem); b == true {
			fmt.Printf("%s%s is true%s\n", green, elem, stop)
			if verbose && i != len(queries)-1 {
				fmt.Printf("\n")
			}
		} else if inf == true {
			fmt.Printf("%s%s can't be evaluated because of an infinite loop%s\n", yellow, elem, stop)
			if verbose && i != len(queries)-1 {
				fmt.Printf("\n")
			}
		} else {
			fmt.Printf("%s%s is false%s\n", red, elem, stop)
			if verbose && i != len(queries)-1 {
				fmt.Printf("\n")
			}
		}
	}
}
