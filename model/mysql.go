package model

import (
	"checkVerification/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func InitDB() (err error) {
	DB, err = sql.Open("mysql", utils.GetConfig().Mysql)
	if err != nil {
		return err
	}

	checkTableSQL := "SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = ? AND table_name = ? LIMIT 1"
	var tableName string
	err = DB.QueryRow(checkTableSQL, "check_verification", "check_info").Scan(&tableName)
	if err != nil {
		if err == sql.ErrNoRows {
			createTableSQL := `CREATE TABLE check_info (
								  id int unsigned NOT NULL AUTO_INCREMENT,
								  uuid char(36) NOT NULL,
								  check_type varchar(255) NOT NULL,
								  check_value varchar(255) DEFAULT NULL,
								  check_success char(1) NOT NULL DEFAULT 'n',
								  update_time datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
								  PRIMARY KEY (id)
								) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci`

			_, err = DB.Exec(createTableSQL)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for range ticker.C {
			result, err := DB.Exec("DELETE FROM check_info WHERE update_time < DATE_SUB(NOW(), INTERVAL 24 HOUR)")
			if err != nil {
				log.Printf("删除冗余记录错误: %s \n", err.Error())
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Printf("删除冗余记录错误: %s \n", err.Error())
			}

			log.Printf("<%d>条冗余记录已删除", rowsAffected)
		}
	}()

	return nil
}
