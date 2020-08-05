package main

import (
	"seckillsys/rabbitmq"
	"strconv"
	"time"
)

func main() {

	rq := rabbitmq.NewPubSubRabbitMq("hello")
	defer rq.Destory()

	for i:=0; i<100; i++ {
		rq.PublishPub(strconv.Itoa(i) + "coming!!!")
		time.Sleep(1 * time.Second)
	}
}
