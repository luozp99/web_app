package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"web_app/settings"
)

var db *sqlx.DB

func InitDB(conf *settings.MySqlConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName,
	)

	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		return err
	}

	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	return
}

func Close() {
	db.Close()
}

func GetDb() *sqlx.DB {
	return db
}
