
package user
import (
	"net/http"
	"fmt"
	"handsongolang/pkg/utils"
	"log"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
	}
}


func (uh *UserHandler) CreateUser(resp http.ResponseWriter, req *http.Request) {
	var ur UserRequest
	err := utils.ParseRequest(req, &ur)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("request parse error"))
		return
	}
	user,_:= NewUser(ur)
	log.Printf("ur.FullName=====>"+user.FullName)
	res, _ := NewUserService(req.Context()).CreateUser(&user);
	fmt.Fprintf(resp, res)
}