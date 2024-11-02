package database

import (
	"go-htmx/internal/app"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(
		newUserRepository,
		fx.As(new(app.UserStorage)),
	)),
)
