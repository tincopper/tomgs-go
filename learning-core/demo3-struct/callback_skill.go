package demo3_struct

import "fmt"

type Configurator func(*Application)

type Application struct {
    // 设置一些属性
    // configs *Configuration
    // ...
    AppName string
}

func New() *Application {
    app := &Application{
    }
    return app
}

func (app *Application) Configure(configurators ...Configurator) *Application {
    for _, cfg := range configurators {
        cfg(app)
    }
    return app
}

func (app *Application) SetAppName(appName string) {
    app.AppName = appName
}

// 自定义的回调函数，相当于Spring里面的ApplicationContextAware里面的setApplicationContext方法
// 这里可以定义为一个模板
func CustomConfigurator(app *Application)  {
    app.SetAppName("demo-base-app")
}

func TestCallback() {
    app := New()
    app.Configure(CustomConfigurator)
    fmt.Println(app.AppName)
}