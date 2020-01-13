package filetest

import (
    "bufio"
    "io"
    "io/ioutil"
    "os"
)

/*
    文件读取的方式
 */
// 一次性读取，ioutil.ReadFile和ioutil.ReadAll，可以一次性将文件内容读取到内存，缺点是会占用大量内存
func ReadAll(filePth string) ([]byte, error) {
    //ioutil.ReadFile(filePth)

    f, err := os.Open(filePth)
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(f)
}

// 分块读取，可在速度和内存占用之间取得很好的平衡。
func processBlock(line []byte) {
    os.Stdout.Write(line)
}

func ReadBlock(filePth string, bufSize int, hookfn func([]byte)) error {
    f, err := os.Open(filePth)
    if err != nil {
        return err
    }
    defer f.Close()

    buf := make([]byte, bufSize) //一次读取多少个字节
    bfRd := bufio.NewReader(f)
    for {
        n, err := bfRd.Read(buf)
        hookfn(buf[:n]) // n 是成功读取字节数

        if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
            if err == io.EOF {
                return nil
            }
            return err
        }
    }

    return nil
}

func TestProcessBlock() {
    ReadBlock("test.txt", 10000, processBlock)
}

// 逐行读取,逐行读取有的时候真的很方便，性能可能慢一些，但是仅占用极少的内存空间。
func processLine(line []byte) {
    os.Stdout.Write(line)
}

func ReadLine(filePth string, hookfn func([]byte)) error {
    f, err := os.Open(filePth)
    if err != nil {
        return err
    }
    defer f.Close()

    bfRd := bufio.NewReader(f)
    for {
        line, err := bfRd.ReadBytes('\n')
        hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
        if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
    return nil
}

func TestProcessReadLine() {
    ReadLine("test.txt", processLine)
}