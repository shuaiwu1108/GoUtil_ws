package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleError(err error, message string, flag bool) {
	if err != nil {
		fmt.Println(message, err)

		if flag {
			Pause()
			os.Exit(500)
		}else{
			//什么都不做
		}

	}
}

func HandleResp(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "[Http返回结果处理失败]", false)
	return string(body)
}
