package main

import (
    "github.com/donnie4w/go-logger/logger"
    "tomgs-go/learning-logger/dep"
)

func init() {
    logger.SetConsole(true)
    logger.SetRollingFile(`F:\logtest`, "test.log", 10, 1, logger.KB)
    logger.SetLevel(logger.DEBUG)
}

func main() {
    logger.Debug("Debug log ...")
    logger.Info("Info log ...")
    logger.Warn("Warn log ...")
    logger.Error("Error log ...")
    logger.Fatal("Fatal log ...")
    
    // 测试依赖的包是否生效
    dep.TestDep()
}
