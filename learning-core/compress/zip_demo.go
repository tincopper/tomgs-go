package compress

import (
    "archive/zip"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

// zip compress
func Compress(file, dest string) error {
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    defer f.Close()

    d, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer d.Close()

    writer := zip.NewWriter(d)
    defer writer.Close()

    // 压缩文件内的文件名称
    //w, err := writer.Create("test.txt")
    _, fileName := filepath.Split(file)
    w, err := writer.Create(fileName)
    if err != nil {
        return err
    }

    _, err = io.Copy(w, f)
    return err
}

func Compress2(file string, response http.ResponseWriter, request *http.Request) error {
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    defer func() {_ = f.Close()}()

    /*dest := "test.zip"
    d, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer d.Close()*/

    writer := zip.NewWriter(response)
    defer func() {_ = writer.Close()}()

    // 压缩文件内的文件名称
    _, fileName := filepath.Split(file)
    w, err := writer.Create(fileName)
    if err != nil {
        return err
    }

    _, err = io.Copy(w, f)
    return err
}

func Decompress() error {
    return nil
}