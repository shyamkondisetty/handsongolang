package user

import (
	// "context"
	"context"
	// "handsongolang/pkg/config"
	// "handsongolang/pkg/db"
	"log"
)

type Service interface {
	CreateUser(ctx context.Context,user *User) (*User, error)
}

type userService struct {
	// ctx context.Context
	userRepo UserRepository
}

func (uss userService) CreateUser(ctx context.Context,user *User) (*User, error) {

	log.Printf("first Name  "+user.FirstName)
	log.Printf("last Name  "+user.LastName)
	uss.userRepo.AddUser(ctx,user)
	return user, nil

}
func NewUserService(userRepo UserRepository) Service {
	return &userService{
		userRepo: userRepo,
	}
}
