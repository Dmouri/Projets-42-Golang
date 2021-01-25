package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

/* kind 0 = fact // kind 1 = rules
 */

var nodeMap = make(map[string][2]int)

type node struct {
	kind    bool
	name    string
	valid   bool
	checked bool
}

func exclamationMark(ctx context.Context, right string) bool {
	span, _ := opentracing.StartSpanFromContext(ctx, "exclamationMark")
	defer span.Finish()
	count := 0
	for _, char := range right {
		if char == '!' {
			count++
		}
	}
	if (count % 2) == 0 {
		return false
	}
	return true
}

func reverseMemberMark(ctx context.Context, left, right string) (string, string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "reverseMemberMark")
	defer span.Finish()
	matches := regexMatch(right, "^.*([A-Z]{1}).*$")
	right = matches[1]
	left = "!(" + left + ")"
	return left, right
}

func insertMemberInGraph(ctx context.Context, left, right string, graph [][]node) [][]node {
	span, _ := opentracing.StartSpanFromContext(ctx, "insertMemberInGraph")
	defer span.Finish()
	var nodeIndex [2]int
	if nodeMap[right][1] == 1 {
		newLeftNode := node{true, left, false, false}
		graph[nodeMap[right][0]] = append(graph[nodeMap[right][0]], newLeftNode)
	} else {
		nodeIndex[0] = len(graph)
		nodeIndex[1] = 1
		nodeMap[right] = nodeIndex
		newNodeArray := []node{{false, right, false, false}, {true, left, false, false}}
		graph = append(graph, newNodeArray)
	}
	return graph
}

func cleanMember(ctx context.Context, right string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "cleanMember")
	defer span.Finish()
	matches := regexMatch(right, "^.*([A-Z]{1}).*$")
	return matches[1]
}

func developNegation(ctx context.Context, right string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "developNegation")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	countOpen, countClose := getInnerParenthese(ctx, right)
	count := 0
	if countOpen > 0 && right[countOpen-1] == '!' {
		tmp := right[countOpen-1 : countClose+1]
		for i, elem := range tmp {
			if elem >= 'A' && elem <= 'Z' {
				if tmp[i+count-1] == '!' {
					tmp = tmp[:i+count-1] + tmp[i+count:]
					count--
				} else {
					tmp = tmp[:i+count] + "!" + tmp[i+count:]
					count++
				}
			}
		}
		tmp = strings.Replace(tmp, "!(", "", -1)
		tmp = strings.Replace(tmp, ")", "", -1)
		right = right[:countOpen-1] + tmp + right[countClose+1:]
	} else if countOpen == -1 && countClose == -1 {
		return right
	} else {
		if countOpen == -1 && countClose != -1 {
			countOpen = 0
			countClose++
			right = "(" + right
		}
		if countOpen != -1 && countClose == -1 {
			right = right + ")"
			countClose = len(right) - 1
		}
		if countOpen == 0 && countClose == len(right)-1 {
			right = strings.Replace(right, "(", "", -1)
			right = strings.Replace(right, ")", "", -1)
			return right
		} else if countClose == len(right)-1 {
			tmp := right[countOpen-1 : countClose]
			tmp = strings.Replace(tmp, "(", "", -1)
			tmp = strings.Replace(tmp, ")", "", -1)
			right = right[:countOpen-1] + tmp
		} else if countOpen == 0 {
			tmp := right[countOpen : countClose+1]
			tmp = strings.Replace(tmp, "(", "", -1)
			tmp = strings.Replace(tmp, ")", "", -1)
			right = tmp + right[countClose+1:]
		} else {
			tmp := right[countOpen : countClose+1]
			tmp = strings.Replace(tmp, "(", "", -1)
			tmp = strings.Replace(tmp, ")", "", -1)
			right = right[:countOpen] + tmp + right[countClose+1:]
		}
	}
	right = developNegation(ctx, right)
	return right
}

