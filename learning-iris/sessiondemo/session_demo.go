package main

import (
    "fmt"
    "github.com/kataras/iris"
    "github.com/kataras/iris/sessions"
)

var (
    cookieNameForSessionID ="mycookesinahduixu"
    sess = sessions.New(sessions.Config{Cookie:cookieNameForSessionID})
)


func secret(ctx iris.Context){
    if auth, err := sess.Start(ctx).GetBoolean("authennticated"); !auth{
        fmt.Println(auth, err)
        ctx.StatusCode(iris.StatusForbidden)
        return
    }
    ctx.WriteString("the cake is a lie!")
}

func login(c iris.Context){
    session := sess.Start(c)
    
    session.Set("authennticated", true)
    
    c.WriteString("logging")
}

func loginOut(c iris.Context){
    session := sess.Start(c)
    
    // 撤销用户身份验证
    session.Set("authenticated", false)
}

// curl -XGET -s http://localhost:8080/secret
// curl -XGET -s -I http://localhost:8080/login
// curl -XGET -s -I http://localhost:8080/logout
// curl -XGET -s --cookie "mycookesinahduixu=719d169c-7209-4af8-9080-1fe69d8e7582" http://localhost:8080/secret
func main(){
    app	:= iris.New()
    
    app.Get("/secret", secret)
    app.Get("/login", login)
    app.Get("/loginOut", loginOut)
    app.Run(iris.Addr(":8080"))
}

