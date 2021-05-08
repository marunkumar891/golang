package main

import (
	"context"
	"golang/temporal/worflows"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "GreetingTaskQueue",
	}
	name := "ArunKumar"
	we, err := c.ExecuteWorkflow(context.Background(), options, worflows.Greetings, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

}
