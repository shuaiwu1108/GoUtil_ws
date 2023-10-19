# GoUtil_ws
#### Go语言，工具包，自用

### 1. OrcleDBUtil  
#### 依赖github.com/sijms/go-ora/v2  

* OrclConnInit, oracle连接初始化,返回*sql.DB  
* OrclCustomQuery, 返回[]map[string]interface{}  
* OrclQueryOne, 返回map[string]string  
* OrclQueryAll, 返回[]map[string]string

### 2.MySqlDBUtil
#### 依赖github.com/go-sql-driver/mysql、github.com/jmoiron/sqlx
* MySqlConnInit, mysql连接初始化, 返回*sqlx.DB
* QueryOne, 返回map[string]string
* QueryAll, 返回[]map[string]string

### 3.sql文件夹说明
* 文件夹内部不要直接放入sql脚本  
* 建立二级目录后，在放入sql脚本  
* 如果有多个脚本，请保证文件排序，以便升序执行  
* 执行方法为DBUtil.ReadSqlExec()

### 4. CommonUtil
* 提供Pause方法，用于暂停当前窗口而不退出



 