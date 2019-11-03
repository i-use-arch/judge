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

	id := uint64(3308020806989400253)

	ctx := context.Background()
	dbc, err := dbconn.MakeClient(ctx)

	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}

	runner := runner.Runner{Client: dbc, Timeout: 2 * time.Second}
	fmt.Println("about to run")
	output, err := runner.Run(id)

	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}

	fmt.Println(output)

	var status string
	if err != nil {
		status = "error"
	} else {
		status = "finished"
	}

	err = dbc.WriteOutput(id, output, status)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
