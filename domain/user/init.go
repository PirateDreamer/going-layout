package user

import (
	"going-demo/domain/user/repo"
	"going-demo/domain/user/service"

	"go.uber.org/fx"
)

func UserDomain() fx.Option {
	return fx.Options(
		fx.Provide(repo.NewUserRepo),
		fx.Provide(service.NewUserService),
	)
}
