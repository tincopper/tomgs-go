package controller

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type resultStruct struct {
    Status bool `json:"status"`
    Msg string `json:"msg"`
}

type resultDataStruct struct {
    Status bool `json:"status"`
    Msg string `json:"msg"`
    Data interface{} `json:"data"`
}

func init() {
    //web.AddRoute("/basic", basicHandler)
}

func basicHandler(writer http.ResponseWriter, request *http.Request) {
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