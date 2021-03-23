package user

// import (
// 	"fmt"
// )

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserResp struct {
	UserID  int
	Message string
}

// func AddUserSuccessMessage(name string) string {
// 	return fmt.Sprintf("User details for %s added successfully", name)
// }
