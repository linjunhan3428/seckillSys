package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//连接信息
const MQURL = "amqp://guest:guest@localhost:5672/"

//rabbitMQ结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange  string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl     string
}

// 创建结构体实例
func NewRabbitMq(queueName string, exchange string, key string) *RabbitMQ{

	return &RabbitMQ{
		QueueName: queueName,
		Exchange: exchange,
		Key: key,
		Mqurl: MQURL,
	}
}

// 断开channel和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 处理错误函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err.Error())
	}
}

//   -------------------------------simple------------------------------------
// 创建简单模式下RabbitMQ实例
func NewSimpleRabbitMQ(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMq(queueName, "", "")

	//获取connection
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabb"+"itmq!")

	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// 简单模式下生产
func (r *RabbitMQ) PublishSimple(message string) {
	// 需要通过channel申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	r.failOnErr(err, "apply queue failed !")
	// 通过channel将消息发送到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// simple 模式下消费信息
func (r *RabbitMQ) ConsumeSimple() {
	// 申请队列
	queue, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	r.failOnErr(err, "apply queue failed")

	// 接收消息
	msgs, err := r.channel.Consume(queue.Name, "", true, false, false, false, nil)
	r.failOnErr(err, "消息接收失败")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// ------------------------------------发布订阅模式---------------------------------------------------
// 创建发布订阅模式的rabbitmq
func NewPubSubRabbitMq(exchangeName string) *RabbitMQ {

	rq := NewRabbitMq("", exchangeName, "")

	var err error
	rq.conn, err = amqp.Dial(MQURL)
	rq.failOnErr(err, "failed to connect rabbmq")

	rq.channel, err = rq.conn.Channel()
	rq.failOnErr(err, "failed to open a channel")
	return rq
}

// 发送消息
func (r *RabbitMQ) PublishPub(message string) {
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "declare exchange failed")

	// 发送消息
	err = r.channel.Publish(r.Exchange, "", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	r.failOnErr(err, "push message failed")
}

// 消费消息
func (r *RabbitMQ) RecieveSub() {
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "declare exchange failed")

	queue, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	r.failOnErr(err, "queue declare failed !")

	err = r.channel.QueueBind(queue.Name, "", r.Exchange, false, nil)
	r.failOnErr(err, "queue bind failed !")

	message, err := r.channel.Consume(queue.Name, "", true, false, false, false, nil)
	r.failOnErr(err, "consume failed!")

	forever := make(chan struct{})

	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("请按 ctrl + c 退出")
	<-forever
}