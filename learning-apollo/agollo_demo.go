package main

import (
    "fmt"
    "github.com/shima-park/agollo"
    "github.com/shima-park/agollo/viper-remote"
    "github.com/spf13/viper"
    "time"
)

// https://github.com/shima-park/agollo
func Demo1() {
    a, err := agollo.New("172.20.183.155:8080", "panshi-agent-dev", agollo.AutoFetchOnCacheMiss())
    if err != nil {
        panic(err)
    }
    errCh := a.Start()
    value := a.Get("apollo.spring.datasource.username")
    fmt.Println(value)

    // 监听配置变化
    watchCh := a.Watch()
    for {
        select {
            case err := <- errCh:
                //
                fmt.Println(err)
            case resp := <- watchCh:
                fmt.Println(
                    "error:", resp.Error,
                    "changes:", resp.Changes,
                    "namespace:", resp.Namespace,
                    "newValue:", resp.NewValue,
                    "oldValue:", resp.OldValue)

        }
    }
}

type Config struct {
    AppSalt string         `mapstructure:"appsalt"`
    DB      DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
    Driver   string        `mapstructure:"driver"`
    Host     string        `mapstructure:"host"`
    Port     int           `mapstructure:"port"`
    Timeout time.Duration  `mapstructure:"timeout"`
    // ...
}

type DemoConfig struct {
    PanshiAgent PanshiAgentConfig `mapstructure:"panshi-agent"`
}

type PanshiAgentConfig struct {
    DumpDir string `mapstructure:"dump_dir"`
    BinDump string `mapstructure:"bin_dump"`
    BinBash string `mapstructure:"bin_bash"`
    AdminUrl string `mapstructure:"admin_url"`
    Listen string `mapstructure:"listen"`
    LogDir string `mapstructure:"log_dir"`

}

func Demo2() {
    remote.SetAppID("panshi-agent-dev")
    v := viper.New()
    v.SetConfigType("prop")
    err := v.AddRemoteProvider("apollo", "172.20.183.155:8080", "application")
    // error handle...
    err = v.ReadRemoteConfig()
    // error handle...
    if err != nil {
        fmt.Println(err)
    }

    // 直接反序列化到结构体中
    /*var conf DemoConfig
    err = v.Unmarshal(&conf)
    // error handle...
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%+v\n", conf)*/

    // 各种基础类型配置项读取
    fmt.Println("listen:", v.GetString("panshi-agent.listen"))
    fmt.Println("dump_dir:", v.GetString("panshi-agent.dump_dir"))
    fmt.Println("bin_dump:", v.GetString("panshi-agent.bin_dump"))
    fmt.Println("bin_bash:", v.GetString("panshi-agent.bin_bash"))
    fmt.Println("admin_url:", v.GetString("panshi-agent.admin_url"))
    fmt.Println("log_dir:", v.GetString("panshi-agent.log_dir"))

    fmt.Println("panshi-agent:", v.GetStringMap("panshi-agent"))

    // 获取所有key，所有配置
    fmt.Println("AllKeys", v.AllKeys(), "AllSettings",  v.AllSettings())

    for {
        time.Sleep(1 * time.Second)
        _ = v.WatchRemoteConfig()
        v.WatchRemoteConfigOnChannel()
        //fmt.Println("app.AllSettings:", v.AllSettings())
        fmt.Println("listen:", v.GetString("panshi-agent.listen"))
    }
}