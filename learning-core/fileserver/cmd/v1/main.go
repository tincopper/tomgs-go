package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

// 注意工作目录位置
var dir = "."

// 原始方式
func main() {
	mux := http.NewServeMux()
	// 创建了一个文件处理器，将它挂载到路径/static/上, 这里的注册路径/static/最后的/不能省略
	mux.HandleFunc("/static/", fileHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := dir + r.URL.Path
	fmt.Println(path)

	f, err := os.Open(path)
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}

	if d.IsDir() {
		DirList(w, r, f)
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}

	ext := filepath.Ext(path)
	if contentType := extensionToContentType[ext]; contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))
	w.Write(data)
}

var extensionToContentType = map[string]string {
	".html": "text/html; charset=utf-8",
	".css": "text/css; charset=utf-8",
	".js": "application/javascript",
	".xml": "text/xml; charset=utf-8",
	".jpg":  "image/jpeg",
}

func DirList(w http.ResponseWriter, r *http.Request, f http.File) {
	dirs, err := f.Readdir(-1)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		url := url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), name)
	}
	fmt.Fprintf(w, "</pre>\n")
}

func toHTTPError(err error) int {
	if os.IsNotExist(err) {
		return http.StatusNotFound
	}
	if os.IsPermission(err) {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

func Error(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}