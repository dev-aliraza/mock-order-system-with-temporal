package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dev-aliraza/mock-order-system-with-temporal/internal/order-api/controller"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

func health(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Working Fine.....!")
}

func createWorkflowClient() client.Client {
	c, err := client.Dial(client.Options{
		HostPort:  "127.0.0.1:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client:::::", err)
	}
	return c
}

func registerControllers(engine *gin.Engine, c client.Client) {
	oc := controller.NewOrderController(c)
	engine.GET("/", health)
	engine.POST("/v1/order", oc.Create)
	engine.PATCH("/v1/order", oc.Update)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	workflowClient := createWorkflowClient()
	defer workflowClient.Close()

	router := gin.Default()
	registerControllers(router, workflowClient)

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancelTimeout := context.WithTimeout(ctx, 5)
	defer cancelTimeout()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("fatal shutdown: %s", err)
		os.Exit(1)
	}

	fmt.Println("Shutting down server")
}
