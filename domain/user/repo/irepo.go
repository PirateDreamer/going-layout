package repo

import (
	"context"
	"going-demo/domain/user/model/entity"
)

type IUserRepo interface {
	UserByID(c context.Context, ID int64) (*entity.User, error)
}
