package rabbitMQSimple

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://guest:guest@localhost:5672/alex"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Key       string
	Exchange  string
	Mqurl     string
}

//func NewRabbitMQ(queueName string, key string, exchange string) *RabbitMQ {
//	rabbitMQ := &RabbitMQ{
//		QueueName: queueName,
//		Key:       key,
//		Exchange:  exchange,
//		Mqurl:     MQURL,
//	}
//	var err error
//	rabbitMQ.conn,err = amqp.Dial(rabbitMQ.Mqurl)
//	rabbitMQ.ErrFatal(err,"创建rabbitMQ链接错误")
//	rabbitMQ.channel,err = rabbitMQ.conn.Channel()
//	rabbitMQ.ErrFatal(err,"创建channel错误")
//	return rabbitMQ
//}
//func (r * RabbitMQ)Destory(){
//	// 断开Channel、Channel
//	_ = r.channel.Close()
//	_ = r.conn.Close()
//}
//func (r * RabbitMQ)ErrFatal(err error, message string) {
//	if err != nil {
//		log.Fatalf("%s:%s", err, message)
//		panic(fmt.Sprintf("%s：%s",err,message))
//	}
//}
//
//func NewRabbitMQSimple(queueName string) *RabbitMQ {
//	return NewRabbitMQ(queueName,"","")
//}
// 初始化RabbitMQ
func NewRabbitMQ(queueName string, key string, exchange string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Key:       key,
		Exchange:  exchange,
		Mqurl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailErr(err, "链接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailErr(err, "Channel错误")

	return rabbitmq
}

// 报错信息处理
func (r *RabbitMQ) FailErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s,%s", err, message)
		panic(fmt.Sprintf("%s,%s", err, message))
	}
}

// 关闭链接
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// 简单模式Rabbit创建
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式下生产者创建
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列，不存在就创建，存在就不创建
	_, err := r.channel.QueueDeclare(
		// 队列名称
		r.QueueName,
		// 消息是否持久化，服务器重启数据就没有了
		false,
		// 是否自动删除，最后一个消费者断开连接后，是否删除队列中的数据
		false,
		// 是否具有排他性，只有自己才可以访问的队列，其他消费者不可访问
		false,
		// 阻塞，是否等待服务器响应
		false,
		// 其他参数
		nil)
	if err != nil {
		fmt.Println("err is", err)
	}
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 默认false,如果true，则通过routkey规则和exchange类型查找队列，如果查找不到则把消息返回给消费者
		false,
		//如果为true，当exchange发送消息到队列中时，发现队列没有绑定消费者，则把消息返还给消费者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		fmt.Println("err is", err)
	}
	fmt.Println("消息已经发送")
}

// 简单模式下消费者创建
func (r *RabbitMQ) ConsumeSimple() {
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println("err is err")
	}
	msgs, err := r.channel.Consume(
		// 队列名称
		r.QueueName,
		// 消费者名称，用来区分消费者
		"",
		// 默认为true，是否自动应答，是否通知服务器已经接收到消息
		true,
		// 排他性，我创建了只有我自己可见的队列
		false,
		// 默认false，如果设置为true：不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 默认为false，设置为true：不阻塞
		false, nil)
	if err != nil {
		fmt.Println("err is err")
	}
	forever := make(chan bool)
	fmt.Println("等待消息接收")
	for data := range msgs {
		fmt.Println("等待ing")
		go func() {
			// 处理消息逻辑
			log.Printf("data is：%s", data.Body)
			fmt.Printf("data is：%s", data.Body)
		}()
	}
	fmt.Println("消费端接收消息")
	log.Println("消费端接收消息")
	<-forever
}

// 创建广播模式RabbitMQPubSub
func (r *RabbitMQ) NewRabbitMQPubSub(exchangeName strings) *RabbitMQ {
	return NewRabbitMQ("", "", exchangeName)
}

// 广播模式生产者
func (r *RabbitMQ) PublishPub(message string) {
	err := r.channel.ExchangeDeclare(
		// Exchange名称
		r.Exchange,
		// 队列名称
		"fanout",
		true,
		false,
		//默认为false，true时不可以被client用来推送消息，只能用来exchange之间的绑定
		false,
		//阻塞
		false,
		nil)
	r.FailErr(err,"创建Exchange错误")
	err = r.channel.Publish(r.Exchange, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	r.FailErr(err,"发送消息错误")
}

// 广播模式下的消费者
func (r *RabbitMQ) ReceiveSub() {

}