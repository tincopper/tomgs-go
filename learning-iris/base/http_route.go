package base

import (
    "fmt"
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
    "github.com/kataras/iris/core/router"
)

// 路由分组
func RouteParty() *iris.Application{
    app := iris.New()
    // 方式一
    /*users := app.Party("/users", myAuthMiddlewareHandler)
    // http://localhost:8080/users/42/profile
    users.Get("/{id:int}/profile", userProfileHandler)
    // http://localhost:8080/users/messages/1
    users.Get("/messages/{id:int}", userMessageHandler)*/

    // 方式二
    app.PartyFunc("/users", func(users router.Party) {
        users.Use(myAuthMiddlewareHandler)
        // http://localhost:8080/users/42/profile
        users.Get("/{id:int}/profile", userProfileHandler)
        // http://localhost:8080/users/messages/1
        users.Get("/messages/{id:int}", userMessageHandler)
    })

    return app
}

/*var myAuthMiddlewareHandler = func(ctx context.Context) {
    ctx.HTML("<h1>Welcome</h1>")
}*/
func myAuthMiddlewareHandler(ctx context.Context)  {
    ctx.HTML("<h1>Welcome</h1>")
    // 验证通过则调用Next，否则调用StopExecution
    ctx.Next() // continue
    //ctx.StopExecution()
}

var userProfileHandler = func(ctx context.Context) {
    // post请求体参数
    body := ctx.Request().Body
    fmt.Println(body)
}

var userMessageHandler = func(ctx context.Context) {
    method := ctx.Method()
    path := ctx.Path()
    // 获取？后面的参数
    id := ctx.URLParamIntDefault("id", -1)
    // 获取获取path参数
    id2 := ctx.Params().GetIntDefault("id", -2)
    fmt.Println("method:", method, "path:", path, "id:", id, "id2", id2)
}
