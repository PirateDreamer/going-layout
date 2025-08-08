package application

import (
	"context"
	"going-demo/domain/user/model/entity"
	"going-demo/domain/user/service"
)

type UserApp struct {
	userService service.IUserService
}

func NewUserApp(userService service.IUserService) *UserApp {
	return &UserApp{
		userService: userService,
	}
}

func (userApp *UserApp) UserInfo(ctx context.Context, ID int64) (user *entity.User, err error) {
	user, err = userApp.userService.UserInfo(ctx, ID)
	return
}
