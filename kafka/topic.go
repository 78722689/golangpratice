package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Topic struct {
	Brokers   []string
	Topic     string
	Partition int32
}

func (topic *Topic) Create() {
	config := sarama.NewConfig()
	//config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(topic.Brokers, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	err = admin.CreateTopic(topic.Topic, &sarama.TopicDetail{
		NumPartitions:     topic.Partition,
		ReplicationFactor: 1,
	}, false)

	if err != nil {
		log.Fatal("Error while creating topic: ", err.Error())
	} else {
		log.Println("%s created successfully", topic.Topic)
	}
}

func (topic *Topic) Delete() {
	config := sarama.NewConfig()
	//config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(topic.Brokers, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	err = admin.DeleteTopic(topic.Topic)

	if err != nil {
		log.Fatal("Error while deleting topic: ", err.Error())
	} else {
		log.Printf("%s has been deleted successfully\n", topic.Topic)
	}
}
