package main

import (
	"seckillsys/rabbitmq"
)

func main() {
	rq := rabbitmq.NewSimpleRabbitMQ("ljh")
	defer rq.Destory()

	rq.PublishSimple("i am coming")
}
