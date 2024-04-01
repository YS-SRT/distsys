package adapter

import "distsys/base"

type UserAuthReq struct {
	UserId   string
	Resource string
	Act      string
}

type UserRole struct {
	ID   int
	Role string
}

type ResourceRole struct {
	ID           int
	Resource     string
	ResourceRole string
	Act          []string
}

type UserAuthorizedAdapter interface {
	// IsAdmin(auth *UserAuthReq) (bool, *base.CustomError)
	IsAuthorized(req *UserAuthReq) (bool, *base.CustomError)
}
