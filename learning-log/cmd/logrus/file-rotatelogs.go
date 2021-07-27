package main

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    log "github.com/sirupsen/logrus"
    "io"
    "os"
    "time"
    "tomgs-go/learning-log/logrus"
)

func init() {
    // 设置日志格式为json格式
    log.SetFormatter(&log.JSONFormatter{})

    // 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
    // 日志消息输出可以是任意的io.writer类型
    //log.SetOutput(os.Stdout)

    // 设置多种输出方式
    /*
       var file, err = os.OpenFile("F:\\logtest\\logrus.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
       if err != nil {
           fmt.Println("Could Not Open Log File : " + err.Error())
       }
       log.SetOutput(io.MultiWriter(file, os.Stdout))
    */

    // 日志切割方式
    path := "F:\\logtest\\logrus.log"
    /* 日志轮转相关函数
       `WithLinkName` 为最新的日志建立软连接
       `WithRotationTime` 设置日志分割的时间，隔多久分割一次
       WithMaxAge 和 WithRotationCount二者只能设置一个
         `WithMaxAge` 设置文件清理前的最长保存时间
         `WithRotationCount` 设置文件清理前最多保存的个数
    */
    // 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
    writer, _ := rotatelogs.New(
        path + ".%Y%m%d%H%M",
        rotatelogs.WithLinkName(path),
        rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
        rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
    )
    log.SetOutput(io.MultiWriter(writer, os.Stdout))

    // 设置日志级别为warn以上
    log.SetLevel(log.DebugLevel)
    log.SetReportCaller(true)
    // 添加Hook
    // hook可以添加自定义的信息
    log.AddHook(&logrus.DefaultFieldHook{})
    // 记录行号信息，这个要放在logrus包下面
    log.AddHook(&logrus.LineHook{
        Field: "linenum",
        Skip:  1,
    })
}

// https://blog.csdn.net/wslyk606/article/details/81670713
func main() {
    log.Infof("test log level: %s", "INFO")

    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")

    log.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
    }).Warn("The group's number increased tremendously!")

    log.WithFields(log.Fields{
        "omg":    true,
        "number": 100,
    }).Fatal("The ice breaks!")
}
