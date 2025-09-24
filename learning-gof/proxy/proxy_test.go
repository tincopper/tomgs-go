package proxy

import (
	"fmt"
	"testing"
)

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}

}

// 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	// 处理前
	fmt.Println("前置处理：处理其他业务")
	// 调用真实对象
	proxy.station.sell(name)
	// 处理后
	fmt.Println("后置处理：处理其他业务")
}

func NewStationProxy(seller Seller) Seller {
	proxy := &StationProxy{station: seller.(*Station)}
	return proxy
}

func TestStaticProxy(test *testing.T) {
	station := &Station{stock: 10}
	proxy := NewStationProxy(station)
	proxy.sell("张三")
	proxy.sell("李四")
}
