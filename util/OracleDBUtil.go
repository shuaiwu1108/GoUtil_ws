package util

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"github.com/wslio/GoUtil_ws/model"
)

var oracledb *sql.DB

func OracleInit(dbInfo model.DBInfo) {
	osqlInfo := fmt.Sprintf("%s/%s@%s:%d/%s", dbInfo.Name, dbInfo.Pass, dbInfo.Url, dbInfo.Port, dbInfo.DbName)
	fmt.Println("链接Oarcle：", osqlInfo)
	dd, err := sql.Open("godror", osqlInfo)
	HandleError(err, "Oracle 链接失败", true)
	oracledb = dd
	fmt.Println("Oracle链接成功！")
}

func OracelQueryOne(sqlStr string, args ...interface{}) (map[string]string, error) {
	rows, err := oracledb.Query(sqlStr, args...)
	HandleError(err, "[Oracle SQL 查询出错]", true)
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
	HandleError(err, "[Oracle SQL查询，结果解析出错]", true)
	for k, v := range values{
		key := cols[k]
		row[key] = string(v)
	}
	return row, nil
}

func OracelQueryAll(sqlStr string, args ...interface{}) ([]map[string]string, error) {
	rows, err := oracledb.Query(sqlStr, args...)
	HandleError(err, "[Oracle SQL 查询出错]", true)
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
		HandleError(err, "[Oracle SQL 查询结果解析出错]", true)
		row := make(map[string]string, 10)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}