package logrus

import log "github.com/sirupsen/logrus"

type DefaultFieldHook struct {
}

/*
    添加appName字段名称
 */
func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
    entry.Data["appName"] = "MyAppName"
    return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
    return log.AllLevels
}

