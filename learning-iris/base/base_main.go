package base

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
)

func IrisBaseMain() *iris.Application {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    app.Handle("GET", "/", func(ctx context.Context) {
        ctx.HTML("<h1>Welcome</h1>")
    })

    app.Get("/ping", func(ctx context.Context) {
        ctx.WriteString("pong")
    })

    app.Get("/hello", func(ctx context.Context) {
        ctx.JSON(iris.Map{"message": "Hello Iris!"})
    })

    return app
    /*err := app.Run(iris.Addr(":8080"))
    if err != nil {
        fmt.Println(err)
    }*/
}