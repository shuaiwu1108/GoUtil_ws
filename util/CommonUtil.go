package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Pause() {
	fmt.Printf("按任意键退出...")
	b := make([]byte, 1)
	os.Stdin.Read(b)

	/**
	fmt.Println("--------------------------------------")
	var s string
	fmt.Println("输入exit退出:")
	fmt.Scan(&s)
	if(s == "exit") {
	} else {
	pause()
	}
	*/
}

func HandleError(err error, message string, flag bool) {
	if err != nil {
		log.Println(message, err)
		if flag {
			Pause()
			os.Exit(500)
		} else {
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
