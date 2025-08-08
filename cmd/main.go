package main

import (
	"going-demo/application"
	"going-demo/domain"
	"going-demo/internal/rest"
	"going-demo/pkg/conf"
	"going-demo/pkg/ginc"
	"going-demo/pkg/infra/store"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	ginc.R = gin.Default()

	fx.New(
		fx.Provide(conf.Load),
		fx.Provide(store.InitMysql),

		rest.Rest(),
		application.App(),
		domain.Domain(),

		rest.Router(),
		// fx.Invoke(func(conf *viper.Viper) {
		// 	ginc.R.Run()
		// }),
	).Run()
}
