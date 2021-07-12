package util

import (
	"fmt"
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
