package handlers

import (
	"handsongolang/pkg/utils"
	"handsongolang/pkg/user"
	"net/http"
	// "fmt"
	"log"
)

type UserHandler struct {
	userService  user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService : userService,
	}
}


func (uh *UserHandler) CreateUser(resp http.ResponseWriter, req *http.Request) {
	var ur user.UserRequest
	err := utils.ParseRequest(req, &ur)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("request parse error"))
		return
	}
	newUser,_:= user.NewUser(ur)
	log.Printf("ur.FullName=====>"+newUser.FullName)
	res, _ := uh.userService.CreateUser(req.Context(),&newUser)
	userResponse:=user.UserResp{
		UserID  : res.ID,
		Message : "user created successfully",
	}
	// resp.Write([]byte("request parse error"))
	utils.WriteSuccessResponse(resp,http.StatusCreated,userResponse)
}