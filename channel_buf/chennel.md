## Channel
### 初始化
> channel为进程之间的通信管道
```
// 声明变量
var ch chan int
// 初始化
ch = make(chan int,10)
```
> channel的长度和容量
- 长度：channel内有多少个元素
- 容量：channel最多能添加多少元素

```
func channeLength() {
	ch := make(chan int, 10)
	for i:=0;i<5 	;i++  {
		ch<-i
	}
	Println("channel length is",len(ch))
    //channel length is 5

	ch<-100
	Println("channel length is",len(ch))
    //channel length is 6

	Println("channel cap is",cap(ch))
    //channel cap is 10
}
```

### 阻塞
> channel 有虚拟空间
- channel内没有元素时，此时，获取元素会阻塞。
- channel内的空间已经填满时，再向channel添加元素，会阻塞。

> channel 没有虚拟空间
- channel没有分配空间时，向channel添加元素，会阻塞

### 主线程等待全部goroutine执行完毕
> 使用无capacity阻塞特性使主线程等待goroutine执行完毕
```
func generateInt(i int,ch chan bool){
	time.Sleep(time.Second)
	Println("func generateInt",i)
	// 执行完一次，在ch中添加一个数值
	ch<-true
}

func waitGoroutine() {
	no := 3
	var c chan bool
	c = make(chan bool,no)
	for i := 0; i < no; i++ {
		go generateInt(i,c)
		<-c
	}
}
```
> 使用sync.WaitGroup 进行等待
```
func syncGoroutine(i int,wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	Println("syncGoroutine:",i)
	wg.Done()
}

func syncWait(){
	no :=3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go syncGoroutine(i,&wg)
	}
	wg.Wait()
	Println("all goroutine is over")
}
```