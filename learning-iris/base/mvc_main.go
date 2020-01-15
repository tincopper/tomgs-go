package base

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
    "github.com/kataras/iris/mvc"
)

func IrisMvcMain() *iris.Application {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())
    mvc.New(app).Handle(new(ExampleController))
    // app.Run(iris.Addr(":8080"))
    return app
}

// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct {}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
    return mvc.Response{
        ContentType: "text/html",
        Text: "<h1>Welcome</h1>",
    }
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
    return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
    return map[string]string{"message": "Hello Iris!"}
}

