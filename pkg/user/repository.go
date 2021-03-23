package user

import (
	"context"
	// "errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
	// "payment-svc/pkg/liberror"
	// "payment-svc/pkg/user/model"
	"time"
)

type UserRepository interface {
	AddUser(ctx context.Context,user *User) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func (ur *gormUserRepository) AddUser(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	db := ur.db.WithContext(ctx).Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}
