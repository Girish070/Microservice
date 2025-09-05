package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Girish070/go-microservice/application"
)

func main() {
	app := application.New()
	ctx, cancle := signal.NotifyContext(context.Background(),os.Interrupt)
	defer cancle()
	err  := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}