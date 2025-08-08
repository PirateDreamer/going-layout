package repo

import (
	"context"
	"going-demo/domain/user/model/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (userRepo *UserRepo) UserByID(c context.Context, ID int64) (user *entity.User, err error) {
	result := userRepo.db.WithContext(c).Where("id = ?", ID).Limit(1).Find(&user)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
