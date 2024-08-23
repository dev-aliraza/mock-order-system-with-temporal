package service

import (
	"context"
	"log"
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-api/model"
	"go.temporal.io/sdk/client"
)

type OrderService struct {
	workflowClient client.Client
}

type OrderServiceProvider interface {
	Create() interface{}
}

func NewOrderService(workflowClient client.Client) *OrderService {
	orderService := &OrderService{
		workflowClient: workflowClient,
	}
	return orderService
}

func (s *OrderService) Create(orderReq model.CreateOrderRequest) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	workflowExecution, err := s.workflowClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        orderReq.Schedule,
		TaskQueue: "create-order",
	}, "CreateOrderWorkflow")
	if err != nil {
		log.Fatalln("Execution error", err)
	}
	var result interface{}
	if orderReq.Schedule == "now" {
		workflowExecution.Get(ctx, &result)
	}
	return result
}

func (s *OrderService) Update(orderReq model.UpdateOrderRequest) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	workflowExecution, err := s.workflowClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        orderReq.Schedule,
		TaskQueue: "update-order",
	}, "UpdateOrderWorkflow")
	if err != nil {
		log.Fatalln("Execution error", err)
	}
	var result interface{}
	if orderReq.Schedule == "now" {
		workflowExecution.Get(ctx, &result)
	}
	return result
}
