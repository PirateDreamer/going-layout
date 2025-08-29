package domain

import (
	"going-demo/domain/user"

	"go.uber.org/fx"
)

func Domain() fx.Option {
	return fx.Module("domain",
		user.UserDomain(),
	)
}
