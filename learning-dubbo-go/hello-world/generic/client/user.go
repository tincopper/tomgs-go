package main

type HelloRequest struct {
    Name string
}

func (u *HelloRequest) JavaClassName() string {
    return "com.tomgs.learning.dubbo.api.Helloworld$HelloRequest"
}