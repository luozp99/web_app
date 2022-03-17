package dao

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/dao/mysql"
	"web_app/modles"
)

var secret string = "aa123456"

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

func QueryUserByNameAndPwd(name string, pwd string) {

}
