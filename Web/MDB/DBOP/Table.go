package DBOP

import (
	"database/sql"
	"fmt"
	"go-proxy-cache/Web/MDB/Base"
	"reflect"
	"strconv"
	"strings"
)

var (
	TableFieldTypeInt           = "INTEGER"
	TableFieldTypeDouble        = "DOUBLE"
	TableFieldTypeCHARACTER36   = "CHARACTER(36)"
	TableFieldTypeCHARACTER256  = "CHARACTER(256)"
	TableFieldTypeCHARACTER2048 = "CHARACTER(2048)"
)

type TableField struct {
	Name       string
	Type       string
	PrimaryKey bool
	NotNull    bool
}

func (tf TableField) toCreateSql() string {
	i := 0
	if tf.PrimaryKey {
		i += 1
	}
	if tf.NotNull {
		i += 2
	}
	if i == 0 {
		return fmt.Sprintf("`%s` %s ", tf.Name, tf.Type)
	} else if i == 1 {
		return fmt.Sprintf("`%s` %s primary key", tf.Name, tf.Type)
	} else if i == 2 {
		return fmt.Sprintf("`%s` %s not null", tf.Name, tf.Type)
	} else if i == 3 {
		return fmt.Sprintf("`%s` %s primary key not null", tf.Name, tf.Type)
	} else {
		return ""
	}
}

type TableFields []TableField

func (tfs TableFields) toCreateSql() string {
	tfsql := []string{}
	for _, tf := range tfs {
		tfsql = append(tfsql, tf.toCreateSql())
	}
	return strings.Join(tfsql, ",")
}
func (tfs TableFields) toSelectSql() string {
	sSql := []string{}
	for _, tf := range tfs {
		sSql = append(sSql, tf.Name)
	}
	return "`" + strings.Join(sSql, "`,`") + "`"
}

type Table struct {
	Name   string
	Fields TableFields
}

func (t Table) BuildCreateSql() string {
	return fmt.Sprintf("create table if not exists %s(%s)", t.Name, t.Fields.toCreateSql())
}
func (t Table) SelectAll(db *sql.DB, where string) [][]string {
	selectSql := fmt.Sprintf("select %s from %s %s", t.Fields.toSelectSql(), t.Name, where)
	rows, _ := Base.ExecuteWithResult(db, selectSql, true)
	//  [][]string
	//rets := Base.Row2Json(rows, len(t.Fields))
	return Base.Row2Json(rows, len(t.Fields))
}
func toInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 10)
	return n
}
func toDouble(s string) float64 {
	f, _ := strconv.ParseFloat(s, 10)
	return f
}
func (t Table) SelectResultToObj(obj interface{}, content []string) {
	ps := reflect.ValueOf(obj)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		for find, field := range t.Fields {
			f := s.FieldByName(field.Name)
			if f.IsValid() && f.CanSet() {
				if field.Type == TableFieldTypeInt {
					f.SetInt(toInt(content[find]))
				} else if field.Type == TableFieldTypeCHARACTER36 ||
					field.Type == TableFieldTypeCHARACTER256 ||
					field.Type == TableFieldTypeCHARACTER2048 {
					f.SetString(content[find])
				} else if field.Type == TableFieldTypeDouble {
					f.SetFloat(toDouble(content[find]))
				}
			}
		}
	}
}
