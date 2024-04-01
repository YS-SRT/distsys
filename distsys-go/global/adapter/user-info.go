package adapter

import "distsys/base"

type UserInfo struct {
	ID       int
	Name     string
	Phone    string
	Password string
	Age      int
	Portrait string

	AccessToken  string
	RefreshToken string
}

type UserInfoAdapter interface {
	GetUserInfoByUIDS(uids string) (*UserInfo, *base.CustomError)
	GetUserInfoByID(uid int) (*UserInfo, *base.CustomError)
}
