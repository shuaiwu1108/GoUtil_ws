package util

import (
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"os"
	"strconv"
)

var (
	iniFile *ini.File
)

func GetIniInt(sectionName string, key string) int {
	tempVal := GetIniVal(sectionName, key)
	value, e := strconv.Atoi(tempVal)
	HandleError(e, "string convert string fail, key:"+key)
	return value
}

func GetIniVal(sectionName string, key string) string {
	var tempVal string
	section := GetSection(sectionName)
	if nil != section {
		tempVal = section.Key(key).Value()
	}
	fmt.Printf("配置文件[%s],Key:[%s],value:[%s]\n", sectionName, key, tempVal)
	return tempVal
}

func GetSection(sectionName string) *ini.Section {
	section, e := iniFile.GetSection(sectionName)
	HandleError(e, "未找到配置信息")
	return section
}

func ReadIniInit(fileName string) {
	file, e := ini.Load(fileName)
	HandleError(e, "conf.ini配置文件解析失败")
	iniFile = file
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteByteArray(filename string, tmp string) {
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	HandleError(err1, "")
	n, err1 := io.WriteString(f, tmp) //写入文件(字符串)
	HandleError(err1, "")
	fmt.Printf("写入 %d 个字节 \n", n)
}
