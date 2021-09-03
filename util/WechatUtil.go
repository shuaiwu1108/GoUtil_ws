package util

import (
	"encoding/json"
	"net/url"
)

func getToken(appid string, appsecret string) string {
	params := url.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", appid)
	params.Set("secret", appsecret)
	tmp := HttpGet("https://api.weixin.qq.com/cgi-bin/token", params)
	var f interface{}
	err := json.Unmarshal([]byte(tmp), &f)
	HandleError(err, "token json analysis fail", true)
	m := f.(map[string]interface{})
	tokenStr := m["access_token"]
	return tokenStr.(string)
}

func UserGet(appid string, appsecret string) {
	token := getToken(appid, appsecret)
	params := url.Values{}
	params.Set("access_token", token)
	tmp := HttpGet("https://api.weixin.qq.com/cgi-bin/user/get", params)
	WriteByteArray("UserList.json", tmp)
}
