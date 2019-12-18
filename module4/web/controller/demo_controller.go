package controller

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func init() {
    //web.AddRoute("/demo", demoHandler)
}

func demoHandler(writer http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    // Form可以获取到body和url中的参数
    values := request.Form
    name := values.Get("name")

    fmt.Println("request param name:: ", name)
    result, _ := json.Marshal(resultStruct{Status: true, Msg: name})
    _, err := writer.Write(result)
    if err != nil {
        log.Printf("writer data error, msg: %s", err)
    }
}