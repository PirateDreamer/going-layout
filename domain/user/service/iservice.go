package service

import (
	"context"
	"going-demo/domain/user/model/entity"
)

type IUserService interface {
	UserInfo(context.Context, int64) (*entity.User, error)
}
