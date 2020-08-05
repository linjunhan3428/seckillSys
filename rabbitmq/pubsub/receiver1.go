package main

import "seckillsys/rabbitmq"

func main() {

	rq := rabbitmq.NewPubSubRabbitMq("hello")
	defer rq.Destory()

	rq.RecieveSub()
}
