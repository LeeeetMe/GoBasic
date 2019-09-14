package basic_pro

import (
	. "fmt"
	"math/rand"
	"time"
)

func Random_obj() {
	// 伪随机数：每次运行的随机数都相同
	// 设置种子数，种子数变，生成的伪随机数就会变
	rand.Seed(10)
	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		Println("伪随机数为", num)
	}
	// 真随机数
	// 通过设置时间为种子
	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i < 5; i++ {
		num := rand.Intn(10)
		Println("生成真随机数", num)
	}
}
