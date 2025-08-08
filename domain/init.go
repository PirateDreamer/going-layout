package domain

import (
	"going-demo/domain/user"

	"go.uber.org/fx"
)

func Domain() fx.Option {
	return fx.Options(
		user.UserDomain(),
	)
}
