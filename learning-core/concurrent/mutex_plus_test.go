package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 基于Mutex实现tryLock
// 获取等待的goroutine的数量和状态

// 这个从Mutex源码复制过来
const (
	mutexLocked = 1 << iota // 加锁标识位置 1 0001
	mutexWoken				// 唤醒标识位置 2 0010
	mutexStarving			// 锁饥饿标识   4 0100
	mutexWaiterShift = iota	// 标识waiter的起始bit位置 3
)

// 扩展一个Mutex，通过组合原始的Mutex实现扩展
type Mutex struct {
	//Mutex sync.Mutex
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 先通过unsafe去获取sync.Mutex锁的标识位指针地址（即当前mutexLocked的地址值），这里使用了unsafe操作
	rawMutexLockedAddr := (*int32)(unsafe.Pointer(&m.Mutex))

	// 锁初始值是0，如果加锁成功则为1，这里表示能够将加锁标识位从0设置为1成功，则标识获取锁成功
	if atomic.CompareAndSwapInt32(rawMutexLockedAddr, 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁、或者饥饿状态，这次请求就不参与竞争了，直接返回false
	old := atomic.LoadInt32(rawMutexLockedAddr)
	//0001 & 0111 = 0001 加锁
	//0010 & 0111 = 0010 唤醒
	//0100 & 0111 = 0100 饥饿
	//0000 & 0111 = 0000 没有获取锁
	if old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
		return false
	}

	// 尝试在竞争状态下请求锁
	// 其实到这一步old 就是 0
	// new值就是 1
	newValue := old | mutexLocked
	return atomic.CompareAndSwapInt32(rawMutexLockedAddr, old, newValue)
}

// 实现锁的超时机制
// 一、可以通过Context.WithTimeout进行超时机制添加 也可以通过select time.After配合使用
// 二、最简单直接的是采用channel实现，用select监听锁和timeout两个channel，不在今天的讨论范围内。
//1. 用for循环+TryLock实现：
//先记录开始的时间，用for循环判断是否超时：没有超时则反复尝试tryLock，直到获取成功；如果超时直接返回失败。
//问题：高频的CAS自旋操作，如果失败的太多，会消耗大量的CPU。
//
//
//2. 优化1：TryLock的fast的拆分
//TryLock的抢占实现分为两部分，一个是fast path，另一个是竞争状态下的，后者的cas操作很多。我会考虑减少slow方法的频率，比如调用n次fast path失败后，再调用一次整个Trylock。
//3. 优化2：借鉴TCP重试机制
//for循环中的重试增加休眠时间，每次失败将休眠时间乘以一个系数（如1.5），直到达到上限（如10ms），减少自旋带来的性能损耗
func (m *Mutex) LockWithTimeout(timeout time.Duration) {

}

// 获取锁的数量
// Mutex包含两个字段state和sema。前4个字节(int32)就是state字段
// Mutex 结构中的 state 字段有很多个含义，通过 state 字段，你可以知道锁是否已经被某个 goroutine 持有、当前是否处于饥饿状态、是否有等待的 goroutine 被唤醒、等待者的数量等信息。
// 但是，state 这个字段并没有暴露出来，所以，我们需要想办法获取到这个字段，并进行解析。
// 通过unsafe获取state字段的值
// state 这个字段的第一位是用来标记锁是否被持有，第二位用来标记是否已经唤醒了一个等待者，第三位标记锁是否处于饥饿状态，
// 通过分析这个 state 字段我们就可以得到这些状态信息。我们可以为这些状态提供查询的方法，这样就可以实时地知道锁的状态了。
func (m *Mutex) Count() int {
	// 通过 unsafe 操作，我们可以得到 state 字段的值。
	value := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))

	//右移三位（这里的常量 mutexWaiterShift 的值为 3），就得到了当前等待者的数量。
	//value = value >> mutexWaiterShift 		// 得到等待者的数值
	// 如果当前的锁已经被其他 goroutine 持有，那么，我们就稍微调整一下这个值，加上一个 1，你基本上可以把它看作是当前持有和等待这把锁的 goroutine 的总数。
	//value = value + (value & mutexLocked) 	// 再加上锁持有者的数量，0或者1

	value = value >> mutexWaiterShift + (value & mutexLocked)
	return int(value)
}

// 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 锁是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

func TestTryLock(t *testing.T) {
	var mu Mutex
	for true {
		go func() {
			mu.Lock()
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			mu.Unlock()
		}()

		time.Sleep(time.Second)

		ok := mu.TryLock()
		if ok {
			fmt.Println("got the lock")
			mu.Unlock()
			return
		}

		// 没有获取到锁
		fmt.Println("can't get the lock")
	}
}

func TestLockCount(t *testing.T) {
	var mu Mutex
	for true {
		go func() {
			mu.Lock()
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			//mu.Unlock()
		}()
		time.Sleep(time.Second)
		fmt.Println("count: ", mu.Count())
	}
}

func TestLockInfo(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("count: ", mu.Count(), "isLocked:", mu.IsLocked(), "IsWoken:", mu.IsWoken(), "IsStarving:", mu.IsStarving())
}

