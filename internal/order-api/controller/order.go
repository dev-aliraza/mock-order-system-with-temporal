package controller

import (
	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-api/model"
	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-api/service"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

type OrderController struct {
	workflowClient client.Client
}

func NewOrderController(workflowClient client.Client) *OrderController {
	orderController := &OrderController{workflowClient: workflowClient}
	return orderController
}

func (oc *OrderController) Create(ctx *gin.Context) {
	var orderReq model.CreateOrderRequest
	ctx.BindJSON(&orderReq)
	service := service.NewOrderService(oc.workflowClient)
	ctx.JSON(200, service.Create(orderReq))
}

func (oc *OrderController) Update(ctx *gin.Context) {
	var orderReq model.UpdateOrderRequest
	ctx.BindJSON(&orderReq)
	service := service.NewOrderService(oc.workflowClient)
	ctx.JSON(200, service.Update(orderReq))
}
