package main

import "io"

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Socket 既实现了Writer接口也实现Closer接口
type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
    return 0, nil
}

func (s *Socket) Close() error {
    return nil
}

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter( writer io.Writer){
    writer.Write( nil )
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser( closer io.Closer) {
    closer.Close()
}

func UseMultiImpl() {
    // 实例化Socket
    s := new(Socket)
    usingWriter(s)
    usingCloser(s)
}