package kafka

import (
	ckafka "github.com/confluentic/confluent-kafka-go/kafka"
	"fmt"
)

type KafkaProducer struct {
	Producer *ckafka.Producer
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}

// Recebe o host e a porta do produtor para se conectar
func (k *KafkaProducer) SetupProducer(bootstrapServer string) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":bootstrapServer
	}

	k.Producer, err = ckafka.NewProducer(configMap)

	if err != nil {
		fmt.Println(err)
	}
}

func (k *KafkaProducer) Publish(msg string, topic string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: []byte(msg),
	}

	err := k.Producer.Produce(message, deliveryChan:nil)

	if err != nil {
		return err
	}

	return nil
}