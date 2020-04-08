package main

/**
* @Author: tangzy
* @Date: 2020/4/7 11:16
 */
import (
    "fmt"
    "github.com/donnie4w/go-logger/logger"
    "strconv"
    "time"
)

func _log(i int) {
    logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
    //	logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
    //	logger.Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
    //	logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
    //	logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
}

func main() {
    // go 日志框架使用
    //runtime.GOMAXPROCS(runtime.NumCPU())
    
    //指定是否控制台打印，默认为true
    //	logger.SetConsole(true)
    //	logger.SetFormat("=====>%s##%s")
    //指定日志文件备份方式为文件大小的方式
    //第一个参数为日志文件存放目录
    //第二个参数为日志文件命名
    //第三个参数为备份文件最大数量
    //第四个参数为备份文件大小
    //第五个参数为文件大小的单位
    logger.SetRollingFile(`F:\logtest`, "test.log", 10, 1, logger.KB)
    
    //指定日志文件备份方式为日期的方式
    //第一个参数为日志文件存放目录
    //第二个参数为日志文件命名
    //	logger.SetRollingDaily(`C:\Users\Thinkpad\Desktop\logtest`, "test.log")
    
    //指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
    //一般习惯是测试阶段为debug，		 生成环境为info以上
    logger.SetLevel(logger.DEBUG)
    
    for i := 100; i > 0; i-- {
        go _log(i)
    }
    time.Sleep(2 * time.Second)
    var lg = logger.GetLogger()
    
    //重新指定log文件
    lg.SetRollingFile(`F:\logtest`, "test.log", 10, 1, logger.KB)
    lg.SetLevelFile(logger.INFO, `F:\logtest`, "info.log")
    lg.SetLevelFile(logger.WARN, `F:\logtest`, "warn.log")
    lg.Debug("debug hello world")
    for i := 100; i > 0; i-- {
        go lg.Info("info hello world >>>>>>>>>>>>>>>>>> ", i)
    }
    lg.Warn("warn hello world")
    lg.Error("param1 ", 1, ", param2", 2)
    lg.Error("param1: %s, param2: %s", 1, 2)
    
    time.Sleep(2 * time.Second)
    
    lg.Debug(fmt.Sprintf("param1: %s, param2: %s","abc", "123"))
}
