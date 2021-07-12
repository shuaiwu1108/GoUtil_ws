package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleError(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		Pause()
		os.Exit(500)
	}
}

func HandleResp(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "[Http返回结果处理失败]")
	return string(body)
}
