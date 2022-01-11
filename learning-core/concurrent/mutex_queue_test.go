package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 使用mutex实现线程安全的队列
// 直接使用slice实现的队列在Dequeue和Enqueue时是线程不安全的，通过Mutex在入队和出队时加锁保证数据安全
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{
		data: make([]interface{}, 0, n),
	}
}

// 将元素放在队尾
func (q *SliceQueue) Enqueue(ele interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, ele)
}

// 从队头获取元素
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func TestSliceQueue(t *testing.T) {
	queue := NewSliceQueue(10)
	go func() {
		for i := 0; i < 5; i++ {
			queue.Enqueue(fmt.Sprint("a:", i))
		}
	}()

	go func() {
		for i := 5; i < 10; i++ {
			queue.Enqueue(fmt.Sprint("a:", i))
		}
	}()

	time.Sleep(time.Duration(1000))

	for i := 0; i < 10; i++ {
		fmt.Println(queue.Dequeue())
	}
}