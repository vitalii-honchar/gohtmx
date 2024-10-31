package internal

import (
	"fmt"
	"go-htmx/internal/app"
	"go-htmx/internal/infra/config"
	"go-htmx/internal/infra/database"
	"go-htmx/internal/infra/web"
	"go-htmx/internal/infra/web/view/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			asMiddleware(web.NewGinLoggerMiddleware),
			asHandler(user.NewProfileView),
			app.NewGetUserUseCase,
			fx.Annotate(
				database.NewUserRepository,
				fx.As(new(app.UserStorage)),
			),
			fx.Annotate(
				web.NewGinServer,
				fx.ParamTags(`group:"middleware"`, `group:"handler"`),
			),
			web.NewRenderer,
			zap.NewExample,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*gin.Engine) {}),
	)
}

func asMiddleware(f any) any {
	return asGroup(f, new(web.Middleware), "middleware")
}

func asHandler(f any) any {
	return asGroup(f, new(web.Handler), "handler")
}

func asGroup(f any, t any, group string) any {
	return fx.Annotate(
		f,
		fx.As(t),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, group)),
	)
}
