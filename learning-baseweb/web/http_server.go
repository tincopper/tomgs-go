package web

import (
    "fmt"
    "log"
    "net/http"
    "runtime/debug"
    "strings"
)

type Handler interface {
    Callback(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

type FirstCallbackFun func(http.ResponseWriter, *http.Request) bool

func (f HandlerFunc) Callback(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}

var routeMap = make(map[string]Handler, 16)
var firstCallBack FirstCallbackFun

func SetFirstCallBack(callbackFunc func(http.ResponseWriter, *http.Request) bool) {
    if firstCallBack == nil {
        firstCallBack = callbackFunc
    }
}

func AddRoute(route string, callbackFunc func(http.ResponseWriter, *http.Request)) error {
    if _, ok := routeMap[route]; ok {
        return fmt.Errorf(route + " is exsit")
    }
    routeMap[route] = HandlerFunc(callbackFunc)
    http.HandleFunc(route, routeFunc)
    return nil
}

func routeFunc(w http.ResponseWriter, req *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("err:", err, string(debug.Stack()))
        }
    }()
    var route string
    i := strings.IndexAny(req.RequestURI, "?")
    if i > 0 {
        route = strings.TrimSpace(req.RequestURI[0:i])
    } else {
        route = req.RequestURI
    }
    if _, ok := routeMap[route]; ok {
        if firstCallBack(w, req) == true {
            routeMap[route].Callback(w, req)
        }
    }
}

func AddStaticRoute(route string, dir string) {
    http.Handle(route, http.FileServer(http.Dir(dir)))
}

func Start(ipAndPort string) error {
    return http.ListenAndServe(ipAndPort, nil)
}

func StartTLS(ipAndPort string, serverKey string, serverCrt string) error {
    return http.ListenAndServeTLS(ipAndPort, serverCrt, serverKey, nil)
}

func InitAndStartServe(firstCallback func(http.ResponseWriter, *http.Request) bool, loadRoute func(), ipAndPort string) error {
    SetFirstCallBack(firstCallBack)
    loadRoute()
    err := Start(ipAndPort)
    return err
}

func LoadRouteAndStartServe(loadRoute func(), ipAndPort string) error {
    SetFirstCallBack(func(writer http.ResponseWriter, request *http.Request) bool {
        return true
    })
    loadRoute()
    err := Start(ipAndPort)
    return err
}