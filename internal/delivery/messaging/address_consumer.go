package messaging

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/model"
)

type AddressConsumer struct {
	Log *logrus.Logger
}

func NewAddressConsumer(log *logrus.Logger) *AddressConsumer {
	return &AddressConsumer{
		Log: log,
	}
}

func (c AddressConsumer) Consume(message *kafka.Message) error {
	addressEvent := new(model.AddressEvent)
	if err := json.Unmarshal(message.Value, addressEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling address event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic addresses with event: %v from partition %d", addressEvent, message.TopicPartition.Partition)
	return nil
}
