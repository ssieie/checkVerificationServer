package model

import (
	"checkVerification/utils"
	"database/sql"
)

var DB *sql.DB

func InitDB() (err error) {
	DB, err = sql.Open("mysql", utils.GetConfig().Mysql)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
