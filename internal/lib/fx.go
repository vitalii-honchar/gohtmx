package lib

import (
	"fmt"

	"go.uber.org/fx"
)

func AsGroup(f any, t any, group string) any {
	return fx.Annotate(
		f,
		fx.As(t),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, group)),
	)
}
