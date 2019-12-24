package http

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func GetRequest(url string) {
    client := http.DefaultClient
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    defer func() {_ = resp.Body.Close()}()

    status := resp.Status
    fmt.Println(status)
    body := resp.Body
    all, _ := ioutil.ReadAll(body)
    fmt.Println(string(all))
}

type Test struct {
    Age string `json:"age"`
    Name string `json:"name"`
}

func PostRequest(url string) {
    client := http.DefaultClient
    // 字符串的形式
    //reader := strings.NewReader("")
    byteSlice, _ := json.Marshal(&Test{
        Age:  "18",
        Name: "tomgs",
    })
    reader := bytes.NewReader(byteSlice)
    resp, _ := client.Post(url, "application/json", reader)
    respBytes, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(respBytes))
}