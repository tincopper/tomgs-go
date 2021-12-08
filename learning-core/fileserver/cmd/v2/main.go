package main

import (
	"log"
	"net/http"
)

// http.FileServer 方式
// https://darjun.github.io/2020/01/13/goweb/fileserver/
func main() {
	mux := http.NewServeMux()
	//http.Dir表示文件的起始路径，空即为当前路径。调用Open方法时，传入的参数需要在前面拼接上该起始路径得到实际文件路径。
	//
	//http.FileServer的返回值类型是http.Handler，所以需要使用Handle方法注册处理器。
	//http.FileServer将收到的请求路径传给http.Dir的Open方法打开对应的文件或目录进行处理。
	//在上面的程序中，如果请求路径为/static/hello.html，那么拼接http.Dir的起始路径.，最终会读取路径为./static/hello.html的文件。
	//
	// 传入http.Dir类型变量，注意http.Dir是一个类型，其底层类型为string，并不是方法。因而http.Dir("")只是一个类型转换，而非方法调用
	// http.Dir表示文件的起始路径，空即为当前路径。
	mux.Handle("/static/", http.FileServer(http.Dir("")))

	//有时候，我们想要处理器的注册路径和http.Dir的起始路径不相同。有些工具在打包时会将静态文件输出到public目录中。
	//这时需要使用http.StripPrefix方法，该方法会将请求路径中特定的前缀去掉，然后再进行处理
	//这时，请求localhost:8080/static/hello.html将会返回./public/hello.html文件。
	//mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))

	server := &http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
