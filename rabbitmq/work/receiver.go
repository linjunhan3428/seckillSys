package main

import "seckillsys/rabbitmq"

func main() {
	rq := rabbitmq.NewSimpleRabbitMQ("hello")
	defer rq.Destory()
	rq.ConsumeSimple()
}
