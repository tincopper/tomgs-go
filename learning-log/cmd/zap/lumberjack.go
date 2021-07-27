package main

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "time"
)

/*
   uber日志框架示例：https://github.com/uber-go/zap
   学习文档：https://studygolang.com/articles/18780?fr=sidebar
            https://studygolang.com/articles/17394
*/
var MainLogger *zap.Logger
var GatewayLogger *zap.Logger

func init() {
    MainLogger = NewLogger("./logs/main.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
    GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
}

func main() {
    fmt.Println("init main")
    MainLogger.Debug("hello main Debug")
    MainLogger.Info("hello main Info")
    GatewayLogger.Debug("Hi Gateway Im Debug")
    GatewayLogger.Info("Hi Gateway  Im Info")
    MainLogger.Info("testtttt",
        zap.String("key1", "value1"),
        zap.Int("attempt", 3),
        zap.Duration("backoff", time.Second))
}

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
    core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
    return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
    //日志文件路径配置2
    hook := lumberjack.Logger{
        Filename:   filePath,   // 日志文件路径
        MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: maxBackups, // 日志文件最多保存多少个备份
        MaxAge:     maxAge,     // 文件最多保存多少天
        Compress:   compress,   // 是否压缩
    }
    // 设置日志级别
    atomicLevel := zap.NewAtomicLevel()
    atomicLevel.SetLevel(level)
    //公用编码器
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,  // 大写编码器，LowercaseLevelEncoder：小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }
    return zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
        atomicLevel,                                                                     // 日志级别
    )
}
