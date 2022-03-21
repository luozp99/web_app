package dao

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/dao/mysql"
	"web_app/modles"
)

var secret = "aa123456"

func QueryUserByName(name string) error {
	strSql := "select count(id) from user where name = ?"
	db := mysql.GetDb()
	var count int
	if err := db.Get(&count, strSql, name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return nil
}

func InsertUser(user *modles.UserDO) (err error) {
	strSql := "insert into user(id,name,password) value(?,?,?)"
	user.Password = encryptPassword(user.Password)
	db := mysql.GetDb()
	_, err = db.Exec(strSql, user.Id, user.Name, user.Password)
	return err
}

func encryptPassword(password string) (pwd string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func QueryUserByNameAndPwd(name string, pwd string) (err error) {
	var password = encryptPassword(pwd)

	strSql := "select id,name,password from user where name = ?"
	db := mysql.GetDb()
	var user = &modles.UserDO{}
	err = db.Get(user, strSql, name)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}

	if err != nil {
		return err
	}

	if password != user.Password {
		return errors.New("密码输入错误")
	}
	return err
}
