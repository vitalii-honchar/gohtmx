package web

import (
	"go-htmx/internal/infra/web/view/user"
	"go-htmx/internal/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(asMiddleware(NewGinLoggerMiddleware)),
	fx.Provide(asHandler(user.NewProfileView)),
	fx.Provide(NewRenderer),
	fx.Provide(fx.Annotate(
		NewGinServer,
		fx.ParamTags(`group:"middleware"`, `group:"handler"`),
	)),
)

func asHandler(f any) any {
	return lib.AsGroup(f, new(Handler), "handler")
}

func asMiddleware(f any) any {
	return lib.AsGroup(f, new(Middleware), "middleware")
}