func doParentheseNeedDiscard(ctx context.Context, str string) (int, int, bool) {
	span, _ := opentracing.StartSpanFromContext(ctx, "doParentheseNeedDiscard")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if !regex(str, "^.*[(]{1}.*[\\^\\|]{1}.*[)]{1}.*$") {
		return 0, 0, false
	}
	for true {
		countOpen, countClose := getInnerParenthese(ctx, str)
		if countClose == -1 && countOpen == -1 {
			str = strings.Replace(str, "#", "(", -1)
			str = strings.Replace(str, "$", ")", -1)
			return countOpen, countClose, false
		}
		if regex(str[countOpen:countClose+1], "^.*[(]{1}.*[\\^\\|]{1}.*[)]{1}.*$") {
			str = strings.Replace(str, "#", "(", -1)
			str = strings.Replace(str, "$", ")", -1)
			if countOpen == 0 && countClose == len(str)-1 {
				return 0, 0, false
			}
			return countOpen, countClose, true
		}
		str = replaceAtIndex(str, '#', countOpen)
		str = replaceAtIndex(str, '$', countClose)
	}
	return 0, 0, false
}
func isRunePresent(ctx context.Context, str string, char string) (bool, []string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "isRunePresent")
	defer span.Finish()
	reg := "^.*[\\" + char + "]{1}.*$"
	if !regex(str, reg) {
		return false, nil
	}
	reg = "^([!]{0,1}[(]{0,1}.*)[\\" + char + "]{1}(.*[)]{0,1})$"
	matches := regexMatch(str, reg)
	return true, matches[1:]
}

func manageAmbiguity(ctx context.Context, matches []string, char string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "manageAmbiguity")
	defer span.Finish()
	match := ""
	white := "\x1b[37;1m"
	close := "\x1b[0m"
	entry := "default"
	if char == "|" {
		fmt.Printf("There is an ambiguity in the following expression : %s%s %s %s%s\n", white, matches[0], char, matches[1], close)
		fmt.Printf("To fix the ambiguity, please choose the member you want to set as true\n")
		fmt.Printf("For %s%s%s type %s1%s, for %s%s%s type %s2%s or for both type %s3%s\n", white, matches[0]+")", close, white, close, white, "("+matches[1], close, white, close, white, close)
		for true {
			fmt.Scanf("%s", &entry)
			if !(len(entry) != 1 && (entry != "1" || entry != "2" || entry != "3")) {
				if entry == "1" || entry == "2" {
					i, _ := strconv.Atoi(entry)
					if i == 1 {
						match = matches[0] + ")"
					} else {
						match = "(" + matches[1]
					}
					fmt.Printf("You typed %s%s%s and decided to set %s%s%s as true\n", white, entry, close, white, match, close)
				} else {
					fmt.Printf("You typed %s%s%s and decided to set %s%s%s and %s%s%s as true\n", white, entry, close, white, matches[0]+")", close, white, "("+matches[1], close)
				}
				break
			}
			fmt.Printf("Wrong input: Please type 1, 2 or 3\n")
		}
		if entry == "1" || entry == "2" {
			return match
		} else {
			return matches[0] + "+" + matches[1]
		}
	} else if char == "^" {
		fmt.Printf("There is an ambiguity in the following expression : %s%s %s %s%s\n", white, matches[0], char, matches[1], close)
		fmt.Printf("To fix the ambiguity, please choose the member you want to set as true\n")
		fmt.Printf("For %s%s%s type %s1%s or for %s%s%s type %s2%s\n", white, matches[0]+")", close, white, close, white, "("+matches[1], close, white, close)
		for true {
			fmt.Scanf("%s", &entry)
			if len(entry) == 1 && (entry == "1" || entry == "2") {
				i, _ := strconv.Atoi(entry)
				if i == 1 {
					match = matches[0] + ")"
				} else {
					match = "(" + matches[1]
				}
				fmt.Printf("You typed %s%s%s and decided to set %s%s%s as true\n", white, entry, close, white, match, close)
				break
			}
			fmt.Printf("Wrong input: Please type 1 or 2\n")
		}
		return match
	}
	return ""
}

