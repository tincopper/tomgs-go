package main

import (
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	helloworldpb "tomgs-go/learning-grpc-gateway/hello-world/api"
	"tomgs-go/learning-grpc-gateway/hello-world/pkg/ui/data"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux, err := newGateway(ctx)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/serve", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	mux.HandleFunc("/docs/", fileHandler)
	serveSwaggerUI(mux)

	log.Println("grpc-gateway listen on localhost:9090")
	return http.ListenAndServe(":9090", mux)
}

func newGateway(ctx context.Context) (http.Handler, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := helloworldpb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, ":50051", opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	log.Println("start serveSwaggerFile")

	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("api/", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	var fileServer = http.FileServer(&assetfs.AssetFS{
		Asset:    data.Asset,
		AssetDir: data.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	//mux.Handle(prefix, fileServer)
}

var dir = "api"
func fileHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s\n", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/docs/")
	p = path.Join(dir, p)

	//p := "." + r.URL.Path
	//path := "api"
	fmt.Println(p)

	f, err := os.Open(p)
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

	dataBytes, err := ioutil.ReadAll(f)
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}

	ext := filepath.Ext(p)
	if contentType := extensionToContentType[ext]; contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))
	_, _ = w.Write(dataBytes)
}

var extensionToContentType = map[string]string{
	".html": "text/html; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".js":   "application/javascript",
	".xml":  "text/xml; charset=utf-8",
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
		if strings.HasSuffix(name, "json") {
			//url := url.URL{Path: name}
			//fmt.Sprintf("http://localhost:9090/swagger-ui/index1.html?file=http://localhost:9090/swagger/%s", name)
			host := "localhost:9090"
			urlInfo := url.URL{
				Path:     "/swagger-ui/index1.html",
				Host:     host,
				RawQuery: fmt.Sprintf("file=http://%s/swagger/%s", host, name),
			}
			fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", urlInfo.String(), name)
		} else {
			urlInfo := url.URL{Path: name}
			fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", urlInfo.String(), name)
		}
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

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
