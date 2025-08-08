package service

import (
	"context"
	"going-demo/domain/user/model/entity"
	"going-demo/domain/user/repo"
	"going-demo/types/xerrs"
)

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) UserInfo(ctx context.Context, ID int64) (user *entity.User, err error) {
	if user, err = s.userRepo.UserByID(ctx, ID); err != nil {
		return
	}
	if user == nil {
		err = xerrs.ErrUserNotFound
		return
	}
	return user, nil
}
