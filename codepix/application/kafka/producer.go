package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)
// NewKafkaProducer é a função que nos retorna um novo kafka.Producer gerado pela lib.
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return p
}
// Publish é a nossa função para produzir e publicar uma mensagem no Kafka.
func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}
// DeliveryReport() é a função que nos informa se a mensagem foi entregue ou não. Não vamos retornar nada, mas precisamos fazer algo com o kafka.Event...
func DeliveryReport(deliveryChan chan ckafka.Event) {
	// Como um loop pra escultar o nosso Channel.
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
