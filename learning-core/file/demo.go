package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}

	return result, nil
}

func main() {
	//var files []string
	//files, _ = GetAllFile("learning-core")
	//fmt.Println("目录下的所有文件如下")
	//for i := 0; i < len(files); i++ {
	//	fmt.Println(files[i])
	//}

	extMap := make(map[string]string)
	err := json.Unmarshal([]byte("null"), &extMap)
	fmt.Println(err)
	result := extMap["a"]
	fmt.Println(result)
	extMap["a"] = "1"
}
