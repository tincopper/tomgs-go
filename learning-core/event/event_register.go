package event

// 事件容器
var eventContainer = make(map[string][]func(interface{}))

// 事件注册
func RegisterEvent(name string, callback func(interface{})) {
    // 通过事件名称查找事件列表
    list := eventContainer[name]
    // 注册callback到事件列表
    list = append(list, callback)
    // 重新赋值，将新的事件列表添加到容器
    eventContainer[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
    list := eventContainer[name]
    for _, callback := range list {
        callback(param)
    }
}

