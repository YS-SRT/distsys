package convertor

import (
	"distsys/user/model"
	"distsys/user/model/dto"
	"distsys/utils"
)

func RegisterInput2User(input *dto.UserRegisterInput) *model.User {

	return &model.User{
		Nickname:   input.Nickname,
		WechatUIDS: input.WechatUIDS,
		Phone:      input.Phone,
		IsActive:   true,
		IsDeleted:  false,
		Password:   utils.AESEncode4Conf(input.Password),
	}

}

func User2VO(user *model.User, accessToken string, refreshToken string) *dto.UserVO {
	return &dto.UserVO{
		ID:           user.ID,
		Nickname:     user.Nickname,
		Portrait:     user.Portrait,
		Phone:        user.Phone,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
}
