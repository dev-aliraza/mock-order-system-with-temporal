package main

import (
	"log"
	"os"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/activity"
	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func createWorkflowClient() (client.Client, error) {
	c, err := client.Dial(client.Options{
		HostPort:  "127.0.0.1:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	return c, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("No worker name provided")
	}

	workflowClient, _ := createWorkflowClient()
	defer workflowClient.Close()

	var w worker.Worker

	if os.Args[1] == "create-order" {
		w = worker.New(workflowClient, "create-order", worker.Options{})
		w.RegisterWorkflow(workflow.CreateOrderWorkflow)
		w.RegisterActivity(activity.CreateOrderActivity)
	} else if os.Args[1] == "update-order" {
		w = worker.New(workflowClient, "update-order", worker.Options{})
		w.RegisterWorkflow(workflow.UpdateOrderWorkflow)
		w.RegisterActivity(activity.UpdateOrderActivity)
	}

	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker:::::", err)
	}
}
