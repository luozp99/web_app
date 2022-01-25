package logic

import (
	userDao "web_app/dao/user"
	"web_app/modles"
	"web_app/pkg/snowflake"
)

func SignUp(user *modles.UserSignUp) {

	userDao.QueryUserByName(user.UserName)

	_ = snowflake.GenId()

	userDao.InsertUser()

}
