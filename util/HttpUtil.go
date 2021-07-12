package util

import (
	"net/http"
	"net/url"
)

func HttpGet(uu string, params url.Values) string {
	reqUrl, _ := url.ParseRequestURI(uu)
	reqUrl.RawQuery = params.Encode()
	resp, err := http.Get(reqUrl.String())
	HandleError(err, "get fail ")
	tmp := HandleResp(resp)
	return tmp
}
