package main

import "io"

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Socket 既实现了Writer接口也实现Closer接口
// 因为Socket类型实现了io.Writer接口的所有方法（其实就是一个Write()函数），所以
// Socket实现了接口io.Write
// 所以从这里也可以理解，go的实现是怎么一个过程，只要实现了接口的方法，那么就实现了该接口，而不需要像Java里面的implements关键字来实现接口。
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