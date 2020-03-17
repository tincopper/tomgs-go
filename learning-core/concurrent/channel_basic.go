package main

import (
    "fmt"
    "time"
)

func ChannelBasic() {
    //一般channel的声明形式为：
    //var chanName chan ElementType
    // 声明一个传递类型为int的channel
    // var i chan int
    // 声明一个map，元素是bool型的channel，即map的value的类型是bool类型的channel的
    // var m map[string] chan bool
    
    // 定义一个channel也很简单，直接使用内置的函数make()即可：
    // ch := make(chan int)
    // 这就声明并初始化了一个int型的名为ch的channel。
    
    // 写如数据到channel
    // ch <- value
    
    // 从channel读取数据
    // result := <- ch
}

// select机制
//早在Unix时代， select机制就已经被引入。通过调用select()函数来监控一系列的文件句
//柄，一旦其中一个文件句柄发生了IO动作，该select()调用就会被返回。后来该机制也被用于
//实现高并发的Socket服务器程序。 Go语言直接在语言级别支持select关键字，用于处理异步IO
//问题。
//
//select的用法与switch语言非常类似，由select开始一个新的选择块，每个选择条件由
//case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比， select有比较多的
//限制，其中最大的一条限制就是每个case语句里必须是一个channel操作，大致的结构如下：
//select {
// case <-chan1: 如果chan1成功读到数据，则进行该case处理语句
// case chan2 <- 1: 如果成功向chan2写入数据，则进行该case处理语句
// default: 如果上面都没有成功，则进入default处理流程
//}

// 只要select中的case有一个条件满足，那么程序就会往下执行，这个特性也为后面的channel超时机制提供了支持。
func SelectDemo() {
    ch := make(chan int, 1)
    for {
        select {
        case ch <- 0:
        case ch <- 1:
        }
        i := <-ch
        fmt.Println("Value received:", i)
    }
}

// channel缓冲机制
func ChannelCache() {
    // 创建一个带有缓存的channel
    // 在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面这个例子就创建了一个大小
    //为1024的int类型channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被
    //填完之前都不会阻塞。
    ch := make(chan int, 1024)
    // 循环读取channel中的数据
    for i := range ch {
        fmt.Println("Received:", i)
    }
}

// channel超时机制
//在并发编程的通信过程中，最需要处理的就是超时问题，即向channel写数据时发现channel
//已满，或者从channel试图读取数据时发现channel为空。如果不正确处理这些情况，很可能会导
//致整个goroutine锁死。
//
//超过设定的时间时，仍然没有处理完任务，则该方法会立即终止并
//返回对应的超时信息。超时机制本身虽然也会带来一些问题，比如在运行比较快的机器或者高速
//的网络上运行正常的程序，到了慢速的机器或者网络上运行就会出问题，从而出现结果不一致的
//现象，但从根本上来说，解决死锁问题的价值要远大于所带来的问题。
func ChannelTimeOut() {
    //使用channel时需要小心，比如对于以下这个用法：
    //i := <-ch
    //不出问题的话一切都正常运行。但如果出现了一个错误情况，即永远都没有人往ch里写数据，那
    //么上述这个读取动作也将永远无法从ch中读取到数据， 导致的结果就是整个goroutine永远阻塞并
    //没有挽回的机会。
    
    //Go语言没有提供直接的超时处理机制，但我们可以利用select机制。虽然select机制不是
    //专为超时而设计的，却能很方便地解决超时问题。因为select的特点是只要其中一个case已经
    //完成，程序就会继续往下执行，而不会考虑其他case的情况。
    // 首先，我们实现并执行一个匿名的超时等待函数
    ch := make(chan int)
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(1e9) // 等待1秒钟
        timeout <- true
    }()
    // 然后我们把timeout这个channel利用起来
    select {
    case <- ch:
        // 从ch中读取到数据
    case <-timeout:
        // 一直没有从ch中读取到数据，但从timeout中读取到了数据
    }
}

// channel的传递
//在Go语言中channel本身也是一个原生类型，与map之类的类型地位一样，因
//此channel本身在定义后也可以通过channel来传递。
//
//
func Pipeline() {
    p := make(chan *PipeData)
    
    p2 := new(PipeData)
    p2.value = 1
    p2.handler = func(i int) int {
        return 1
    }
    p2.next <- 1
    
    p <- p2
    handle(p)
}

type PipeData struct {
    value int
    handler func(int) int
    next chan int
}

func handle(queue chan *PipeData) {
    for data := range queue {
        data.next <- data.handler(data.value)
    }
}

// 单向通道
//channel本身必然是同时支持读写的，
//否则根本没法用。假如一个channel真的只能读，那么肯定只会是空的，因为你没机会往里面写数
//据。同理，如果一个channel只允许写，即使写进去了，也没有丝毫意义，因为没有机会读取里面
//的数据。所谓的单向channel概念，其实只是对channel的一种使用限制。
//
//我们在将一个channel变量传递到一个函数时，可以通过将其指定为单向channel变量，从
//而限制 该函数中可 以对此 channel的操作， 比如只能往 这个 channel写，或者只 能从这个
//channel读。
//
//为什么要做这样的限制呢？从设计的角度考虑，所有的代码应该都遵循“最小权限原则”，
//从而避免没必要地使用泛滥问题，进而导致程序失控。
func simplexChannel() {
    //单向channel变量的声明非常简单，如下：
    //var ch1 chan int // ch1是一个正常的channel，不是单向的
    //var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
    //var ch3 <-chan int // ch3是单向channel，只用于读取int数据
    
    //channel是一个原生类型，因此不仅支持被传递，还支持类型转换。就是在单向channel和双向channel之间进行转换
    //ch4 := make(chan int)
    //ch5 := <- chan int(ch4) // ch5就是一个单向的读取channel
    //ch6 := chan <- int(ch4) // ch6 是一个单向的写入channel
}

// 下面我们来看一下单向channel的用法：
func Parse(ch <-chan int) {
    for value := range ch {
        fmt.Println("Parsing value", value)
    }
}

// channel关闭
func closeChannel() {
    ch := make(chan int)
    // go提供了内置函数close来进行channel的关闭动作
    close(ch)
    // 判断ch是否已经关闭
    value, ok := <-ch
    if !ok {
        fmt.Println("ch已经关闭")
    }
    fmt.Println("value:", value)
}

func main() {

}