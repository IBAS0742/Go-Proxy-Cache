package Base

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

var FieldTypeText = "TEXT"
var FieldTypeINTEGER = "INTEGER"
var FieldTypeREAL = "REAL"
var FieldTypeVARCHAR = "VARCHAR"
var FieldTypeBLOB = "BLOB"

// 获取将 string 类型转为具体的 表字段对应的类型
func getTableFieldTypeInstanceFunc(fieldType string) (f func(string) interface{}) {
	switch fieldType {
	case FieldTypeVARCHAR:
		f = func(s string) interface{} {
			return s
		}
		break
	case FieldTypeText:
		f = func(s string) interface{} {
			return s
		}
		break
	case FieldTypeBLOB:
		f = func(s string) interface{} {
			return s
		}
		break
	case FieldTypeINTEGER:
		f = func(s string) interface{} {
			i, _ := strconv.Atoi(s)
			return i
		}
		break
	case FieldTypeREAL:
		f = func(s string) interface{} {
			//i,_ := strconv.Atoi(s)
			f, _ := strconv.ParseFloat(s, 64)
			return f
		}
		break
	}
	return f
}

func ExecuteSqlNotReturn(db *sql.DB, sql string, info bool) error {
	if info {
		fmt.Println(sql)
	}
	statement, err := db.Prepare(sql)
	if statement != nil {
		statement.Exec()
		defer statement.Close()
		return nil
	} else {
		//fmt.Println("sql fail")
		if err != nil {
			return err
		}
		return errors.New("sql execute fail")
	}
}
func ExecuteWithResult(db *sql.DB, sql string, info bool) (*sql.Rows, error) {
	if info {
		fmt.Println(sql)
	}
	rows, err_ := db.Query(sql)
	return rows, err_
}

func ExecuteCustomSelectSql(db *sql.DB, sql string, fieldLen int) [][]string {
	rows, _ := ExecuteWithResult(db, sql, true)
	return Row2Json(rows, fieldLen)
}
