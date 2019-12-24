package filetest

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

// 获取当前目录下的文件或目录信息(不包含多级子目录，不包含路径)
func ListDir(path string) {
    if path == "" {
        pwd, _ := os.Getwd()
        path = pwd
    }

    dirs, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for i := range dirs {
        fileInfo := dirs[i]
        fmt.Println(fileInfo.Name(), fileInfo.IsDir(), fileInfo.ModTime(), fileInfo.Size())
    }
}

// 获取当前目录下的文件或目录名(包含路径)
// E:\go_workspace\src\tomgs-go\module1\filetest\file_demo.go
func ListDir2() {
    pwd, _ := os.Getwd()
    filepathNames, err := filepath.Glob(filepath.Join(pwd, "*"))
    if err != nil {
        log.Fatal(err)
    }
    for i := range filepathNames {
        filepathName := filepathNames[i]
        fmt.Println(filepathName)
    }
}

// 获取当前文件或目录下的所有文件或目录信息(包括子目录)
func ListDir3() {
    pwd, _ := os.Getwd()
    err := filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        fmt.Println(info.Name())
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}

// 获取目录的绝对路径
func ListDir4() {
    tmp, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    current, _ := filepath.Abs(filepath.Dir(os.Args[1]))
    unknown, _ := filepath.Abs(filepath.Dir(os.Args[2]))
    fmt.Println(tmp)
    fmt.Println(current)
    fmt.Println(unknown)
}

func WriteData() {
    file, err := ioutil.ReadFile("F:\\test.txt")
    if err != nil {
        fmt.Println(err)
    }
    i := len(file)
    fmt.Println(i)
}