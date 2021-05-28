package kafka

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type Producer struct {
	Brokers   []string
	Topic     string
	Partition int32
}

func (p *Producer) Send(message string) {
	go p.ProduceMessage(message)
}

func (p *Producer) ConnectProducer() ( /*sarama.SyncProducer*/ sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	if p.Partition != -1 {
		config.Producer.Partitioner = sarama.NewManualPartitioner
	}
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	//conn, err := sarama.NewSyncProducer(brokersUrl, config)
	conn, err := sarama.NewAsyncProducer(p.Brokers, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (p *Producer) ProduceMessage(message string) {
	producer, err := p.ConnectProducer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			select {
			case e := <-producer.Errors():
				fmt.Println(e)
			case s := <-producer.Successes():
				fmt.Printf("Message \"%s\" is sent successfully to topic \"%s\", partition \"%d\"\n", s.Value, s.Topic, s.Partition)
			}
		}
	}()

	for {
		if message == "" {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			message = scanner.Text()

			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
				continue
			}
		}

		msg := &sarama.ProducerMessage{
			Topic:     p.Topic,
			Value:     sarama.StringEncoder(message),
			Partition: p.Partition,
		}

		producer.Input() <- msg
		//fmt.Printf("Message \"%s\" is stored in topic-%s, partition-%d\n", message, p.Topic, msg.Partition)
		message = ""
	}
}
