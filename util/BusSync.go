package util

import (
	"GoUtil_ws/model"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func BusSyncSql(code, vpn, port, sql string) string {
	pub := model.RequestInstance(code)
	req := model.IntegratedQueryRequest{Publicrequest: pub}
	req.Dbname = "1"
	req.Sql = sql
	req.Parameter = "{}"
	h := md5.New()
	h.Write([]byte(pub.Askorgcode+req.Sql+req.Dbname+req.Parameter))
	req.Publicrequest.Signmsg = hex.EncodeToString(h.Sum(nil))
	reqJson, err := json.Marshal(req)
	HandleError(err, "请求参数Json异常!")
	res := HttpPost("http://"+vpn+":"+port+"/integratedquery", string(reqJson))
	fmt.Println(res)
	return res
}
