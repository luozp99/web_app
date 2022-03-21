package logic

import (
	userDao "web_app/dao/user"
	"web_app/modles"
	"web_app/pkg/snowflake"
)

func SignUp(user *modles.UserSignUp) (err error) {
	err = userDao.QueryUserByName(user.UserName)
	if err != nil {
		return err
	}

	userId := snowflake.GenId()

	var signUser = &modles.UserDO{
		Id:       userId,
		Name:     user.UserName,
		Password: user.Password,
	}

	return userDao.InsertUser(signUser)
}

func LoginUser(loginUser *modles.LoginUser) error {
	return userDao.QueryUserByNameAndPwd(loginUser.UserName, loginUser.Password)
}
