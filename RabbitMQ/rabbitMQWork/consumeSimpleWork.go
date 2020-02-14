package rabbitMQWork

import "RabbitMQ/rabbitMQSimple"

func main() {
	rabbitmq := rabbitMQSimple.NewRabbitMQSimple("AreYouOK")
	//defer rabbitmq.Destory()
	rabbitmq.ConsumeSimple()
}

