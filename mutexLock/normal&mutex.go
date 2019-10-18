package main
import (
	. "fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func Add(x *int) {
	Println("normal")
	time.Sleep(time.Second)
	for i := 0; i < 5000; i++ {
		*(x)++
	}
	wg.Done()
}

func mutexAdd(x *int) {
	mutex.Lock()
	time.Sleep(time.Second)
	Println("mutex")
	for i := 0; i < 5000; i++ {
		*(x)++
	}
	wg.Done()
	mutex.Unlock()
}
func testNormal(){
	// 不加锁会导致数据错误
	x := 0
	wg.Add(2)
	go Add(&x)
	go Add(&x)
	wg.Wait()
	Println("不加锁x结果：",x)
}

func testMutexLock(){
	// 枷锁后正确
	x := 0
	wg.Add(2)
	go mutexAdd(&x)
	go mutexAdd(&x)
	wg.Wait()
	Println("加锁x结果：",x)
}