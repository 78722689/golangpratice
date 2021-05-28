package kafka

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	Brokers   []string
	Topic     string
	Partition int32
	Beginning bool
}

/*
func StartConsumers(num, partition int, topic string, brokers []string) {
	for i := 0; i < num; i++ {
		consumer := &Consumer{Brokers: brokers, Topic: topic, Partition: partition}
		consumer.ConsumeMessage(topic, brokers)
	}
}
*/

func (c *Consumer) connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// NewConsumer creates a new consumer using the given broker addresses and configuration
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *Consumer) Consume() {
	worker, err := c.connectConsumer(c.Brokers)
	if err != nil {
		log.Fatal(err)
	}

	offset := sarama.OffsetNewest
	if c.Beginning {
		offset = sarama.OffsetOldest
	}

	consumer, err := worker.ConsumePartition(c.Topic, c.Partition, offset)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := worker.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Consumer started")
	for {
		select {
		case msg := <-consumer.Messages():
			fmt.Printf("Received message \"%s\" on topic \"%s\" \n", string(msg.Value), msg.Topic)
		}
	}
}
