package application

import "go.uber.org/fx"

func App() fx.Option {
	return fx.Options(
		fx.Provide(NewUserApp),
	)
}
