package main

import (
    "fmt"
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
    "github.com/kataras/iris/core/router"
    "github.com/kataras/iris/middleware/recover"
    "github.com/kataras/iris/mvc"
    "github.com/kataras/iris/sessions"
    "time"
    "tomgs-go/learning-iris/mvcdemo/datasource"
    "tomgs-go/learning-iris/mvcdemo/repositories"
    "tomgs-go/learning-iris/mvcdemo/services"
    "tomgs-go/learning-iris/mvcdemo/web/controller"
    "tomgs-go/learning-iris/mvcdemo/web/middleware"
)

/**
* @Author: tangzy
* @Date: 2020/1/20 14:20
 */
func main() {
    app := iris.New()
    // 服务器恢复机制
    app.Use(recover.New())
    app.Logger().SetLevel("debug")
    // Load the template files.
    app.RegisterView(iris.HTML("./web/views", ".html"))
    // 静态文件
    app.StaticWeb("/static", "./web/static")
    // 创建datasource、repositories、service
    // 注入service实例
    // Bind the "userService" to the UserController's Service (interface) field.
    //hero.Register()
    // 404处理
    app.OnErrorCode(iris.StatusNotFound, func(c context.Context) {
        _, _ = c.HTML("<h1>404</h1>")
    })
    //app.OnError(iris.StatusNotFound, func(c *iris.Context) { c.HTML(iris.StatusOK, "<h1>404</h1>") })
    // API
    //handler.RegisterAPI(app)
    
    // 通过这种配置的方式添加路由
    mvc.Configure(app.Party("/movies"), movies)
    //mvc.Configure(app, movies)
    
    // http://localhost:8080/movies
    // http://localhost:8080/movies/1
    app.Run(
        // Start the web server at localhost:8080
        iris.Addr("localhost:8080"),
        // skip err server closed when CTRL/CMD+C pressed:
        iris.WithoutServerError(iris.ErrServerClosed),
        // enables faster json serialization and more:
        iris.WithOptimizations,
    )
}

func movies(movies *mvc.Application) {
    // Add the basic authentication(admin:password) middleware
    // for the /movies based requests.
    movies.Router.Use(middleware.BasicAuth)
    //movies.Router.Use()
    //movies.Router.PartyFunc()
    
    sessionManager := sessions.New(sessions.Config{
        Cookie:       "site_session_id",
        Expires:      60 * time.Second,
        AllowReclaim: true,
    })
    
    // Create our movie repository with some (memory) data from the datasource.
    repo := repositories.NewMovieRepository(datasource.Movies)
    // Create our movie service, we will bind it to the movie movies's dependencies.
    movieService := services.NewMovieService(repo)
    movies.Register(sessionManager.Start, movieService)
    
    // serve our movies controller.
    // Note that you can serve more than one controller
    // you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(movies.Party(...))`
    // if you want.
    
    // error
    /*movies.Party("/findAll", func(context context.Context) {
        context.WriteString("ok")
    })*/
    // http://localhost:8080/movies/findAll
    movies.Party("/findAll").Handle(new(subController))
    
    // http://localhost:8080/movies/list/1
    movies.Router.PartyFunc("/list", func(child router.Party) {
        child.Get("/{id:int}", func(c context.Context) {
            c.WriteString(fmt.Sprintf("%s", c.Params().Get("id")))
        })
    })

    // http://localhost:8080/movies/list1
    movies.Router.Get("/list1", func(c context.Context) {
        c.WriteString("list1")
    })
    
    movies.Handle(new(controllers.MovieController))
}

type subController struct {

}

func (c *subController) Get() string {
    return "ok"
}