package activity

import (
	"context"
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/model"
)

func CreateOrderActivity(ctx context.Context) (order model.CreateOrderResponse, err error) {
	time.Sleep(5 * time.Second)
	return model.CreateOrderResponse{Id: 1, Item: "Pizza", Quantity: 2}, nil
}
