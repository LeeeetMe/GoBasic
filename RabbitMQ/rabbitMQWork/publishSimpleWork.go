package rabbitMQWork

import (
	"RabbitMQ/rabbitMQSimple"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitMQSimple.NewRabbitMQSimple("AreYouOK")
	defer rabbitmq.Destory()
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("Hello World"+strconv.FormatInt(int64(i),10))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}