package main

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// type Cond
//  func NeWCond(l Locker) *Cond
//  func (c *Cond) Broadcast() ==> notifyAll
//  func (c *Cond) Signal() ==> notify
//  func (c *Cond) Wait() ==> wait，调用Wait方法时必须持有c.L的锁

func TestBaseCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			cond.L.Lock()
			ready++
			cond.L.Unlock()

			log.Printf("运动员#%d已准备就绪\n", i)
			cond.Broadcast()
		}(i)
	}

	cond.L.Lock()
	for ready != 10 {
		cond.Wait()
		log.Println("裁判员被唤醒一次")
	}
	cond.L.Unlock()

	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}