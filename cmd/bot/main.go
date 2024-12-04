package main

import (
	"context"
	"tinvest-go/internal/app/bot"
	"tinvest-go/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	err := bot.NewApp(ctx).Run()
	if err != nil {
		logger.Fatalf(ctx, "failed to run bot: %s", err.Error())
	}
}
