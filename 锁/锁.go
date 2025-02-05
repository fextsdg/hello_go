package main

//两种锁的简单实践

//mutex 适用于所有场景，但性能和效率低
//Rwmutex 适用于读多写少的场景，读与读之间可以并行，读写之间不能并行只能获取锁的权限

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var lock sync.Mutex

// var i,j int =0,0
var rwlock sync.RWMutex

func read(i int) {
	defer wg.Done()
	rwlock.RLock() //读上锁
	fmt.Printf("开始读操作%d!\n", i)
	time.Sleep(time.Second)
	fmt.Printf("读操作%d完成!\n", i)
	rwlock.RUnlock() //读放锁
}

func write(i int) {
	defer wg.Done()
	rwlock.Lock() //写上锁 lock默认是写
	fmt.Printf("开始写操作%d!\n", i)
	time.Sleep(time.Second * 3)
	fmt.Printf("写操作%d完成!\n", i)
	rwlock.Unlock() //写放锁
}

func testRW(r int, w int) {
	wg.Add(r + w)
	for i := 0; i < r; i++ {
		go read(i)
	}
	for i := 0; i < w; i++ {
		go write(i)
	}

}

// 使用加减操作测试mutex作用
var t int = 0

func add1000() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock()
		t++
		lock.Unlock()
	}
}

func sub500() {
	defer wg.Done()
	for i := 0; i < 500; i++ {
		lock.Lock()
		t--
		lock.Unlock()
	}
}

func testMutex() {
	wg.Add(2)
	go add1000()
	go sub500()
}

func main() {
	testRW(5, 2)
	testMutex()

	wg.Wait()
	fmt.Println(t)
}
