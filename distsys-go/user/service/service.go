package service

import (
	"distsys/base"
	"distsys/global"
	"distsys/global/adapter"
	"distsys/global/config"
	"distsys/user/model"
	"distsys/utils"

	"gorm.io/gorm"
)

type UserService struct {
	config *config.ConfigDoc
	db     *gorm.DB
}

func BuildService(config *config.ConfigDoc, db *gorm.DB) *UserService {

	return &UserService{
		config: config,
		db:     db,
	}
}

func (svr *UserService) GetUserInfoByID(uid int) (*adapter.UserInfo, *base.CustomError) {

	if user, err := svr.GetUserByID(uid); err != nil {
		return nil, err
	} else {

		if user.IsActive && !user.IsDeleted {

			return &adapter.UserInfo{

				ID:       user.ID,
				Name:     user.Nickname,
				Phone:    user.Phone,
				Age:      user.Age,
				Portrait: user.Portrait,
			}, nil

		} else {

			return nil, base.BuildCustomError(base.CodeUserNotActiveOrIsDeleted, "user is inactive or be deleted")
		}

	}

}

func (svr *UserService) GetUserInfoByUIDS(uids string) (*adapter.UserInfo, *base.CustomError) {

	if user, err := svr.GetUserByWechatIDS(uids); err != nil {
		return nil, err
	} else {

		if user.IsActive && !user.IsDeleted {

			return &adapter.UserInfo{

				ID:       user.ID,
				Name:     user.Nickname,
				Phone:    user.Phone,
				Age:      user.Age,
				Portrait: user.Portrait,
			}, nil

		} else {

			return nil, base.BuildCustomError(base.CodeUserNotActiveOrIsDeleted, "user is inactive or be deleted")
		}

	}

}

func (svr *UserService) CreateUser(u *model.User) (*model.User, *base.CustomError) {

	result := global.GSQLDB.Create(&u)
	if err := utils.VerifyDBResult(result.Error, result.RowsAffected, "create user failed"); err != nil {
		return nil, err
	}

	return u, nil
}

func (svr *UserService) GetUserByPwd(phone string, password string) (*model.User, *base.CustomError) {
	var user *model.User
	result := global.GSQLDB.Where(&model.User{Phone: phone, Password: utils.AESEncode4Conf(password)}).First(&user)

	if err := utils.VerifyDBResult(result.Error, result.RowsAffected, "Get user by pwd failed"); err != nil {

		return nil, err
	}

	return user, nil

}

func (svr *UserService) GetUserByID(uid int) (*model.User, *base.CustomError) {
	var user *model.User
	result := global.GSQLDB.Where(&model.User{ID: uid}).First(&user)

	if err := utils.VerifyDBResult(result.Error, result.RowsAffected, "Get user by id failed"); err != nil {

		return nil, err
	}

	return user, nil

}

func (svr *UserService) GetUserByWechatIDS(uids string) (*model.User, *base.CustomError) {
	var user *model.User
	result := global.GSQLDB.Where(&model.User{WechatUIDS: uids}).First(&user)

	if err := utils.VerifyDBResult(result.Error, result.RowsAffected, "Get user by id failed"); err != nil {

		return nil, err
	}

	return user, nil

}
