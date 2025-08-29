package main

import (
	"context"
	"fmt"
	"going-demo/application"
	"going-demo/domain"
	"going-demo/internal/rest"
	"going-demo/pkg/conf"
	"going-demo/pkg/infra/store"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// @title going-layout
// @version 1.0
// @description going的脚手架
func main() {
	fx.New(
		fx.Provide(conf.Load),
		fx.Provide(store.InitMysql),

		domain.Domain(),
		application.App(),
		rest.Rest(),

		fx.Invoke(func(conf *viper.Viper, userRest *rest.UserRest) {
			r := gin.Default()
			rest.UserRouter(r, userRest)

			r.Run()
		}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
