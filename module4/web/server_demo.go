package web

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    if err != nil {
        log.Println("error: ", err)
    }

    fmt.Println(request.Form)
    fmt.Println("path: ", request.URL.Path)
    fmt.Println("scheme: ", request.URL.Scheme)
    fmt.Println(request.Form["url_long"])
    for k, v := range request.Form {
        fmt.Println("key: ", k)
        fmt.Println("val: ", strings.Join(v, " "))
    }
    _, _ = fmt.Fprintf(writer, "hello chain!")
}

func StartServerDemo()  {
    http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal(err)
    }
}