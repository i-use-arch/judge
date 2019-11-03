package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/i-use-arch/judge/dbconn"
	"github.com/i-use-arch/judge/runner"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	dbc, err := dbconn.MakeClient(ctx)

	if err != nil {
		panic(err)
	}

	runner := runner.Runner{Client: dbc}
	fmt.Println("about to run")
	err = runner.Run(3308020806989400253)

	if err != nil {
		panic(err)
	}
}
