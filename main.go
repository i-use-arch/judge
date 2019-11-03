package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/i-use-arch/judge/dbconn"
	"github.com/i-use-arch/judge/runner"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	dbc, err := dbconn.MakeClient(ctx)

	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}

	runner := runner.Runner{Client: dbc, Timeout: 2 * time.Second}
	fmt.Println("about to run")
	output, err := runner.Run(3308020806989400253)

	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}

	fmt.Println(output)
}
