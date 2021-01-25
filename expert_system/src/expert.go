package main

import (
	"context"
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func mgtEntryParam(ctx context.Context) ([]string, bool) {
	span, _ := opentracing.StartSpanFromContext(ctx, "mgtEntryParam")
	defer span.Finish()
	var strFile []string
	nb := 0
	for i, str := range os.Args {
		if str == "" {
			continue
		}
		if i != 0 && (regex(str, "^--graph$") || regex(str, "^-g$")) {
			dgraph = true
		} else if i != 0 && (regex(str, "^--debug$") || regex(str, "^-d$")) {
			debug = true
		} else if i != 0 && (regex(str, "^--verbose$") || regex(str, "^-v$")) {
			verbose = true
		} else if i != 0 {
			file, err := os.Open(str)
			defer file.Close()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			nb++
			strFile = append(strFile, str)
		} else if i != 0 {
			nb++
		}
	}
	if nb == 0 {
		fmt.Printf("Error: Not enough argument\n")
		writeUsage()
		return nil, false
	}
	if len(strFile) > 1 {
		multiple = true
	}
	return strFile, true
}

func main() {
	var err error
	var checkError bool
	tracer, closer, err = InitJaeger("Expert System", false)
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		fmt.Printf("Jaeger Instance failed\n")
	}
	defer closer.Close()
	span := tracer.StartSpan("Main")
	defer span.Finish()
	span.LogFields(
		log.String("log", "Start main"),
	)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tabPath, valid := mgtEntryParam(ctx)
	cleanExit(valid, span)
	for _, path := range tabPath {
		checkError = false
		fmt.Printf("Process file : %s:\n", path)
		rules, facts, queries, valid := parseFile(ctx, path)
		if !valid {
			checkError = true
			continue
		}
		graph := generateGraph(ctx, rules, facts)
		if debug {
			fmt.Printf("Graph:\n%v\n", graph)
		}
		processQueries(ctx, graph, queries)
		if dgraph && len(graph) > 0 {
			b := pushGraph(ctx, graph)
			cleanExit(b, span)
		}
		nodeMap = make(map[string][2]int)
		currentQuery = ""
	}
	if checkError == true {
		cleanExit(false, span)
	}
	return
}
