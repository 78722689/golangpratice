package kafka

import (
	"os"
	"os/signal"
	"strings"
)

func KafkaMain(arg *Arguments) {

	brokers := []string{arg.Broker}

	singals := make(chan os.Signal, 1)
	signal.Notify(singals, os.Interrupt)

	switch arg.Operation {
	case "create":
		topic := &Topic{Brokers: brokers, Topic: arg.Topic, Partition: int32(arg.Partition)}
		topic.Create()
	case "delete":
		topic := &Topic{Brokers: brokers, Topic: arg.Topic}
		topic.Delete()
	case "list":
		list := &List{Brokers: brokers, Topic: arg.Topic}
		list.Get()
	case "produce":
		producer := &Producer{Brokers: brokers, Topic: arg.Topic, Partition: int32(arg.Partition)}
		producer.Send(arg.Message)
	case "consume":
		consumer := &Consumer{Brokers: brokers, Topic: arg.Topic, Partition: int32(arg.Partition), Beginning: arg.Beginning}
		go consumer.Consume()
	case "cg":
		cg := &ConsumerGroup{Brokers: brokers, Topics: strings.Split(arg.Topic, ","), Beginning: arg.Beginning, Group: arg.Group}
		cg.Start(arg.Reblance)

	}
	<-singals
}
