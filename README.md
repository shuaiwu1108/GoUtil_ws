# GoUtil_ws
Go语言，工具包，自用
### 1. OrcleDBUtil
1.需要安装oracle客户端，并配置系统环境变量CGO_LDFLAGS
2.windows下需要安装mingw环境
3.OracleInit
4.OracelQueryOne
5.OracelQueryAll

### 2. CommonUtil
1.提供Pause方法，用于暂停当前窗口而不退出

### 3.DBUtil
1.MySql的链接操作封装
2.DbInit
3.QueryOne
4.QueryAll

### 4.sql文件夹说明
1.文件夹内部不要直接放入sql脚本
2.建立二级目录后，在放入sql脚本
3.如果有多个脚本，请保证文件排序，以便升序执行
4.执行方法为DBUtil.ReadSqlExec()