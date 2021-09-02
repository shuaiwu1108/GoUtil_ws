package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wslio/GoUtil_ws/model"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	db *sqlx.DB
)

func DbInit(dbInfo model.DBInfo) {
	database, err := sqlx.Open("mysql", dbInfo.Name+":"+dbInfo.Pass+"@tcp("+dbInfo.Url+":"+strconv.Itoa(dbInfo.Port)+")/"+dbInfo.DbName+"?charset=utf8&parseTime=true")
	HandleError(err, "open mysql failed, ")
	db = database
	fmt.Println("Mysql 初始化完毕！", dbInfo.Url)
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

func ExecSql(sqlStr string, args ...interface{}){
	rs,err := db.Exec(sqlStr, args...)
	HandleError(err, "MySql执行出错！")
	num, _ := rs.RowsAffected()
	id,_ := rs.LastInsertId()
	fmt.Println("影响行数：", num, " id：", id)
}

func PrepareSql(sqlString string){
	rs,err := db.Prepare(sqlString)
	HandleError(err, "MySql执行出错！")
	rs.Exec()
}

func ReadSqlExec(){
	dir := GetIniVal("sql", "dir")
	fmt.Printf("Sql执行目录：%s\n", dir)

	fileInfoList, err := ioutil.ReadDir(dir)
	HandleError(err, "Sql目录读取失败！")

	reg := regexp.MustCompile(`{.*}`)

	for i := range fileInfoList{
		fmt.Printf("Sql文件名：%s\n", fileInfoList[i].Name())
		bytes, err := ioutil.ReadFile(filepath.Join(dir, fileInfoList[i].Name()))
		HandleError(err, fileInfoList[i].Name()+" 文件读取失败！")
		allSql := string(bytes)
		sqlArr := strings.Split(allSql, "@repeat")
		for i := range sqlArr{
			sql := sqlArr[i]
			if len(sql)==0 {
				continue
			}
			result := reg.FindAllString(sql, -1)
			if len(result)==0{
				tmpSql := sql[strings.Index(sql,"}")+1:]
				fmt.Println("执行sql：", tmpSql)
				ExecSql(tmpSql)
			}else{
				tmpRes, err := QueryOne(result[0][1:len(result[0])-1])
				HandleError(err, "条件Sql执行失败！")
				coun := tmpRes["count(1)"]
				fmt.Println("条件sql结果：", coun, "跳过执行")
				if("1" == coun){
					continue
				}else{
					tmpSql := sql[strings.Index(sql,"}")+1:]
					fmt.Println("执行sql：", tmpSql)
					ExecSql(tmpSql)
				}
			}
		}
	}
}