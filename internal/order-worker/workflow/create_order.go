package workflow

import (
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/activity"
	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/model"
	"go.temporal.io/sdk/workflow"
)

// define the workflow function
func CreateOrderWorkflow(ctx workflow.Context) (model.CreateOrderResponse, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 60,
	})
	// start the activities
	orderFuture := workflow.ExecuteActivity(ctx, activity.CreateOrderActivity)

	// wait for activities to complete
	order := model.CreateOrderResponse{}
	if err := orderFuture.Get(ctx, &order); err != nil {
		return order, err
	}
	return order, nil
}
