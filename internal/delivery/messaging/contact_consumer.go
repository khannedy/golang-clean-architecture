package messaging

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/model"
)

type ContactConsumer struct {
	Log *logrus.Logger
}

func NewContactConsumer(log *logrus.Logger) *ContactConsumer {
	return &ContactConsumer{
		Log: log,
	}
}

func (c ContactConsumer) Consume(message *kafka.Message) error {
	ContactEvent := new(model.ContactEvent)
	if err := json.Unmarshal(message.Value, ContactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %d", ContactEvent, message.TopicPartition.Partition)
	return nil
}
