#select
- 与switch case用法类似，同时等待多个channel ready，哪个先ready就进入哪个分支
- 如果两个channel 同时ready，则随机选择其中一个分支执行
- 如果全部channel都没有ready，且有default，则进入default分支，没有default则阻塞住等待channel的ready状态

## select使用
+ 快速判断channel 内没有填充元素
```
func selectChannel() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	go server1(ch1)
	go server2(ch2)
	select {
	case s1 := <-ch1:
		Println(s1)
	case s2 := <-ch2:
		Println(s2)
	default:
		Println("default：所有的channel都为空")
	}
}
```
+ 快速判断channel 空间已满
```
func getChannel(ch chan string) {
	time.Sleep(time.Second)
	for value := range ch {
		Println("get value from ch:",value)
	}
}
func selectFullChannel()  {
	/*
		检测channel是否已满
	*/
	ch := make(chan string,10)
	go putIntoChannel(ch)
	go getChannel(ch)
	for{
		select {
		case ch <- "good":
			Println("ch input good")
		default:
			Println("ch已经满了")
		}
	}
}
```