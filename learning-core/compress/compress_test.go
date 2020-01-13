package compress

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCompress(t *testing.T) {
    err := Compress("F:\\test.txt", "F:\\test.zip")
    if err != nil {
        fmt.Println(err)
    }
}

func TestCompress2(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        request.ParseForm()
        Compress2("F:\\test.txt", writer, request)
    }))
    //server.URL += "?filepath=F:/"

    resp, err := http.Get(server.URL)
    if err != nil {
        fmt.Println(err)
    }

    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    err = ioutil.WriteFile("F:\\test.zip", bodyBytes, 777)
    if err != nil {
        fmt.Println(err)
    }
    //result := string(bodyBytes)
    //fmt.Println(result)

    //result := resultDataStruct{}
    //err = json.Unmarshal(bodyBytes, &result)

}