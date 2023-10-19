package util

import (
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	goora "github.com/sijms/go-ora/v2"
	"time"
)

func OrclConnInit(param map[string]interface{}) *sql.DB {
	url := goora.BuildUrl(param["host"].(string), param["port"].(int), param["database"].(string), param["username"].(string), param["password"].(string), nil)
	fmt.Println("链接Oarcle：", url)
	dd, err := sql.Open(param["drivername"].(string), url)
	HandleError(err, "Oracle 链接失败", true)
	fmt.Println("Oracle链接成功！")
	return dd
}

func OrclCustomQuery(oracledb *sql.DB, sqlString string) []map[string]interface{} {
	rows, _ := oracledb.Query(sqlString)
	defer rows.Close()
	var res []map[string]interface{}
	columns, _ := rows.Columns()
	vals := make([]interface{}, len(columns))
	valsPtr := make([]interface{}, len(columns))
	for i := range vals {
		valsPtr[i] = &vals[i]
	}
	for rows.Next() {
		_ = rows.Scan(valsPtr...)
		r := make(map[string]interface{})
		for i, v := range columns {
			if va, ok := vals[i].([]byte); ok {
				r[v] = string(va)
			} else {
				if v == "DEPARTDATE" {
					tmp := vals[i].(time.Time)
					r[v] = tmp.Format("2006-01-02")
				} else if _, ok := vals[i].(time.Time); ok {
					tmp := vals[i].(time.Time)
					r[v] = tmp.Format("2006-01-02 15:04:05")
				} else {
					r[v] = vals[i]
				}
			}
		}
		res = append(res, r)
	}
	return res
}

func OrclQueryOne(oracledb *sql.DB, sqlStr string, args ...interface{}) (map[string]string, error) {
	rows, err := oracledb.Query(sqlStr, args...)
	HandleError(err, "[Oracle SQL 查询出错]", true)
	defer rows.Close()
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	rows.Next()
	row := make(map[string]string)
	err = rows.Scan(scans...)
	HandleError(err, "[Oracle SQL查询，结果解析出错]", true)
	for k, v := range values {
		key := cols[k]
		row[key] = string(v)
	}
	return row, nil
}

func OrclQueryAll(oracledb *sql.DB, sqlStr string, args ...interface{}) ([]map[string]string, error) {
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
