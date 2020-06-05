package mylog

import (
    "log"
    "os"
)

// 自定义日志级别
var (
    Debug *log.Logger
    Info *log.Logger
    Warn *log.Logger
    Error *log.Logger
)

func init() {
    log.Println("my log init ...")
    Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate | log.Ltime | log.Lshortfile)
    Info  = log.New(os.Stdout, "[INFO] ", log.Ldate | log.Ltime | log.Lshortfile)
    Warn  = log.New(os.Stdout, "[WARN] ", log.Ldate | log.Ltime | log.Lshortfile)
    Error = log.New(os.Stderr, "[Error] ", log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
    Debug.Println("debug info...")
}