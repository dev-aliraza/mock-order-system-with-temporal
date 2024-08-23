package workflow

import (
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/activity"
	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-worker/model"
	"go.temporal.io/sdk/workflow"
)

// define the workflow function
func UpdateOrderWorkflow(ctx workflow.Context) (model.UpdateOrderResponse, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 60,
	})
	// start the activities
	orderFuture := workflow.ExecuteActivity(ctx, activity.UpdateOrderActivity)

	// wait for activities to complete
	order := model.UpdateOrderResponse{}
	if err := orderFuture.Get(ctx, &order); err != nil {
		return order, err
	}
	return order, nil
}
