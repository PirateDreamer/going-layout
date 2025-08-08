package rest

import (
	"context"
	"going-demo/application"
	"going-demo/pkg/ginc"

	"github.com/gin-gonic/gin"
)

func UserRouter(userRest *UserRest) {
	g := ginc.R.Group("api/user")
	g.GET("user_info", ginc.Run(userRest.UserInfo))
}

type UserRest struct {
	userApp *application.UserApp
}

func NewUserRest(userApp *application.UserApp) *UserRest {
	return &UserRest{
		userApp: userApp,
	}
}

type UserInfoReq struct {
	ID int64 `json:"id"`
}

type UserInfoResp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (userRest *UserRest) UserInfo(ctx context.Context, c *gin.Context, req UserInfoReq) (resp *UserInfoResp, err error) {
	return
}
