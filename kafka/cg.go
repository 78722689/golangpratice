package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	Brokers   []string
	Topics    []string
	Group     string
	Beginning bool
}

func (cg *ConsumerGroup) Start(reblance string) {
	go cg.start(reblance)
}

func (cg *ConsumerGroup) start(reblance string) {
	config := sarama.NewConfig()
	//config.Version = version

	switch reblance {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
		fmt.Println("Sticky")
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
		fmt.Println("RoundRobin")
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
		fmt.Println("Range")
	default:
		log.Panicf("Unrecognized consumer group partition reblance strategy: %s", reblance)
	}

	if cg.Beginning {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	log.Println(cg.Topics)
	ctx, cancel := context.WithCancel(context.Background())
	group, err := sarama.NewConsumerGroup(cg.Brokers, cg.Group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	defer func() {
		if group != nil {
			if err := group.Close(); err != nil {
				log.Panic(err)
			}
		}
	}()

	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	dispatcher := Dispatcher{}
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Consumer group %s started on topics %v\n", cg.Group, cg.Topics)
		for {
			log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			if err := group.Consume(ctx, cg.Topics, &dispatcher); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			log.Println("yyyyyyyyyyyyyyyyyyyyyyyyyyyy")
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println("Error ctx, %v, exit", ctx.Err())
				return
			}
			//consumer.ready = make(chan bool)
			log.Println("zzzzzzzzzzzzzzzzzzzzz")

			//done <- true
			log.Println("wwwwwwwwwwwwwwwwwwwwwwwwwwww")
		}
	}()

	log.Println("before done")
	<-done
	log.Println("done")
	cancel()
	log.Println("before wait")
	wg.Wait()
	log.Println("Exit")
}

func (cg *ConsumerGroup) RegisterConsumer(topic string, dispatcher *Dispatcher) {

}

type Dispatcher struct {
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (dispatcher *Dispatcher) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	//close(consumer.ready)
	log.Println("Setup")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (dispatcher *Dispatcher) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (dispatcher *Dispatcher) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Printf("ConsumeClaim entry, topic: %s partition: %d memberID: %s GenID: %d\n", claim.Topic(), claim.Partition(), session.MemberID(), session.GenerationID())

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message: value=%s, timestamp=%v, topic=%s, partition=%d", string(message.Value), message.Timestamp, message.Topic, message.Partition)
		session.MarkMessage(message, "")
	}

	log.Println("ConsumeClaim exit")
	return nil
}
