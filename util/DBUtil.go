package util

import (
	"../model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

var (
	db *sqlx.DB
)

func DbInit(dbInfo model.DBInfo) {
	database, err := sqlx.Open("mysql", dbInfo.Name+":"+dbInfo.Pass+"@tcp("+dbInfo.Url+":"+strconv.Itoa(dbInfo.Port)+")/"+dbInfo.DbName+"?charset=utf8&parseTime=true")
	HandleError(err, "open mysql failed, ")
	db = database
}

func QueryOne(sqlStr string, args ...interface{}) (map[string]string, error) {
	rows, err := db.Query(sqlStr, args...)
	HandleError(err, "[SQL 查询出错]")
	defer rows.Close()
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values{
		scans[i] = &values[i]
	}
	rows.Next()
	row := make(map[string]string)
	err = rows.Scan(scans...)
	HandleError(err, "[SQL查询，结果解析出错]")
	for k, v := range values{
		key := cols[k]
		row[key] = string(v)
	}
	return row, nil
}

func QueryAll(sqlStr string, args ...interface{}) ([]map[string]string, error) {
	rows, err := db.Query(sqlStr, args...)
	HandleError(err, "[SQL 查询出错]")
	defer rows.Close()
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0, 10)
	for rows.Next() {
		err := rows.Scan(scans...)
		HandleError(err, "[SQL 查询结果解析出错]")
		row := make(map[string]string, 10)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}
