package http

import "testing"

func TestGetRequest(t *testing.T) {
    GetRequest("http://www.baidu.com")
}

func TestPostRequest(t *testing.T) {
    PostRequest("http://www.baidu.com")
}