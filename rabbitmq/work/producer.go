package main

import (
	"seckillsys/rabbitmq"
	"strconv"
	"time"
)

func main() {

	rq := rabbitmq.NewSimpleRabbitMQ("hello")
	defer rq.Destory()

	for i:=1; i<=100; i++ {
		rq.PublishSimple(strconv.Itoa(i) + "coming!!!")
		time.Sleep(1 * time.Second)
	}
}
