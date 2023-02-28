package web

import (
	"checkVerification/model"
	X "checkVerification/utils"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Website struct {
}

func (b *Website) GetVerify(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(r.Body)

	parse, err := X.Parse(r.Body)
	if err != nil {
		_, _ = w.Write(X.JSON(X.Z{
			"code": http.StatusUnauthorized,
			"msg":  "Invalid request",
		}))
		return
	}

	var recordId int
	sqlStr := "select id from check_info where uuid=?"
	err = model.DB.QueryRow(sqlStr, r.Header.Get("uuid")).Scan(&recordId)
	if err != nil {
		if err == sql.ErrNoRows {
			sqlStr := "insert into check_info(uuid, check_type,check_value,update_time) values (?,?,?,?)"
			ret, err := model.DB.Exec(sqlStr, r.Header.Get("uuid"), parse["checkType"], "111333", time.Now().Format(X.TimeFormat))
			if err != nil {
				_, _ = w.Write(X.JSON(X.Z{
					"code": http.StatusBadRequest,
					"msg":  err.Error(),
				}))
				return
			}
			_, err = ret.LastInsertId() // 新插入数据的id
			if err != nil {
				_, _ = w.Write(X.JSON(X.Z{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				}))
				return
			}

			_, _ = w.Write(X.JSON(X.Z{
				"code": http.StatusOK,
				"data": X.Z{
					"test": parse,
				},
			}))
		} else {
			_, _ = w.Write(X.JSON(X.Z{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			}))
			return
		}
	} else {
		sqlStr := "update check_info set check_type=?,check_value=?,check_success=?,update_time=? where id = ?"
		ret, err := model.DB.Exec(sqlStr, parse["checkType"], "asdas", "n", time.Now().Format(X.TimeFormat), recordId)
		if err != nil {
			_, _ = w.Write(X.JSON(X.Z{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			}))
			return
		}
		_, err = ret.RowsAffected() // 操作影响的行数
		if err != nil {
			_, _ = w.Write(X.JSON(X.Z{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			}))
			return
		}

		_, _ = w.Write(X.JSON(X.Z{
			"code": http.StatusOK,
			"data": X.Z{
				"test": parse,
			},
		}))
	}
}
