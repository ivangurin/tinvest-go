package bot

import (
	"context"
	"tinvest-go/internal/config"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/pkg/service_provider"
)

type IApp interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *service_provider.ServiceProvider
}

func NewApp(ctx context.Context) IApp {
	ctx, cancel := context.WithCancel(ctx)

	sp := service_provider.GetServiceProvider(ctx)
	sp.GetCloser().Add(func() error {
		cancel()
		return nil
	})

	config.Init()

	return &app{
		ctx: ctx,
		sp:  sp,
	}
}

func (a *app) Run() error {
	logger.Info(a.ctx, "app is starting...")
	defer logger.Info(a.ctx, "app has been finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	go func() {
		logger.Info(a.ctx, "bot is starting...")
		err := a.sp.GetBotApi(a.ctx).Serve(a.ctx)
		if err != nil {
			logger.Errorf(a.ctx, "failed to serve bot: %s", err.Error())
			closer.CloseAll()
		}
	}()

	// logger
	// closer.Add(logger.Close)

	// tracer
	// closer.Add(tracer.Close)

	logger.Info(a.ctx, "app started successfully")

	return nil
}
