package internal

import (
	"go-htmx/internal/app"
	"go-htmx/internal/infra/config"
	"go-htmx/internal/infra/database"
	"go-htmx/internal/infra/web"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		web.Module,
		database.Module,
		fx.Provide(
			config.NewConfig,
			app.NewGetUserUseCase,
			zap.NewExample,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*gin.Engine) {}),
	)
}
