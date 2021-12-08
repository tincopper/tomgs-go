package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ServeFileContent(w http.ResponseWriter, r *http.Request, name string, modTime time.Time) {
	f, err := os.Open(name)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "open file error:", err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "call stat error:", err)
		return
	}

	if fi.IsDir() {
		w.WriteHeader(400)
		fmt.Fprint(w, "no such file:", name)
		return
	}

	http.ServeContent(w, r, name, fi.ModTime(), f)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("file")

	if filename == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "filename is empty")
		return
	}

	ServeFileContent(w, r, filename, time.Time{})
}

// http.ServeContent 方式
// 除了直接使用http.FileServer之外，net/http库还暴露了ServeContent方法。这个方法可以用在处理器需要返回一个文件内容的时候，非常易用。
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/show", fileHandler)

	server := &http.Server {
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}