func discardAmbiguity(ctx context.Context, right string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "discardAmbiguity")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if start, end, b := doParentheseNeedDiscard(ctx, right); b == true {
		tmp := discardAmbiguity(ctx, right[start:end+1])
		right = right[:start] + tmp + right[end+1:]
		if _, _, b := doParentheseNeedDiscard(ctx, right); b == true {
			right = discardAmbiguity(ctx, right)
		}
	}
	if b, matches := isRunePresent(ctx, right, "|"); b == true {
		right = manageAmbiguity(ctx, matches, "|")
		if b, _ := isRunePresent(ctx, right, "|"); b == true {
			right = discardAmbiguity(ctx, right)
		}
	}
	if b, matches := isRunePresent(ctx, right, "^"); b == true {
		right = manageAmbiguity(ctx, matches, "^")
		if b, _ := isRunePresent(ctx, right, "^"); b == true {
			right = discardAmbiguity(ctx, right)
		}
	}
	return right
}

func populateGraph(ctx context.Context, graph [][]node, left, right string) [][]node {
	span, _ := opentracing.StartSpanFromContext(ctx, "populateGraph")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if !regex(right, "^.*[A-Z]{1}.*[A-Z]{1}.*$") {
		if exclamationMark(ctx, right) {
			left, right = reverseMemberMark(ctx, left, right)
			graph = insertMemberInGraph(ctx, left, right, graph)
		} else {
			right = cleanMember(ctx, right)
			graph = insertMemberInGraph(ctx, left, right, graph)
		}
	} else {
		if regex(right, "^.*[\\^\\|]{1,}.*$") {
			right = discardAmbiguity(ctx, right)
			right = developNegation(ctx, right)
			reg := buildRegex(ctx, right, "([!]{0,1}[A-Z]{1})[+]{0,1}")
			matches := regexMatch(right, reg)
			for _, str := range matches[1:] {
				if exclamationMark(ctx, str) {
					tmp, str := reverseMemberMark(ctx, left, str)
					graph = insertMemberInGraph(ctx, tmp, str, graph)
				} else {
					graph = insertMemberInGraph(ctx, left, str, graph)
				}
			}
		} else {
			right := developNegation(ctx, right)
			reg := buildRegex(ctx, right, "([!]{0,1}[A-Z]{1})[+]{0,1}")
			matches := regexMatch(right, reg)
			for _, str := range matches[1:] {
				if exclamationMark(ctx, str) {
					tmp, str := reverseMemberMark(ctx, left, str)
					graph = insertMemberInGraph(ctx, tmp, str, graph)
				} else {
					graph = insertMemberInGraph(ctx, left, str, graph)
				}
			}
		}
	}
	return graph
}

func buildRegex(ctx context.Context, right, partReg string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "buildRegex")
	defer span.Finish()
	count := 0
	i := 0
	for _, char := range right {
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}
	reg := "^"
	for i != count {
		reg += partReg
		i++
	}
	reg += "$"
	return reg
}

func processRules(ctx context.Context, rules, facts []string) [][]node {
	span, _ := opentracing.StartSpanFromContext(ctx, "processRules")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	var graph [][]node
	for _, rule := range rules {
		matches := regexMatch(rule, "(^[()!A-Z+\\^\\|]{1,})([<]{0,}[=>]{2})(.*)$")
		left := matches[1]
		operator := matches[2]
		right := matches[3]
		if operator == "=>" {
			graph = populateGraph(ctx, graph, left, right)
		} else if operator == "<=>" {
			graph = populateGraph(ctx, graph, left, right)
			graph = populateGraph(ctx, graph, right, left)
		}
	}
	return graph
}

func processFacts(ctx context.Context, graph [][]node, facts []string) [][]node {
	span, _ := opentracing.StartSpanFromContext(ctx, "generateGraph")
	defer span.Finish()
	for _, fact := range facts {
		if nodeMap[fact][1] == 0 {
			newNode := []node{{false, fact, true, false}}
			graph = append(graph, newNode)
			nodeMap[fact] = [2]int{len(graph) - 1, 1}
		} else {
			graph[nodeMap[fact][0]][0].valid = true
		}
	}
	return graph
}

func generateGraph(ctx context.Context, rules, facts []string) [][]node {
	span, _ := opentracing.StartSpanFromContext(ctx, "generateGraph")
	defer span.Finish()
	span.LogFields(
		log.String("Rules", fmt.Sprintf("%v\n", rules)),
		log.String("Facts", fmt.Sprintf("%v\n", facts)),
	)
	ctx = opentracing.ContextWithSpan(ctx, span)
	graph := processRules(ctx, rules, facts)
	graph = processFacts(ctx, graph, facts)
	return graph
}
