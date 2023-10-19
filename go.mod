module github.com/shuaiwu1108/GoUtil_ws

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/sijms/go-ora/v2 v2.7.19
	gopkg.in/ini.v1 v1.62.0
)

replace github.com/shuaiwu1108/GoUtil_ws => ../GoUtil_ws

require github.com/smartystreets/goconvey v1.6.4 // indirect
