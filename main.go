package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/i-use-arch/judge/dbconn"
	"github.com/i-use-arch/judge/runner"
	"github.com/i-use-arch/judge/workqueue"
)

var (
	mongoStr  = flag.String("mongostr", "", "connection string for mongodb")
	rabbitStr = flag.String("rabbitstr", "", "connection string for rabbitmq")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	dbc, err := dbconn.MakeClient(ctx, *mongoStr)

	if err != nil {
		failOnError(err, "unable to connect to db")
	}

	runner := runner.Runner{Client: dbc, Timeout: 2 * time.Second}

	workqueue.MakeQueue(ctx, *rabbitStr, func(id uint64) error {
		fmt.Println("about to run")
		output, err := runner.Run(id)

		fmt.Println(output)

		var status string
		if err != nil {
			status = "error"
		} else {
			status = "finished"
		}

		err = dbc.WriteOutput(id, output, status)

		if err != nil {
			log.Printf("%v", err)
		}
		return nil
	})

	forever := make(chan bool)
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
