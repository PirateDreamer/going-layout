package rest

import (
	"context"
	"going-demo/application"
	"going-demo/pkg/ginc"

	"github.com/gin-gonic/gin"
)

func UserRouter(userRest *UserRest) {
	g := ginc.R.Group("api/user")
	g.GET("/user_info", ginc.Run(userRest.UserInfo))
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
	ID int64 `form:"id" binding:"required"` // 用户ID
}

type UserInfoResp struct {
	Name string `json:"name"` // 姓名
	Age  int    `json:"age"`  // 年龄
}

// @Summary ID获取用户信息
// @Tags 用户
// @Description ID获取用户信息
// @Param req query UserInfoReq true "查询参数"
// @Success 200 {object}  ginc.Response{data=UserInfoResp}
// @Router /api/user/user_info [get]
func (userRest *UserRest) UserInfo(ctx context.Context, c *gin.Context, req UserInfoReq) (resp *UserInfoResp, err error) {
	return
}
