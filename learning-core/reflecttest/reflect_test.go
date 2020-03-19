package reflecttest

import (
    "fmt"
    "io"
    "reflect"
    "testing"
)

type MyReader struct {
    Name string
}

func (r MyReader) Read(p []byte) (n int, err error) {
    // 实现自己的Read方法
    return 0, nil
}

// 反射基础
// 两个基本概念——Type和Value，它们也是Go语言包中reflect空间里最重要的两个类型。
//
//对所有接口进行反射，都可以得到一个包含Type和Value的信息结构。比如我们对上例的
//reader进行反射，也将得到一个Type和Value， Type为io.Reader， Value为MyReader{"a.txt"}。
//顾名思义， Type主要表达的是被反射的这个变量本身的类型信息，而Value则为该变量实例本身
//的信息
func TestBasicReflect(t *testing.T) {
    //因为MyReader类型实现了io.Reader接口的所有方法（其实就是一个Read()函数），所以
    //MyReader实现了接口io.Reader。我们可以按如下方式来进行MyReader的实例化和赋值：
    var reader io.Reader
    reader = &MyReader{"a.txt"}
    reader.Read(nil)
}

// 获取类型信息
func TestGetTypeInfo(t *testing.T) {
    var f float64 = 3.14
    fmt.Println("type:", reflect.TypeOf(f)) // float64
    v := reflect.ValueOf(f)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())
}

// 获取值信息
// 前提：Go语言中所有的类型都是值类型，即这些变量在传递给函数的时候将发生一次复制。
func TestGetValueInfo(t *testing.T) {
    //var x float64 = 3.4
    //v := reflect.ValueOf(x)
    // fmt.Println("settability of p:" , v.CanSet())
    //v.Set(4)
    /*
    最后一条语句试图修改v的内容。是否可以成功地将x的值改为4.1呢？先要理清v和x的关系。在
    调用ValueOf()的地方，需要注意到x将会产生一个副本，因此ValueOf()内部对x的操作其实
    都是对着x的一个副本。假如v允许调用Set()，那么我们也可以想象出，被修改的将是这个x的
    副本，而不是x本身。如果允许这样的行为，那么执行结果将会非常困惑。调用明明成功了，为
    什么x的值还是原来的呢？为了解决这个问题Go语言，引入了可设属性这个概念（ Settability）。
    如果CanSet()返回false，表示你不应该调用Set()和SetXxx()方法，否则会收到这样的
    错误：
    panic: reflect.Value.SetFloat using unaddressable value
    现在我们知道，有些场景下不能使用反射修改值，那么到底什么情况下可以修改的呢？其实
    这还是跟传值的道理类似。我们知道，直接传递一个float到函数时，函数不能对外部的这个
    float变量有任何影响，要想有影响的话，可以传入该float变量的指针。下面的示例小幅修改
    了之前的例子，成功地用反射的方式修改了变量x的值：
     */
    var x float64 = 3.4
    p := reflect.ValueOf(&x) // 注意：得到X的地址
    fmt.Println("type of p:", p.Type())
    fmt.Println("settability of p:" , p.CanSet())
    v := p.Elem()
    fmt.Println("settability of v:" , v.CanSet())
    v.SetFloat(7.1)
    fmt.Println(v.Interface())
    fmt.Println(x)
}

// 对结构体的反射
//对于结构的反射操作并没有根本上的不同，只是用了Field()方法来按索引获取
//对应的成员。当然，在试图修改成员的值时，也需要注意可赋值属性。
type S struct {
    A int
    B string
}
func TestStructReflect(t *testing.T) {
    v := S{203, "mh203"}
    s := reflect.ValueOf(&v).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}

func TestMap(t *testing.T) {
    mapVal := make(map[string]interface{}) // make创建map返回的是这个map的引用
    var mapVal1 map[string]interface{}

    mapValue := reflect.ValueOf(mapVal)
    //mapValue1 := reflect.ValueOf(mapVal1)
    // 定义的变量mapVal1不能直接进行传值，如果要对值进行修改需要传地址
    mapValue1 := reflect.Indirect(reflect.ValueOf(&mapVal1))
    //mapVal1Addr := reflect.ValueOf(&mapVal1)
    //fmt.Println(mapVal1Addr.CanSet())
    //mapValue1 := mapVal1Addr.Elem()

    if mapValue1.IsNil() {
        fmt.Println("mapValue1 is nil")
        fmt.Println("type is: ", mapValue1.Type())
        tmp := reflect.MakeMap(mapValue1.Type())

        mapValue1.Set(tmp)
    }

    if mapValue1.IsNil() {
        fmt.Println("mapValue1 is nil")
    }

    mapValue.SetMapIndex(reflect.ValueOf("kkk"), reflect.ValueOf("bbb"))
    mapValue.SetMapIndex(reflect.ValueOf("aaa"), reflect.ValueOf("ccc"))
    fmt.Println(mapVal)

    mapValue1.SetMapIndex(reflect.ValueOf("kkk"), reflect.ValueOf("bbb"))
    fmt.Println(mapVal1)
}

func TestReflect(t *testing.T) {
    typeOf := reflect.TypeOf(123)
    fmt.Println(typeOf)
}