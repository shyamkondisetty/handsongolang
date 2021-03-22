package user

import (
	// "context"
	"context"
	"handsongolang/pkg/config"
	"handsongolang/pkg/db"
	"log"
)

type Service interface {
	CreateUser(user *User) (string, error)
}

type userService struct {
	ctx context.Context
}

func (uss *userService) CreateUser(user *User) (string, error) {
	// res, err := bis.repository.GetBankURL(ctx, bankCode)
	// if err != nil {
	// 	return "", liberror.Builder().SetOperation("Service.GetBankURL").SetCause(err).Build()
	// }
	log.Printf("first Name  "+user.FirstName)
	log.Printf("last Name  "+user.LastName)
	dc:=config.NewDBConfig()
	ndh,_:=db.NewDBHandler(dc).GetDB()
	NewUserRepository(ndh).AddUser(uss.ctx,user)
	return "in service  "+user.LastName, nil

}
func NewUserService(ctx context.Context) Service {
	return &userService{
		ctx : ctx,
	}
}
