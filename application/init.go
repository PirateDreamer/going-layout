package application

import "go.uber.org/fx"

func App() fx.Option {
	return fx.Module("application",
		fx.Provide(
			NewUserApp,
		),
	)
}
