package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
)

var debug bool
var verbose bool
var multiple bool

func randomString() string {
	msec := time.Now().UnixNano()
	str := strconv.Itoa(int(msec))
	return str[10:]
}

func regex(str, regex string) bool {
	valid, _ := regexp.Compile(regex)
	return valid.MatchString(str)
}

func regexMatch(str, regex string) []string {
	match, _ := regexp.Compile(regex)
	return match.FindStringSubmatch(str)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func getInnerParenthese(ctx context.Context, right string) (int, int) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getInnerParenthese")
	defer span.Finish()
	countOpen := -1
	countClose := -1
	for i, elem := range right {
		if elem == '(' {
			countOpen = i
		}
	}
	for i, elem := range right {
		if elem == ')' {
			countClose = i
			if countClose > countOpen {
				break
			}
		}
	}
	return countOpen, countClose
}

func writeUsage() {
	fmt.Printf("Usage: ./expert_system [--graph|-g] [--debug|-d] file.txt\n")
	fmt.Printf("\t\t--graph: Activate graph render in the UI\n")
	fmt.Printf("\t\t--debug: Print rules, facts and queries\n")
}

func cleanExit(exit bool, span opentracing.Span) {
	if !exit {
		span.Finish()
		closer.Close()
		os.Exit(1)
	}
}
