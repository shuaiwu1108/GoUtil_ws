package util

import (
	"net/http"
	"net/url"
)

func HttpGet(uu string, params url.Values) string {
	reqUrl, _ := url.ParseRequestURI(uu)
	reqUrl.RawQuery = params.Encode()
	resp, err := http.Get(reqUrl.String())
	HandleError(err, "get fail ", false)
	tmp := HandleResp(resp)
	return tmp
}

func HttpPost(uu string, data string) string {
	resp, err := http.PostForm(uu, url.Values{"jsonrequest": {data}})
	HandleError(err, uu + " POST请求失败", false)
	tmp := HandleResp(resp)
	return tmp
}