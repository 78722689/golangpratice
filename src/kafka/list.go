package kafka

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type List struct {
	Brokers []string
	Topic   string
}

func (list *List) Get() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//get broker
	consumer, err := sarama.NewConsumer(list.Brokers, config)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	var topics []string
	if list.Topic == "" {
		//get all topic from consumer
		topics, err = consumer.Topics()
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		topics = append(topics, list.Topic)
	}

	for _, topic := range topics {
		partitions, err := consumer.Partitions(topic)
		if err != nil {
			log.Fatal(err)
			continue
		}

		combine := ""
		for _, partition := range partitions {
			offset, err := list.offset(topic, partition)
			if err != nil {
				log.Fatal(err)
				combine += fmt.Sprintf("(%d:%d) ", partition, -1)
				continue
			}
			combine += fmt.Sprintf("(%d:%d) ", partition, offset)
		}

		fmt.Printf("Topic: %s, Partitions-offsets: %s\n", topic, combine)
	}
}

func (list *List) offset(topic string, partition int32) (int64, error) {
	client, err := sarama.NewClient(list.Brokers, nil)
	if err != nil {
		return 0, err
	}

	return client.GetOffset(topic, partition, sarama.OffsetNewest)
}
