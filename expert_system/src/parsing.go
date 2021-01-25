package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func manageParenthese(ctx context.Context, str string) bool {
	span, _ := opentracing.StartSpanFromContext(ctx, "manageParenthese")
	defer span.Finish()
	span.LogFields(
		log.String("String", str),
	)
	if regex(str, "^.*[+\\^\\|!][)]{1,}.*$|^.*[(]{1,}[+\\^\\|]{1}.*$") {
		return false
	}
	countOpen := 0
	countClose := 0
	for _, elem := range str {
		if elem == '(' {
			countOpen++
		} else if elem == ')' {
			countClose++
			if countClose > countOpen {
				return false
			}
		}
	}
	if countOpen != countClose {
		return false
	}
	return true
}

func verifyFormat(ctx context.Context, str string) bool {
	span, _ := opentracing.StartSpanFromContext(ctx, "verifyFormat")
	defer span.Finish()
	span.LogFields(
		log.String("String", str),
	)
	if !regex(str, "^[(]{0,}[+\\^\\|]{1,}.*$|^.*[!+\\|\\^]{1,}[)]{0,}$") && !regex(str, "^.*[A-Z]{2,}.*$") && !regex(str, "^.*[!+\\|\\^]{1}[()]{0,}[+\\|\\^].*$|^.*[!]{1,}[!]{1,}.*$") {
		return true
	}
	return false
}

func checkDuplicateEntry(ctx context.Context, facts, queries []string) bool {
	span, _ := opentracing.StartSpanFromContext(ctx, "checkDuplicateEntry")
	defer span.Finish()
	span.LogFields(
		log.String("Check duplicate in :", fmt.Sprintf("%v", facts)),
		log.String("Check duplicate in :", fmt.Sprintf("%v", queries)),
	)
	exit := false
	encounteredMapFact := map[string]bool{}
	encounteredMapQuery := map[string]bool{}
	for _, key := range facts {
		if encounteredMapFact[key] == false {
			encounteredMapFact[key] = true
		} else {
			fmt.Printf("Error: Duplicate fact\n")
			exit = true
		}
	}
	for _, key := range queries {
		if encounteredMapQuery[key] == false {
			encounteredMapQuery[key] = true
		} else {
			fmt.Printf("Error: Duplicate query\n")
			exit = true
		}
	}
	if exit == true {
		return false
	}
	return true
}

func splitStr(ctx context.Context, purifyLine, char string) []string {
	span, _ := opentracing.StartSpanFromContext(ctx, "splitStr")
	defer span.Finish()
	span.LogFields(
		log.String("String", purifyLine),
		log.String("Char", char),
	)
	myReg := "^[" + char + "]{1}"
	for i := 0; i < (len(purifyLine) - 1); i++ {
		myReg += "([A-Z]{1})"
	}
	myReg += "$"
	tab := regexMatch(purifyLine, myReg)
	return tab[1:]
}

func parseFile(ctx context.Context, path string) ([]string, []string, []string, bool) {
	span, _ := opentracing.StartSpanFromContext(ctx, "parseFile")
	defer span.Finish()
	countFact := false
	countQuery := false
	line := 0
	errorCount := 0
	var rulesArray []string
	var factArray []string
	var queryArray []string
	ctx = opentracing.ContextWithSpan(ctx, span)
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		trimLine := strings.Replace(scanner.Text(), " ", "", -1)
		matches := regexMatch(trimLine, "^([^#]{0,})[#]{0,1}.{0,}$")
		if len(matches) == 1 || (len(matches) == 2 && len(matches[1]) == 0) {
			continue
		}
		purifyLine := matches[1]
		if countFact == false {
			if regex(purifyLine, "^[=]{1}[A-Z]{0,}$") {
				countFact = true
				factArray = splitStr(ctx, purifyLine, "=")
			} else if regex(purifyLine, "^[^<]{1,}[<]{0,1}[=]{1}[>]{1}.{1,}$") {
				matchSides := regexMatch(purifyLine, "^([^=<>]{1,})[<]{0,}[=]{1}[>]{1}([^=<>]{1,})$")
				left := matchSides[1]
				right := matchSides[2]
				if regex(left, "^[A-Z()!+\\^\\|]{1,}$") && regex(right, "^[A-Z()!+\\^\\|]{1,}$") {
					if manageParenthese(ctx, left) && manageParenthese(ctx, right) {
						if verifyFormat(ctx, left) && verifyFormat(ctx, right) {
							rulesArray = append(rulesArray, purifyLine)
						} else {
							fmt.Printf("Error: line %d: Wrong format\n", line)
							errorCount++
						}
					} else {
						fmt.Printf("Error: line %d: Wrong format\n", line)
						errorCount++
					}
				} else {
					fmt.Printf("Error: line %d: Wrong format\n", line)
					errorCount++
				}
			} else {
				fmt.Printf("Error: line %d: Wrong format\n", line)
				errorCount++
			}
		} else if countQuery == false {
			if regex(purifyLine, "^[?]{1}[A-Z]{1,}$") {
				countQuery = true
				queryArray = splitStr(ctx, purifyLine, "?")
			} else {
				fmt.Printf("Error: line %d: Wrong format\n", line)
				errorCount++
			}
		} else {
			fmt.Printf("Error: line %d: Wrong format\n", line)
			errorCount++
		}
	}
	if errorCount != 0 {
		return nil, nil, nil, false
	}
	if len(queryArray) == 0 {
		fmt.Printf("Error: No query found\n")
		return nil, nil, nil, false
	}
	if !checkDuplicateEntry(ctx, factArray, queryArray) {
		return nil, nil, nil, false
	}
	if debug {
		fmt.Printf("Rules:\n%v\nFact:\n%v\nQuery:\n%v\n", rulesArray, factArray, queryArray)
	}
	return rulesArray, factArray, queryArray, true
}
