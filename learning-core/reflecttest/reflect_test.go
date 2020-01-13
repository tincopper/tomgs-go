package reflecttest

import (
    "fmt"
    "reflect"
    "testing"
)

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