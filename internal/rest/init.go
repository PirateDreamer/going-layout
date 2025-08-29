package rest

import (
	"go.uber.org/fx"
)

func Rest() fx.Option {
	return fx.Options(
		fx.Provide(NewUserRest),
	)
}

// func Router() fx.Option {
// 	return fx.Options(
// 		fx.Invoke(UserRouter),
// 	)
// }
