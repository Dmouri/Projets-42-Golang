package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

var dgraph bool
var tracer opentracing.Tracer
var closer io.Closer

type Entry struct {
	UID  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
	Fact []F    `json:"fact,omitempty"`
}

type F struct {
	UID  string `json:"uid,omitempty"`
	Tag  string `json:"tag,omitempty"`
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
	Rule []R    `json:"rule,omitempty"`
}

type R struct {
	UID  string `json:"uid,omitempty"`
	Tag  string `json:"tag,omitempty"`
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout if b = true

func InitJaeger(service string, b bool) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           b,
			LocalAgentHostPort: "127.0.0.1:5775",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Printf("ERROR: cannot init Jaeger: %v\n", err)
	}
	return tracer, closer, err
}

func pushGraph(ctx context.Context, graph [][]node) bool {
	span, _ := opentracing.StartSpanFromContext(ctx, "pushGraph")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if multiple { // Wait end of transaction dgraph
		time.Sleep(1000000000)
	}
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Dgraph is not run, please run with \"make dgraph\"\n")
	}
	defer conn.Close()
	Dclient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	op := &api.Operation{}
	op.Schema = `
	name: string @index(term,fulltext,trigram) .
	id: string @index(term,fulltext,trigram) .
	tag: string @index(term,fulltext,trigram) .
	rule: uid @reverse .
	fact: uid @reverse .
	`
	err = Dclient.Alter(ctx, op)
	if err != nil {
		fmt.Printf("Dgraph is not run, please run with \"make dgraph\"\n")
		return false
	}
	rID := randomString()
	p := generateDgraph(ctx, graph, rID)
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
	}
	txn := Dclient.NewTxn()
	defer txn.Discard(ctx)
	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You can visualize the graph at : http://127.0.0.1:8000 with query :\n")
	fmt.Printf("{\n\tgraph(func: eq(name, \"Entry\")) @filter(eq(id, \"%s\")) @recurse(depth: 4, loop: false) {\n\t\texpand(_all_)\n\t}\n}\n", rID)
	return true
}

func generateDgraph(ctx context.Context, graph [][]node, rID string) Entry {
	span, _ := opentracing.StartSpanFromContext(ctx, "generateDgraph")
	defer span.Finish()
	ret := Entry{
		UID:  "_:entry:" + rID,
		Name: "Entry",
		ID:   rID,
	}
	Facts := []F{}
	for key := range nodeMap {
		fact := F{"_:" + graph[nodeMap[key][0]][0].name + rID, "fact", key, rID, nil}
		newR := []R{}
		for _, k := range graph[nodeMap[key][0]][1:] {
			rule := R{"_:" + k.name + ":" + rID, "rule", k.name, rID}
			newR = append(newR, rule)
		}
		if len(newR) != 0 {
			fact.Rule = newR
		}
		Facts = append(Facts, fact)
	}
	ret.Fact = Facts
	return ret
}
