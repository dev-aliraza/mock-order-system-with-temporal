package activity

import (
	"context"
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/model"
)

func UpdateOrderActivity(ctx context.Context) (order model.UpdateOrderResponse, err error) {
	time.Sleep(5 * time.Second)
	return model.UpdateOrderResponse{Id: 1, Item: "Pizza2", Quantity: 2}, nil
}
